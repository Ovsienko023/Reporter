package repository

import (
	"context"
)

const sqlSignIn = `
	select u.id,
           hash
    from main.users as u
         inner join main.user_passwords on u.id = user_passwords.user_id
         inner join main.user_logins on user_passwords.id = user_logins.grant_id
    where login = $1 and
          u.deleted_at is null`

func (c *Client) SignIn(ctx context.Context, msg *SignIn) (*SignInResponse, error) {
	row, err := c.driver.Query(ctx, sqlSignIn, msg.Login)
	if err != nil {
		return nil, NewInternalError(err)
	}

	auth := &SignInResponse{}

	for row.Next() {
		err := row.Scan(
			&auth.UserId,
			&auth.Password,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	if auth.Password == nil {
		return nil, ErrCredentials
	}

	return auth, nil
}

type SignIn struct {
	Login string `json:"login,omitempty"`
}

type SignInResponse struct {
	UserId   *string `json:"id,omitempty"`
	Password *string `json:"password,omitempty"`
}
