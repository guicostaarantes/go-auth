package hash_util

type I interface {
	Hash(plain string) (string, error)
	Compare(plain string, hashed string) (bool, error)
}
