package users_commands

import (
	"fmt"

	"github.com/gofrs/uuid"
	users_models "github.com/guicostaarantes/go-auth/modules/users/models"
	hash_util "github.com/guicostaarantes/go-auth/utils/hash"
	log_util "github.com/guicostaarantes/go-auth/utils/log"
	store_util "github.com/guicostaarantes/go-auth/utils/store"
	token_store_util "github.com/guicostaarantes/go-auth/utils/token_store"
)

type AuthenticateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateUser struct {
	HashUtil hash_util.I
	LogUtil log_util.I
	StoreUtil store_util.I
	TokenStoreUtil token_store_util.I
}

func (c AuthenticateUser) ExecuteCommand(input AuthenticateUserInput) (string, error) {
	user := users_models.User{}

	err := c.StoreUtil.First("email", input.Email, &user)
	if err != nil {
		c.LogUtil.Error("2c9a97bc", err)
		return "", fmt.Errorf("internal server error")
	}

	if user.ID == "" {
		return "", fmt.Errorf("forbidden")
	}

	valid, err := c.HashUtil.Compare(input.Password, user.Password)
	if err != nil {
		c.LogUtil.Error("964ea795", err)
		return "", fmt.Errorf("internal server error")
	}

	if !valid {
		return "", fmt.Errorf("forbidden")
	}

	token := uuid.Must(uuid.NewV4()).String()

	err = c.TokenStoreUtil.Create(token, user.ID)
	if err != nil {
		c.LogUtil.Error("3b2e0b8a", err)
		return "", fmt.Errorf("internal server error")
	}

	return token, nil
}
