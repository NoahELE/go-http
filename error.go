package main

import "errors"

var (
	ErrNoArg      = errors.New("no args provided")
	ErrInvalidArg = errors.New("invalid arg")
)
