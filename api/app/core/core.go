package core

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Core struct {
	db     repository.InterfaceDatabase
	Cache  *Cache
	Config configuration.Config
}

type CustomClaims struct {
	UserId     string `json:"user_id"`
	ServerHost string `json:"server_host"`
	jwt.StandardClaims
}

func NewCore(cfg configuration.Config, db repository.InterfaceDatabase) *Core {
	cache := newCache()
	go cache.Clean()

	return &Core{
		db:     db,
		Cache:  newCache(),
		Config: cfg,
	}
}

// authorize возвращает InvokerId или ошибку:
// ErrUnauthorized
func (c *Core) authorize(token string) (string, error) {
	if token == "" {
		return "", ErrUnauthorized
	}

	token = strings.Replace(token, "Bearer ", "", 1)

	data, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("SecretKey"), nil // todo secret into config
	})
	if err != nil {
		return "", ErrUnauthorized
	}

	if _, ok := data.Method.(*jwt.SigningMethodHMAC); !ok {
		return "", ErrUnauthorized // signing method: token.Header["alg"]
	}

	if claims, ok := data.Claims.(jwt.MapClaims); ok && data.Valid {
		return claims["iss"].(string), nil
	} else {
		return "", ErrUnauthorized
	}
}

func (c *Core) generateHash(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, 13)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

func (c *Core) checkPassword(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
