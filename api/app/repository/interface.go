package repository

import (
	"context"
)

type InterfaceDatabase interface {
	Auth(ctx context.Context, msg Auth) (*Authed, error)
	SignUp(ctx context.Context, msg *SignUp) error
	SignIn(ctx context.Context, msg *SignIn) (*SignInResponse, error)

	GetProfile(ctx context.Context, msg *GetProfile) (*Profile, error)
	UpdateProfile(ctx context.Context, msg *UpdateProfile) error

	GetUsers(ctx context.Context, msg *GetUsers) ([]UserItem, *int, error)

	GetCalendarEvents(ctx context.Context, msg *GetCalendarEvents) ([]CalendarEvent, *int64, error)

	GetDayOff(ctx context.Context, msg *GetDayOff) (*DayOff, error)
	CreateDayOff(ctx context.Context, msg *CreateDayOff) (*CreatedDayOff, error)
	DeleteDayOff(ctx context.Context, msg *DeleteDayOff) error

	GetSickLeave(ctx context.Context, msg *GetSickLeave) (*SickLeave, error)
	CreateSickLeave(ctx context.Context, msg *CreateSickLeave) (*CreatedSickLeave, error)
	DeleteSickLeave(ctx context.Context, msg *DeleteSickLeave) error

	GetVacationPaid(ctx context.Context, msg *GetVacationPaid) (*VacationPaid, error)
	CreateVacationPaid(ctx context.Context, msg *CreateVacationPaid) (*CreatedVacationPaid, error)
	DeleteVacationPaid(ctx context.Context, msg *DeleteVacationPaid) error

	GetVacationUnpaid(ctx context.Context, msg *GetVacationUnpaid) (*VacationUnpaid, error)
	CreateVacationUnpaid(ctx context.Context, msg *CreateVacationUnpaid) (*CreatedVacationUnpaid, error)
	DeleteVacationUnpaid(ctx context.Context, msg *DeleteVacationUnpaid) error

	GetReport(ctx context.Context, msg *GetReport) (*Report, error)
	GetReports(ctx context.Context, msg *GetReports) ([]ReportItem, *int, error)
	CreateReport(ctx context.Context, msg *CreateReport) (*CreatedReport, error)
	UpdateReport(ctx context.Context, msg *UpdateReport) error
	DeleteReport(ctx context.Context, msg *DeleteReport) error
	GetSystemUser(ctx context.Context) (*SystemUser, error)

	GetStatistics(ctx context.Context, msg *GetStatistics) (*Statistics, error)

	AddObjectToUserPermission(ctx context.Context, msg *AddObjectToUserPermission) error
	RemoveObjectFromUserPermission(ctx context.Context, msg *RemoveObjectFromUserPermission) error
}
