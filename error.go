package expr

import (
	"errors"
)

var (
	ErrTooFewArguments   = errors.New("too few arguments")
	ErrWrongArgumentType = errors.New("wrong argument type")
)
