package token_store_util

type I interface {
	First(token string)	(string, error)
	Create(token string, userID string) error
}
