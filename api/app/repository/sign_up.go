package repository

import (
	"context"
	"fmt"
)

const sqlSignUp = `
	select error
	from main.create_user(
	    _invoker_id := $1, 
	    _login := $2, 
	    _password := $3, 
	    _display_name := $4
	);`

// SignUp возвращает следующие ошибки:
// ErrInternal
// ErrUnauthorized
// ErrLoginAlreadyInUse
func (c *Client) SignUp(ctx context.Context, msg *SignUp) error {
	raw, err := c.driver.Query(ctx, sqlSignUp,
		msg.InvokerId,
		msg.Login,
		msg.Password,
		msg.DisplayName,
	)
	if err != nil {
		return NewInternalError(err)
	}

	if !raw.Next() {
		if err = raw.Err(); err != nil {
			return AnalyzeRowsError(err)
		}
	}

	var queryErr []byte

	err = raw.Scan(&queryErr)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}

	if queryErr != nil {
		if err = AnalyzeError(queryErr); err != nil {
			return err
		}
	}

	return nil
}

type SignUp struct {
	InvokerId   string  `json:"invoker_id,omitempty"`
	Login       string  `json:"login,omitempty"`
	Password    string  `json:"password,omitempty"`
	DisplayName *string `json:"display_name,"`
}
