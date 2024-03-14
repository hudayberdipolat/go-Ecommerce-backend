package adminToken

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AdminClaims struct {
	AdminID     int    `json:"admin_id"`
	PhoneNumber string `json:"phone_number"`
	AdminStatus string `json:"admin_status"`
	jwt.StandardClaims
}

var SecretAdminKey = []byte("das#jd!ah^wr$we$ry$wbw_w^#$sa)adEd&$sda23*Hgaas2!5dfsdf344342!@3424")

func GenerateAdminToken(adminID int, phoneNumber, adminStatus string) (string, error) {
	adminClaims := AdminClaims{
		AdminID:     adminID,
		PhoneNumber: phoneNumber,
		AdminStatus: adminStatus,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, adminClaims)
	tokenString, err := token.SignedString(SecretAdminKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
