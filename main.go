package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/guicostaarantes/go-auth/graph"
	"github.com/guicostaarantes/go-auth/graph/resolvers"
	users_commands "github.com/guicostaarantes/go-auth/modules/users/commands"
	users_models "github.com/guicostaarantes/go-auth/modules/users/models"
	hash_util "github.com/guicostaarantes/go-auth/utils/hash"
	log_util "github.com/guicostaarantes/go-auth/utils/log"
	store_util "github.com/guicostaarantes/go-auth/utils/store"
	stream_util "github.com/guicostaarantes/go-auth/utils/stream"
	token_store_util "github.com/guicostaarantes/go-auth/utils/token_store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	postgresDsn := os.Getenv("POSTGRES_DSN")
	if postgresDsn == "" {
		panic("missing POSTGRES_DSN environment variable")
	}

	redisConnectionUrl := os.Getenv("REDIS_CONNECTION_URL")
	if redisConnectionUrl == "" {
		panic("missing REDIS_CONNECTION_URL environment variable")
	}

	kafkaBootstrapServers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if kafkaBootstrapServers == "" {
		panic("missing KAFKA_BOOTSTRAP_SERVERS environment variable")
	}

	kafkaUniqueConsumerId := os.Getenv("KAFKA_UNIQUE_CONSUMER_ID")
	if kafkaUniqueConsumerId == "" {
		panic("missing KAFKA_UNIQUE_CONSUMER_ID environment variable")
	}

	port := "8080"

	db, err := gorm.Open(postgres.Open(postgresDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&users_models.User{},
	)
	if err != nil {
		panic(err)
	}
	
	redisOptions, err := redis.ParseURL(redisConnectionUrl)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(redisOptions)

	utils := &resolvers.ResolverUtils{
		Hash:       hash_util.BCryptImpl{Cost: 8},
		Log:        log_util.FmtImpl{},
		Store:      store_util.GormImpl{Db: db},
		Stream:     stream_util.KafkaImpl{
      BootstrapServers: kafkaBootstrapServers,
      UniqueConsumerID: kafkaUniqueConsumerId,
    },
		TokenStore: token_store_util.RedisImpl{RedisClient: rdb},
	}

	res, err := resolvers.CreateResolver(utils)
	if err != nil {
		panic(err)
	}

  bootstrapAdmin := os.Getenv("BOOTSTRAP_ADMIN")
  if bootstrapAdmin != "" {
    s := strings.Split(bootstrapAdmin, "|")
    input := users_commands.CreateUserInput{
      Active: true, 
      Email: s[0],
      Password: s[1],
      Role: users_models.Admin,
    }
    res.Commands.CreateUser.ExecuteCommand(input)
  }

	res.Workers.UserCreated.ExecuteWorker()

	router := graph.CreateServer(res)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
