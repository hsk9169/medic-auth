package auth

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/medic-basic/auth/pkg/domain/model"
	"github.com/medic-basic/auth/pkg/dto/http"
	"github.com/medic-basic/auth/pkg/external/http/basic"
	"github.com/medic-basic/auth/pkg/util"
)

func (a *Aggregate) CheckTokenValid(ctx context.Context, param model.TokenClaimsParam) error {
	// Check cache hit in case of user signout state
	fmt.Println(param.Uuid)
	hitData, err := a.Cache.Get(ctx, param.Uuid, new(string))
	fmt.Println(hitData)
	if err != nil {
		fmt.Println("check token cache hit error")
		return err
	}
	return nil
}

func (a *Aggregate) VerifyUser(ctx context.Context, param model.UserVerifyParam) (any, error) {
	// Validate input
	if err := validator.New().Struct(param); err != nil {
		return nil, err
	}

	account, err := basic.GetAccountInfo(http.GetAccountInfoReqToBasic{
		UserID: util.GetHash(param.Phone),
		Role:   param.Role,
	})
	if err != nil {
		fmt.Println(err)
		fmt.Println("error get account info")
		return nil, err
	}
	fmt.Println(account)

	return account, nil
}
