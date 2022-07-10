package users_workers

import (
	"encoding/json"

	users_models "github.com/guicostaarantes/go-auth/modules/users/models"
	log_util "github.com/guicostaarantes/go-auth/utils/log"
	store_util "github.com/guicostaarantes/go-auth/utils/store"
	stream_util "github.com/guicostaarantes/go-auth/utils/stream"
)

type UserCreatedWorker struct {
	LogUtil    log_util.I
	StoreUtil  store_util.I
	StreamUtil stream_util.I
}

func (w UserCreatedWorker) ExecuteWorker() {
	msgChannel, errChannel, unsubChannel := w.StreamUtil.Subscribe("users")

	go func() {
		for {
			select {
			case msg := <-msgChannel:
				user := &users_models.User{}

				err := json.Unmarshal(msg, user)
				if err != nil {
					w.LogUtil.Error("80f46331", err)
				}

				err = w.StoreUtil.Create(user)
				if err != nil {
					w.LogUtil.Error("b2ebde2d", err)
				}
			case err := <-errChannel:
				w.LogUtil.Error("6f181262", err)
				panic(err)
			case <-unsubChannel:
				w.LogUtil.Debug("f8d7e844", string("unsubscribed from UserCreated worker"))
				return
			}
		}
	}()
}
