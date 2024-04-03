package domain

import "github.com/Ovsienko023/reporter/app/repository"

// --------------------------------
//      		AUTH
// --------------------------------

func (r AuthRequest) ToDb(login string, displayName *string, avatar []byte) repository.Auth {
	return repository.Auth{
		Login:       login,
		DisplayName: displayName,
		Avatar:      avatar,
	}
}

func (r *AuthResponse) From(token string, clientOrigin string) *AuthResponse {
	if r == nil {
		return nil
	}

	r.Token = token
	r.ClientOrigin = clientOrigin

	return r
}

// --------------------------------
//         GET PROVIDER URI
// --------------------------------

func (res *GetProviderUriResponse) From(uri string) *GetProviderUriResponse {
	if res == nil {
		return nil
	}

	res.Uri = uri

	return res
}
