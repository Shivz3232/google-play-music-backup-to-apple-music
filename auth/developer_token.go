package auth

import (
	"os"
	"shivu/google-play-music-backup-reader/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenrateDeveloperToken(obj *models.DeveloperToken, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"iss": obj.TeamID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = obj.KeyID

	privateKeyStr, err := os.ReadFile(obj.PrivateKeyPath)
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyStr)
	if err != nil {
		return "", err
	}

	return token.SignedString(privateKey)
}
