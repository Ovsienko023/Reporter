package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"time"
)

type CreateReportRequest struct {
	Token       string `json:"token,omitempty" swaggerignore:"true"`
	DisplayName string `json:"display_name,omitempty"`
	Date        int64  `json:"date,omitempty"`
	StartTime   int64  `json:"start_time,omitempty"`
	EndTime     int64  `json:"end_time,omitempty"`
	BreakTime   int64  `json:"break_time,omitempty"`
	WorkTime    int64  `json:"work_time,omitempty"`
	Body        string `json:"body,omitempty"`
}

func (r *CreateReportRequest) ToDbCreateReport(invokerId string) *repository.CreateReport {
	return &repository.CreateReport{
		InvokerId:   invokerId,
		DisplayName: r.DisplayName,
		Date:        time.Unix(r.Date, 0).UTC(),
		StartTime:   time.Unix(r.StartTime, 0).UTC(),
		EndTime:     time.Unix(r.EndTime, 0).UTC(),
		BreakTime:   time.Unix(r.BreakTime, 0).UTC(),
		WorkTime:    time.Unix(r.WorkTime, 0).UTC(),
		Body:        r.Body,
	}
}

type CreateReportResponse struct {
	Id string `json:"id,omitempty"`
}

func FromCreateReportResponse(resp *repository.CreatedReport) *CreateReportResponse {
	if resp == nil {
		return nil
	}

	return &CreateReportResponse{
		Id: resp.Id,
	}
}
