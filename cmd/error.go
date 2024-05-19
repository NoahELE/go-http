package cmd

import "errors"

var (
	errNoMethod   = errors.New("no method provided")
	errInvalidArg = errors.New("invalid arg")
)
