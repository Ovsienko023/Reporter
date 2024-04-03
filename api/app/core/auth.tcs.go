package core

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
	"strings"
)

type tcsAuthWithAuthCode struct {
	Address      string `json:"address,omitempty"`
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Code         string `json:"code,omitempty"`
	//RedirectUri  string `json:"redirect_uri,omitempty"`
}

func tcsAuthWithAuthorizationCode(ctx context.Context, msg tcsAuthWithAuthCode) (*tcsAuth, error) {
	urlAuth := "https://%s/oauth2/v1/token"
	urls := fmt.Sprintf(urlAuth, msg.Address)

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", msg.ClientId)
	data.Set("client_secret", msg.ClientSecret)
	data.Set("code", msg.Code)
	//data.Set("redirect_uri", msg.RedirectUri)
	encodeData := data.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, urls, strings.NewReader(encodeData))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "Basic XXXX")

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}
	var auth tcsAuth

	err = dec.Decode(&auth)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}

	return &auth, nil
}

type tcsAuth struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	Scope        string `json:"scope,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func getTcsAvatar(host, userLogin, token string) ([]byte, error) {
	// Инициализируем http клиент todo добавить поддержку tls
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	url := fmt.Sprintf(
		"https://%s/api/v4/users/%s/avatar?access_token=%s",
		host,
		userLogin,
		token,
	)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	_ = resp.Body.Close()

	errData := struct {
		Error any `json:"error"`
	}{}

	if err := json.Unmarshal(buf.Bytes(), &errData); err != nil {
		return buf.Bytes(), nil
	}

	return buf.Bytes(), nil
}

func getTcsServerInfo(ctx context.Context, address string) (*tcsServerInfo, error) {
	urlGetServer := "https://%s/api/v3.5/server"

	url := fmt.Sprintf(urlGetServer, address)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	request.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}

	var server *tcsServerInfo

	err = dec.Decode(&server)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}

	return server, nil
}

type tcsServerInfo struct {
	Name        *string `json:"name,omitempty,omitempty"`
	ProductName *string `json:"product_name,omitempty,omitempty"`
	Version     *string `json:"version,omitempty,omitempty"`
}

type getTcsProfileRequest struct {
	Address     string `json:"address,omitempty"`
	AccessToken string `url:"access_token,omitempty"`
}

func getTcsProfile(ctx context.Context, msg getTcsProfileRequest) (*tcsProfile, error) {
	urlGetProfile := "https://%s/api/v3.5/me"
	url := fmt.Sprintf(urlGetProfile, msg.Address)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	request.Header.Add("Content-Type", "application/json")

	v, err := query.Values(msg)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}
	request.URL.RawQuery = v.Encode()

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}

	profile := struct {
		User struct {
			Id          *string `json:"id,omitempty"`
			Uid         *string `json:"uid,omitempty"`
			LoginName   *string `json:"login_name,omitempty"`
			DisplayName *string `json:"display_name,omitempty"`
			Avatar      *string `json:"avatar,omitempty"`
		}
	}{}

	err = dec.Decode(&profile)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err.Error())
	}

	if profile.User.Id == nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, "login not found")
	}

	return &tcsProfile{
		Login:       *profile.User.Id,
		DisplayName: profile.User.DisplayName,
	}, nil
}

type tcsProfile struct {
	Login       string  `json:"id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}
