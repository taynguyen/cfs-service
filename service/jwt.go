package service

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"gitlab.com/cfs-service/utils"
)

type TokenMetadata struct {
	AgencyID string

	CreatedAt time.Time
	ExpiredAt time.Time
}

const (
	// TokenTypeGenerall for generral query, get, post data
	TokenTypeGenerall = int16(1)

	// TokenTypeEmailConfirm indicate that token is used for confirming email
	TokenTypeEmailConfirm = int16(2)

	// TokenTypePasswordReset token to reset password request
	TokenTypePasswordReset = int16(3)

	// TokenGetPrivateKey used for getting user private key at other service
	TokenGetPrivateKey = int16(4)
)

var (
	publicKey *rsa.PublicKey
)

func InitializeKeyService(publicKeyPath string) error {
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return errors.Wrap(err, "Read public key file failed. ")
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return errors.Wrap(err, "Parse public key failed. ")
	}
	publicKey = verifyKey

	return nil
}

func ExtractToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, err
	}

	return token, nil
}

const JWTSeperator = "."

func ValidateByPublicKey(tokenString string) bool {
	parts := strings.Split(tokenString, JWTSeperator)
	if len(parts) < 3 {
		return false
	}
	err := jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], publicKey)
	return err == nil
}

func ExtractTokenMetadata(tokenString string) (*TokenMetadata, error) {
	token, err := ExtractToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		agencyID := fmt.Sprintf("%s", claims["a_id"])

		createdAtMls, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["c"]), 10, 64)
		if err != nil {
			return nil, err
		}

		expiredAtMls, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["exp"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &TokenMetadata{
			AgencyID:  agencyID,
			CreatedAt: utils.TimeFromUnixMillis(createdAtMls),
			ExpiredAt: utils.TimeFromUnixMillis(expiredAtMls),
		}, nil
	}
	return nil, err
}
