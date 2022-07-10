package store_util

type I interface {
	First(field string, value string, entity interface{}) error	
	Create(entity interface{}) error
}
