package resolvers

import (
	"fmt"
	"reflect"

	users_commands "github.com/guicostaarantes/go-auth/modules/users/commands"
	users_queries "github.com/guicostaarantes/go-auth/modules/users/queries"
	users_workers "github.com/guicostaarantes/go-auth/modules/users/workers"
	hash_util "github.com/guicostaarantes/go-auth/utils/hash"
	log_util "github.com/guicostaarantes/go-auth/utils/log"
	store_util "github.com/guicostaarantes/go-auth/utils/store"
	stream_util "github.com/guicostaarantes/go-auth/utils/stream"
	token_store_util "github.com/guicostaarantes/go-auth/utils/token_store"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Utils    *ResolverUtils
	Commands *ResolverCommands
	Queries  *ResolverQueries
	Workers  *ResolverWorkers
}

type ResolverUtils struct {
	Hash       hash_util.I
	Log        log_util.I
	Store      store_util.I
	Stream     stream_util.I
	TokenStore token_store_util.I
}

type ResolverCommands struct {
	AuthenticateUser *users_commands.AuthenticateUser
	CreateUser *users_commands.CreateUser
}

type ResolverQueries struct {
	UserByID          *users_queries.UserByID
	UserIDByAuthToken *users_queries.UserIDByAuthToken
}

type ResolverWorkers struct {
	UserCreated *users_workers.UserCreatedWorker
}

// Function that checks if any dependencies for a query or command are missing
func anyNilFields(v interface{}) (string, string) {
	w := reflect.ValueOf(v).Elem()
	for j := 0; j < w.NumField(); j++ {
		x := w.Field(j).Elem()
		for k := 0; k < x.NumField(); k++ {
			if x.Field(k).Interface() == nil {
				return w.Type().Field(j).Name, x.Type().Field(k).Name
			}
		}
	}
	return "", ""
}

func CreateResolver(utils *ResolverUtils) (*Resolver, error) {
	commands := &ResolverCommands{
		AuthenticateUser: &users_commands.AuthenticateUser{
			HashUtil:   utils.Hash,
			LogUtil:    utils.Log,
			StoreUtil:  utils.Store,
			TokenStoreUtil:  utils.TokenStore,
		},
		CreateUser: &users_commands.CreateUser{
			HashUtil:   utils.Hash,
			LogUtil:    utils.Log,
			StreamUtil: utils.Stream,
			StoreUtil:  utils.Store,
		},
	}

	queries := &ResolverQueries{
		UserByID: &users_queries.UserByID{
			LogUtil:   utils.Log,
			StoreUtil: utils.Store,
		},
		UserIDByAuthToken: &users_queries.UserIDByAuthToken{
			LogUtil:        utils.Log,
			TokenStoreUtil: utils.TokenStore,
		},
	}

	workers := &ResolverWorkers{
		UserCreated: &users_workers.UserCreatedWorker{
			LogUtil:    utils.Log,
			StoreUtil:  utils.Store,
			StreamUtil: utils.Stream,
		},
	}

	missingService, missingDependency := anyNilFields(commands)
	if missingService != "" {
		return nil, fmt.Errorf("missing dependency %s for command %s", missingDependency, missingService)
	}

	missingService, missingDependency = anyNilFields(queries)
	if missingService != "" {
		return nil, fmt.Errorf("missing dependency %s for query %s", missingDependency, missingService)
	}

	missingService, missingDependency = anyNilFields(workers)
	if missingService != "" {
		return nil, fmt.Errorf("missing dependency %s for worker %s", missingDependency, missingService)
	}

	return &Resolver{
		Utils:    utils,
		Commands: commands,
		Queries:  queries,
		Workers:  workers,
	}, nil
}
