package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type Auth struct {
	Login       string  `json:"call_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Avatar      []byte  `json:"avatar,omitempty"`
}

func (c *Client) Auth(ctx context.Context, msg Auth) (*Authed, error) {
	if userId, err := c.foundUser(ctx, msg.Login); err != nil {
		return &Authed{
			UserId: userId,
		}, nil
	}

	systemUser, err := c.GetSystemUser(ctx)
	if err != nil {
		return nil, NewInternalError(err)
	}

	userId, err := c.createUser(ctx, createUser{
		InvokerId:   *systemUser.UserId,
		Login:       msg.Login,
		DisplayName: msg.DisplayName,
	})

	return &Authed{
		UserId: userId,
	}, nil
}

type Authed struct {
	UserId string `json:"user_id,omitempty"`
}

const sqlFoundUser = `
	select users.id
	from main.users
	inner join main.user_logins ul on ul.user_id = users.id
	where ul.login = $1
`

func (c *Client) foundUser(ctx context.Context, login string) (string, error) {
	row, err := c.driver.Query(ctx, sqlFoundUser, login)
	if err != nil {
		return "", NewInternalError(err)
	}

	var userId string

	for row.Next() {
		err := row.Scan(
			&userId,
		)
		if err != nil {
			return "", NewInternalError(err)
		}
	}

	if userId == "" {
		return "", ErrUserIdNotFound
	}

	return userId, nil
}

const sqlCreateUser = `
    INSERT INTO main.users
        (creator_id, display_name)
    VALUES
    ($1, $2)
    RETURNING id
`

const sqlCreateUserLogin = `
    INSERT INTO main.user_logins
        (creator_id, login)
    VALUES ($1, $2)
`

type createUser struct {
	InvokerId   string  `json:"invoker_id,omitempty"`
	Login       string  `json:"call_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}

func (c *Client) createUser(ctx context.Context, msg createUser) (string, error) {
	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrInternal, err)
	}

	row, err := transaction.Query(ctx, sqlCreateUser,
		msg.InvokerId,
		msg.DisplayName,
	)
	if err != nil {
		return "", NewInternalError(err)
	}

	var userId string

	for row.Next() {
		err = row.Scan(
			&userId,
		)
		if err != nil {
			return "", NewInternalError(err)
		}
	}

	_ = transaction.Commit(ctx)

	row, err = transaction.Query(ctx, sqlCreateUserLogin,
		msg.InvokerId,
		msg.Login,
	)

	if err != nil {
		return "", NewInternalError(err)
	}

	row.Next()
	status := row.CommandTag()
	if status != nil && !status.Insert() {
		return "", NewInternalError(err)
	}

	return userId, nil
}
