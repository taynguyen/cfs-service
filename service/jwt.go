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
	cfsservice "gitlab.com/cfs-service"
	"gitlab.com/cfs-service/utils"
)

type TokenMetadata struct {
	UserID int64

	// Email should have value when token type is Reset Password
	Email string // Optional

	UserIP    string
	StatusID  int32
	TypeID    int16
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

func InitializeJWT(cf *cfsservice.RuntimeConfig) error {
	verifyBytes, err := ioutil.ReadFile(cf.JWTPublicKeyPath)
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
		email := fmt.Sprintf("%s", claims["email"])

		userid, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["userid"]), 10, 64)
		if err != nil {
			return nil, err
		}

		stt, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["stt"]), 10, 32)
		if err != nil {
			return nil, err
		}

		tokenType, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["type"]), 10, 16)
		if err != nil {
			return nil, err
		}

		createdAtMls, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["c"]), 10, 64)
		if err != nil {
			return nil, err
		}

		expiredAtMls, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["exp"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &TokenMetadata{
			UserID:    userid,
			Email:     email,
			StatusID:  int32(stt),
			TypeID:    int16(tokenType),
			UserIP:    fmt.Sprintf("%v", claims["ip"]),
			CreatedAt: utils.TimeFromUnixMillis(createdAtMls),
			ExpiredAt: utils.TimeFromUnixMillis(expiredAtMls),
		}, nil
	}
	return nil, err
}
