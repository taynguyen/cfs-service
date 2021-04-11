package main

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/k0kubun/pp"
	"gitlab.com/cfs-service/utils"
)

func main() {
	// dateTime, err := gostradamus.Parse("2020-11-25 07:36:04.000001", "YYYY-MM-DD HH:mm:ss.S")
	// if err != nil {
	// 	panic(err)
	// }

	// pp.Println("dateTime:", dateTime.GoString())

	// layout := "2006-01-02 15:04:05.000"
	// d, e := time.Parse(layout, "2020-11-25 07:36:04.193")
	// // time.Parse("", "")
	// rs := time.Now().Format("2006-01-02 15:04:05.000")
	// pp.Println("rs:", rs)

	keyPath := "/Users/taynguyen/Projects/Personal/cfs-service/dev/jwtRS256.key"
	signBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		pp.Println("err:", err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		pp.Println("err:", err)
	}

	privateKey := signKey
	token, err := CreateToken(privateKey)
	pp.Println("Token:", token, " err:", err)

}

// CreateToken create jwt token
func CreateToken(privateKey *rsa.PrivateKey) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["a_id"] = "4f9b99eb-490a-484e-bade-15e3841dfda9"
	atClaims["c"] = utils.TimeUnixMilli(time.Now())
	atClaims["exp"] = utils.TimeUnixMilli(time.Now().Add(time.Hour * 24 * 30))

	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	token, err := at.SignedString(privateKey)

	if err != nil {
		return "", err
	}
	return token, nil
}
