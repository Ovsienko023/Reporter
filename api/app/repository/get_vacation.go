package repository

import (
	"context"
	"time"
)

const sqlGetVacation = `
	select id,
		   date_from,
		   date_to,
		   is_paid,
		   status,
		   description,
		   creator_id,
		   created_at,
		   updated_at,
		   payload
    from main.vacations as v
    where id = $1 and exists(select 1
        					 from main.permissions_users_to_objects as pto
            				 where pto.user_id = $2 and
       							   pto.object_id =  v.creator_id)`

func (c *Client) GetVacation(ctx context.Context, msg *GetVacation) (*Vacation, error) {
	isAuth, err := c.checkUserPermission(ctx, msg.InvokerId, msg.InvokerId)
	if !isAuth {
		return nil, ErrPermissionDenied
	}

	row, err := c.driver.Query(ctx,
		sqlGetVacation,
		msg.VacationId,
		msg.InvokerId,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	vacation := &Vacation{}

	for row.Next() {
		err := row.Scan(
			&vacation.Id,
			&vacation.DateFrom,
			&vacation.DateTo,
			&vacation.IsPaid,
			&vacation.Status,
			&vacation.Description,
			&vacation.CreatorId,
			&vacation.CreatedAt,
			&vacation.UpdatedAt,
			&vacation.Payload,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	if vacation.Id == nil {
		return nil, ErrVacationIdNotFound
	}

	return vacation, nil
}

type GetVacation struct {
	InvokerId  string `json:"invoker_id,omitempty"`
	VacationId string `json:"vacation_id,omitempty"`
}

type Vacation struct {
	Id          *string          `json:"id,omitempty"`
	DateFrom    *time.Time       `json:"date_from,omitempty"`
	DateTo      *time.Time       `json:"date_to,omitempty"`
	IsPaid      *bool            `json:"is_paid,omitempty"`
	Status      *string          `json:"status,omitempty"`
	Description *string          `json:"description,omitempty"`
	CreatorId   *string          `json:"creator_id,omitempty"`
	CreatedAt   *time.Time       `json:"created_at,omitempty"`
	UpdatedAt   *time.Time       `json:"updated_at,omitempty"`
	Payload     *VacationPayload `json:"payload,omitempty"`
}

type VacationPayload struct{}