package userToken

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	UserID      int    `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	UserStatus  string `json:"user_status"`
	jwt.StandardClaims
}

var SecretUserKey = []byte("das#jd!ah^wr$we$ry$wbw_w^#$sa)adEd&$sda23*Hgaas2!5")

func GenerateUserToken(userID int, phoneNumber, userStatus string) (string, error) {
	userClaims := UserClaims{
		UserID:      userID,
		PhoneNumber: phoneNumber,
		UserStatus:  userStatus,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString(SecretUserKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
