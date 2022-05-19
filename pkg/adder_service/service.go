package adder_service

import (
	"context"
	"github.com/rs/zerolog"
)

type Service struct {
	log zerolog.Logger
}

func NewService(log zerolog.Logger) (svc *Service) {
	return &Service{log: log}
}

func (s *Service) SumTwoNumbers(ctx context.Context, first int, second int) (result int, err error) {
	result = first + second
	return
}
