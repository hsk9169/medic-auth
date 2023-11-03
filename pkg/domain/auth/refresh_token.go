package auth

import (
	"context"
	"time"

	"github.com/medic-basic/auth/pkg/domain/model"
	"github.com/medic-basic/auth/pkg/util"
)

const (
	AccessTokenTTL  = 60 * 5           // 5 minutes
	RefreshTokenTTL = 60 * 60 * 24 * 7 // 1 week
	AccessPrefix    = "access_"
	RefreshPrefix   = "refresh_"
)

func (a *Aggregate) CreateToken(ctx context.Context, param model.UserVerifyParam) (*model.TokenDetailsParam, error) {
	var td *model.TokenDetailsParam
	// need to hash uuid
	accessKey := AccessPrefix + param.Phone + param.Role
	refreshKey := RefreshPrefix + param.Phone + param.Role

	td, err := util.CreateJWT(model.CreateTokenParam{
		AccessKey:  accessKey,
		RefreshKey: refreshKey,
		AccessTTL:  AccessTokenTTL,
		RefreshTTL: RefreshTokenTTL,
		UserInfo:   param,
	})
	if err != nil {
		return nil, err
	}

	if err = a.Cache.Set(ctx, td.AccessUuid, &td.AccessToken, AccessTokenTTL*time.Second); err != nil {
		return nil, err
	}
	if err = a.Cache.Set(ctx, td.RefreshUuid, &td.RefreshToken, RefreshTokenTTL*time.Second); err != nil {
		return nil, err
	}

	return td, nil
}
