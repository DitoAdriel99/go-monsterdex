package jwt_parse

import (
	"fmt"
	"log"
	"strings"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"

	"github.com/golang-jwt/jwt/v4"
)

func GetClaimsFromToken(param string) (*entity.Claims, error) {
	tokenString := strings.TrimPrefix(param, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &entity.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return entity.JWTKEY, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Error parsing token: %s", err)
	}

	if claims, ok := token.Claims.(*entity.Claims); ok && token.Valid {
		return claims, nil
	} else {
		log.Println("Token is Invalid")
		return nil, fmt.Errorf("Invalid token")
	}
}
