package rlimiter

import "errors"

var (
	ErrSenderRequired error = errors.New("rlimiter: Sender is required")
)
