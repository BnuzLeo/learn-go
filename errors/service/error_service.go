package service

import (
	pkErr "github.com/pkg/errors"
	"learn-go/errors/dao"
)

type ErrService struct {
}

func NewErrService() *ErrService {
	return &ErrService{}
}

func (s *ErrService) Errors() error {
	return dao.Errors()
}

func (s *ErrService) Wrap() error {
	return dao.Wrap()
}

func (s *ErrService) WithMessage() error {
	return dao.WithMessage()
}

func (s *ErrService) WithStack() error {
	return dao.WithStack()
}

func (s *ErrService) WithMessageAndStack() error {
	return pkErr.WithMessage(dao.WithStack(), "service reason: WithMessageAndStack")
}

func (s *ErrService) WithMessageAndWrap() error {
	return pkErr.WithMessage(dao.Wrap(), "service reason: WithMessageAndWrap")
}

func (s *ErrService) WithMessageAndWithMessage() error {
	return pkErr.WithMessage(dao.WithMessage(), "service reason: WithMessageAndWithMessage")
}

func (s *ErrService) WrapAndWrap() error {
	return pkErr.Wrap(dao.Wrap(), "service reason: WrapAndWrap")
}
