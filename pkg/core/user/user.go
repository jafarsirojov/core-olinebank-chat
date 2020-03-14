package user

import (
	"context"
	"github.com/jafarsirojov/core-olinebank-chat/pkg/core/token"
    "github.com/jafarsirojov/mux/pkg/mux/middleware/jwt"
	"errors"
)

type Service struct { }

func NewService() *Service {
	return &Service{}
}

type ResponseDTO struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
}

func (s *Service) Profile(ctx context.Context) (response ResponseDTO, err error) {
	auth, ok := jwt.FromContext(ctx).(*token.Payload)
	if !ok {
		return ResponseDTO{}, errors.New("...")
	}

	return ResponseDTO{
		Id: auth.Id,
		Name: "Vasya",
		Avatar: "https://i.pravatar.cc/50",
	}, nil
}
