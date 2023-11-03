package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/medic-basic/auth/pkg/domain/model"
)

func ExtractToken(bearerToken string) string {
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	} else {
		return ""
	}
}

func VerifyAuthToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		//return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
		return []byte("AUTH_TOKEN_SECRET_KEY"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func CreateJWT(tokenParam model.CreateTokenParam) (*model.TokenDetailsParam, error) {
	td := &model.TokenDetailsParam{}
	var err error
	atClaims := jwt.MapClaims{}
	rtClaims := jwt.MapClaims{}

	td.AtExpires = time.Now().Add(time.Second * time.Duration(tokenParam.AccessTTL)).Unix()
	td.AccessUuid = GetHash(tokenParam.AccessKey)
	td.RtExpires = time.Now().Add(time.Second * time.Duration(tokenParam.RefreshTTL)).Unix()
	td.RefreshUuid = GetHash(tokenParam.RefreshKey)

	atClaims["uuid"] = td.AccessUuid
	atClaims["phone"] = tokenParam.UserInfo.Phone
	atClaims["role"] = tokenParam.UserInfo.Role
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	//td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET_KEY")))
	td.AccessToken, err = at.SignedString([]byte("AUTH_TOKEN_SECRET_KEY"))
	if err != nil {
		return nil, err
	}

	rtClaims["uuid"] = td.RefreshUuid
	rtClaims["phone"] = tokenParam.UserInfo.Phone
	rtClaims["role"] = tokenParam.UserInfo.Role
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	//td.AccessToken, err = at.SignedString([]byte(os.Getenv("REFRESH_SECRET_KEY")))
	td.RefreshToken, err = rt.SignedString([]byte("AUTH_TOKEN_SECRET_KEY"))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func GetJWTClaims(value any) (*model.TokenClaimsParam, error) {
	token := value.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	if claims["phone"] != nil && claims["role"] != nil && claims["uuid"] != nil {
		return &model.TokenClaimsParam{
			Phone: claims["phone"].(string),
			Role:  claims["role"].(string),
			Uuid:  claims["uuid"].(string),
		}, nil
	} else {
		return nil, errors.New("not enough claim data in token")
	}
}
