package uid_util

import "github.com/gofrs/uuid"

type UUIDv4Impl struct {}

func (u UUIDv4Impl) Generate() string {
  return uuid.Must(uuid.NewV4()).String()
}
