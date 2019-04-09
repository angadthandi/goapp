package auth

import (
	"errors"
	"time"

	log "github.com/angadthandi/goapp/log"
	"github.com/angadthandi/goapp/util"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(jwtAuthSecret string, userID int) (UserToken, error) {
	var ret UserToken

	expiresAt := time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix()

	// Create a new token object,
	// specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userid": userID,
			"nbf":    expiresAt,
		},
	)

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := token.SignedString(
		[]byte(jwtAuthSecret))
	if err != nil {
		return ret, err
	}

	ret.Token = tokenStr
	ret.ExpiresAt = int(expiresAt)

	return ret, err
}

func ParseToken(jwtAuthSecret, tokenStr string) (int, error) {
	var (
		userID int
		err    error
	)

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys
	// for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use,
	// but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenStr,
		func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Errorf("Unexpected signing method: %v",
					token.Header["alg"])
				return nil,
					errors.New("Invalid signing method!")
			}

			// return []byte containing your secret,
			// e.g. []byte("my_secret_key")
			return []byte(jwtAuthSecret), nil
		},
	)
	if err != nil {
		log.Errorf("Parse token error: %v", err)
		return userID, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		log.Errorf("Parse token error: %v", err)
		return userID, err
	}

	// fmt.Println(claims["userid"], claims["nbf"])
	userID, err = util.InterfaceToInt(claims["userid"])
	if err != nil {
		log.Errorf("Parse token: invalid userID type error: %v",
			err)
		return userID,
			errors.New("Parse token error: Invalid userID type")
	}

	return userID, err
}
