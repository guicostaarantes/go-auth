package stream_util

type I interface {
	Send(topic string, message []byte) error
	Subscribe(topic string) (chan []byte, chan error, chan bool)
}
