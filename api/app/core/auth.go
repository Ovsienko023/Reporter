package core

import (
	"context"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/domain/constants"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

// --------------------------------
//      		AUTH
// --------------------------------

// Auth возвращает следующие ошибки:
//
// errcore.ErrHostNotFound
// errcore.ErrInternal
func (c *Core) Auth(ctx context.Context, request domain.AuthRequest) (*domain.AuthResponse, error) {
	state, ok := c.Cache.AuthState[request.StateId]
	if !ok {
		//c.logger.Error("Failed to read state cache.")
		return nil, ErrInternal // state протух
	}

	tcsAuth, err := tcsAuthWithAuthorizationCode(ctx, tcsAuthWithAuthCode{
		Address:      c.Config.Tcs.Host,
		ClientId:     c.Config.Tcs.ClientId,
		ClientSecret: c.Config.Tcs.ClientSecret,
	})

	if err != nil {
		//c.logger.Error(fmt.Sprintf("TCS error, failed to auth with authorization code: %s", err.Error()))
		return nil, ErrInternal
	}

	profile, err := getTcsProfile(ctx, getTcsProfileRequest{
		Address:     c.Config.Tcs.Host,
		AccessToken: tcsAuth.AccessToken,
	})

	if err != nil {
		//c.logger.Error(fmt.Sprintf("TCS error, failed to get profile: %s", err.Error()))
		return nil, ErrInternal
	}

	avatar, err := getTcsAvatar(c.Config.Tcs.Host, profile.Login, tcsAuth.AccessToken)
	if err != nil {
		// todo log
	}

	auth, err := c.db.Auth(ctx, request.ToDb(
		profile.Login,
		profile.DisplayName,
		avatar,
	))
	if err != nil {
		//c.logger.Error(fmt.Sprintf("Database error, failed to Auth: %s", err.Error()))
		return nil, ErrInternal
	}

	claims := &CustomClaims{
		auth.UserId,
		state.ServerHost,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(tcsAuth.ExpiresIn) * time.Second).Unix(),
		},
	}

	signingKey := []byte(c.Config.Api.TokenSecret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		//c.logger.Error(fmt.Sprintf("Failed to get the complete, signed token: %s", err.Error()))
		return nil, ErrInternal
	}

	return new(domain.AuthResponse).From(signedToken, state.ClientOrigin), nil
}

// --------------------------------
//         GET PROVIDER URI
// --------------------------------

// GetProviderUri метод для получения uri необходимого для oauth авторизации
func (c *Core) GetProviderUri(ctx context.Context, request *domain.GetProviderUriRequest) (*domain.GetProviderUriResponse, error) {
	_, err := getTcsServerInfo(ctx, c.Config.Tcs.Host)
	if err != nil {
		return nil, ErrProviderServerNotAvailable
	}

	// Сохраняем в кэш поле state из oauth, он будет использован
	//после редиректа с авторизации tcs в эндпоинт "GET: {host}/api/v1/auth"
	stateId := uuid.New().String()
	tcsHost := c.Config.Tcs.Host
	tcsClientId := c.Config.Tcs.ClientId

	c.Cache.AuthState[stateId] = AuthState{
		ServerHost:   tcsHost,
		ClientOrigin: request.ClientOrigin,
	}

	uri := fmt.Sprintf(constants.Oauth2TcsFormatString, tcsHost, tcsClientId, stateId)

	return new(domain.GetProviderUriResponse).From(uri), nil
}
