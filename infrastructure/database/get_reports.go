package database

import (
	"context"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
	"time"
)

const sqlGetReports = `
	select id,
	       title,
		   date,
		   start_time,
		   end_time,
		   break_time,
		   work_time,
		   body,
		   creator_id,
		   created_at,
		   updated_at,
		   deleted_at
    from main.reports`

func (c *Client) GetReports(ctx context.Context, msg *GetReports) ([]ReportItem, *int, error) {
	row, err := c.driver.Query(ctx, sqlGetReports)
	if err != nil {
		return nil, nil, NewInternalError(err)
	}

	reports := make([]ReportItem, 0, 0)

	for row.Next() {
		report := ReportItem{}
		err := row.Scan(
			&report.Id,
			&report.Title,
			&report.Date,
			&report.StartTime,
			&report.EndTime,
			&report.BreakTime,
			&report.WorkTime,
			&report.Body,
			&report.CreatorId,
			&report.CreatedAt,
			&report.UpdatedAt,
			&report.DeletedAt,
		)
		if err != nil {
			return nil, nil, NewInternalError(err)
		}
		reports = append(reports, report)
	}

	return reports, ptr.Int(len(reports)), nil
}

type GetReports struct {
	InvokerId string `json:"invoker_id,omitempty"`
}

type ReportItem struct {
	Id        *string    `json:"id,omitempty"`
	Title     *string    `json:"title,omitempty"`
	Date      *time.Time `json:"date,omitempty"`
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty"`
	BreakTime *time.Time `json:"break_time,omitempty"`
	WorkTime  *time.Time `json:"work_time,omitempty"`
	Body      *string    `json:"body,omitempty"`
	CreatorId *string    `json:"creator_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
