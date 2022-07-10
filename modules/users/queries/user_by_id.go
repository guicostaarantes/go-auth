package users_queries

import (
	"errors"

	users_models "github.com/guicostaarantes/go-auth/modules/users/models"
	log_util "github.com/guicostaarantes/go-auth/utils/log"
	store_util "github.com/guicostaarantes/go-auth/utils/store"
)

type UserByID struct {
	LogUtil   log_util.I
	StoreUtil store_util.I
}

func (q UserByID) ExecuteQuery(ID string) (*users_models.User, error) {
	user := users_models.User{}

	err := q.StoreUtil.First("id", ID, &user)
	if err != nil {
		q.LogUtil.Error("3d3313d4", err)
		return nil, errors.New("internal server error")
	}

	return &user, nil
}
