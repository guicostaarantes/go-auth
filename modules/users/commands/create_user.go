package users_commands

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	users_models "github.com/guicostaarantes/go-auth/modules/users/models"
	hash_util "github.com/guicostaarantes/go-auth/utils/hash"
	log_util "github.com/guicostaarantes/go-auth/utils/log"
	store_util "github.com/guicostaarantes/go-auth/utils/store"
	stream_util "github.com/guicostaarantes/go-auth/utils/stream"
)

type CreateUserInput struct {
	Active   bool              `json:"active"`
	Email    string            `json:"email"`
	Password string            `json:"password"`
	Role     users_models.Role `json:"role"`
}

type CreateUser struct {
	HashUtil   hash_util.I
	LogUtil    log_util.I
	StreamUtil stream_util.I
	StoreUtil  store_util.I
}

func (c CreateUser) ExecuteCommand(input CreateUserInput) (bool, error) {
	user := users_models.User{}

	err := c.StoreUtil.First("email", input.Email, &user)
	if err != nil {
		c.LogUtil.Error("4b3f31a2", err)
		return false, fmt.Errorf("internal server error")
	}

	if user.ID != "" {
		return false, fmt.Errorf("user with same email already exists")
	}

	hashedPassword, err := c.HashUtil.Hash(input.Password)
	if err != nil {
		c.LogUtil.Error("fc168033", err)
		return false, fmt.Errorf("internal server error")
	}

	user.ID = uuid.Must(uuid.NewV4()).String()
	user.Active = true
	user.Email = input.Email
	user.Password = hashedPassword
	user.Role = input.Role

	bytes, err := json.Marshal(user)
	if err != nil {
		c.LogUtil.Error("3b2e0b8a", err)
		return false, fmt.Errorf("internal server error")
	}

	err = c.StreamUtil.Send("users", bytes)
	if err != nil {
		c.LogUtil.Error("da8f3a4b", err)
		return false, fmt.Errorf("internal server error")
	}

	return true, nil
}
