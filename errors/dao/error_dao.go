package dao

import (
	"errors"
	pkErr "github.com/pkg/errors"
)

type User struct {
	Name string
	Sex  int
}

func Errors() error {
	return errors.New("not found")
}

func Wrap() error {
	return pkErr.Wrap(errors.New("not found"), "dao reason: Wrap")
}

func WithStack() error {
	return pkErr.WithStack(errors.New("not found"))
}

func WithMessage() error {
	return pkErr.WithMessage(errors.New("not found"), "dao reason: WithMessage")
}
