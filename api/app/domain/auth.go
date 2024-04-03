package domain

// --------------------------------
//      		AUTH
// --------------------------------

type (
	AuthRequest struct {
		AuthorizationCode string `json:"authorization_code,omitempty"`
		StateId           string `json:"state_id,omitempty"`
	}

	AuthResponse struct {
		Token        string `json:"token,omitempty"`
		ClientOrigin string `json:"client_origin,omitempty"`
	}
)

// --------------------------------
//         GET PROVIDER URI
// --------------------------------

type (
	GetProviderUriRequest struct {
		ClientOrigin string `json:"client_origin,omitempty"`
	}

	GetProviderUriResponse struct {
		Uri string `json:"uri,omitempty"`
	}
)
