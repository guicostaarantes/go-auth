package graph

import (
	"context"
	"errors"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/guicostaarantes/go-auth/graph/generated"
	"github.com/guicostaarantes/go-auth/graph/resolvers"
	users_models "github.com/guicostaarantes/go-auth/modules/users/models"
)

func CreateServer(res *resolvers.Resolver) *chi.Mux {
	c := generated.Config{Resolvers: res}

	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []users_models.Role) (interface{}, error) {
		userID := ctx.Value("userID").(string)

		if userID == "" {
			return nil, errors.New("forbidden")
		}

		user, userErr := res.Queries.UserByID.ExecuteQuery(userID)
		if userErr != nil {
			return nil, errors.New("forbidden")
		}

		for _, v := range roles {
			if v == user.Role {
				return next(ctx)
			}
		}

		return nil, errors.New("forbidden")
	}

	router := chi.NewRouter()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			if token == "" {
				ctx := context.WithValue(r.Context(), "userID", "")
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			userID, tokenErr := res.Queries.UserIDByAuthToken.ExecuteQuery(token)
			if tokenErr != nil {
				ctx := context.WithValue(r.Context(), "userID", "")
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), "userID", userID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	})

	srv.Use(&extension.ComplexityLimit{
		Func: func(ctx context.Context, rc *graphql.OperationContext) int {
			if rc != nil && rc.Operation.Name == "IntrospectionQuery" {
				return 200
			}
			return 100
		},
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/gql"))
	router.Handle("/gql", srv)

	return router
}
