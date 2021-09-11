package utils

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"shopping-cart/common"
)

// SdtClaims defines the custom claims
type SdtClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt_lib.StandardClaims
}

type Jwt struct {
	Token string
}

// parserJWT parser token to get customerId
func (u Jwt) GetCustomerIdByJWT() (string, error) {
	jwtToken, err := jwt_lib.ParseWithClaims(u.Token, &SdtClaims{}, func(token *jwt_lib.Token) (i interface{}, e error) {
		return []byte(common.Config.JwtSecretPassword), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := jwtToken.Claims.(*SdtClaims); ok && jwtToken.Valid {
		return claims.Name, nil
	}
	return "", nil
}
