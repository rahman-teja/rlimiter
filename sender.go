package rlimiter

type Sender interface {
	Send(msg interface{})
}
