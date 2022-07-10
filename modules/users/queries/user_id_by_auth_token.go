package users_queries

import (
	"errors"

	log_util "github.com/guicostaarantes/go-auth/utils/log"
	token_store_util "github.com/guicostaarantes/go-auth/utils/token_store"
)

type UserIDByAuthToken struct {
	LogUtil        log_util.I
	TokenStoreUtil token_store_util.I
}

func (q UserIDByAuthToken) ExecuteQuery(token string) (string, error) {
	ID, err := q.TokenStoreUtil.First(token)
	if err != nil {
		q.LogUtil.Error("57d84202", err)
		return "", errors.New("internal server error")
	}

	return ID, nil
}
