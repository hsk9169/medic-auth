package model

import (
	"context"
)

type AuthAggregate interface {
	VerifyUser(ctx context.Context, param UserVerifyParam) (any, error)
	CreateToken(ctx context.Context, param UserVerifyParam) (*TokenDetailsParam, error)
	CheckTokenValid(ctx context.Context, param TokenClaimsParam) error
}

type UserVerifyParam struct {
	Phone string `validate:"required"`
	Role  string `validate:"required"`
}

type TokenVerifyParam struct {
	AccessToken  string `validate:"required"`
	RefreshToken string `validate:"required"`
}

type CreateTokenParam struct {
	AccessKey  string          `validate:"required"`
	RefreshKey string          `validate:"required"`
	AccessTTL  int64           `validate:"required"`
	RefreshTTL int64           `validate:"required"`
	UserInfo   UserVerifyParam `validate:"required"`
}

type TokenDetailsParam struct {
	AccessToken  string `validate:"required"`
	RefreshToken string `validate:"required"`
	AccessUuid   string `validate:"required"`
	RefreshUuid  string `validate:"required"`
	AtExpires    int64  `validate:"required"`
	RtExpires    int64  `validate:"required"`
}

type TokenClaimsParam struct {
	Phone string `validate:"required"`
	Role  string `validate:"required"`
	Uuid  string `validate:"required"`
}
