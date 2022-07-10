package log_util

import "fmt"

type FmtImpl struct {}

func (u FmtImpl) Debug(uid string, message string) error {
	fmt.Printf("Log (Debug): %s (uid %s)\n", message, uid)
	return nil
}

func (u FmtImpl) Info(uid string, message string) error {
	fmt.Printf("Log (Info): %s (uid %s)\n", message, uid)
	return nil
}

func (u FmtImpl) Warn(uid string, message string) error {
	fmt.Printf("Log (Warn): %s (uid %s)\n", message, uid)
	return nil
}

func (u FmtImpl) Error(uid string, err error) error {
	fmt.Printf("Log (Error): %s (uid %s)\n", err.Error(), uid)
	return nil
}
