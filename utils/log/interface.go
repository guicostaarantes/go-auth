package log_util

type I interface {
	Debug(uid string, message string) error
	Info(uid string, message string) error
	Warn(uid string, message string) error
	Error(uid string, err error) error
}
