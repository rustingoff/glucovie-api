package jwtoken

import (
	"fmt"
	"glucovie/internal/models"
	"glucovie/pkg/dotenv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	accesstokenTTL  = 8 * time.Hour  // 8 hours
	refreshtokenTTL = 24 * time.Hour // 1 day
)

var signInKey = dotenv.GetEnvironmentVariable("TOKEN_SECRET_KEY")
var refreshSignKey = dotenv.GetEnvironmentVariable("REFRESH_TOKEN_KEY")

type AuthTokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func GenerateToken(user models.User) (*AuthTokenDetails, error) {
	var err error
	atd := &AuthTokenDetails{}
	atd.AtExpires = time.Now().Add(accesstokenTTL).Unix()
	atd.AccessUuid = uuid.NewV4().String()

	atd.RtExpires = time.Now().Add(refreshtokenTTL).Unix()
	atd.RefreshUuid = uuid.NewV4().String()

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["authorized"] = true
	accessTokenClaims["access_uuid"] = atd.AccessUuid
	accessTokenClaims["sub"] = user.ID.Hex()
	accessTokenClaims["exp"] = atd.AtExpires

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	atd.AccessToken, err = accessToken.SignedString([]byte(signInKey))
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["refresh_uuid"] = atd.RefreshUuid
	refreshTokenClaims["sub"] = user.ID.Hex()
	refreshTokenClaims["exp"] = atd.RtExpires

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	atd.RefreshToken, err = refreshToken.SignedString([]byte(refreshSignKey))
	if err != nil {
		return nil, err
	}

	return atd, nil
}

func DoPasswordsMatch(hashedPassword, currPassword string) bool {

	hashed := []byte(hashedPassword)
	current := []byte(currPassword)

	err := bcrypt.CompareHashAndPassword(
		hashed, current)
	return err == nil
}

func GenerateHashPassword(password string) string {

	var passwordBytes = []byte(password)

	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err.Error())
	}

	return string(hashedPasswordBytes)
}
