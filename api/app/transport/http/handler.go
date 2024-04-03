package http

import (
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/transport/http/handlers"
	"net/http"
)

type Transport struct {
	core core.Core
}

func NewTransport(c core.Core) *Transport {
	return &Transport{
		core: c,
	}
}

// AUTH

// SignIn ...  Sign In
// @Summary Sign In
// @Description Getting an authorization token
// @Tags Auth
// @Param request body domain.SignInRequest true "body params"
// @Success 200 {object} domain.SignInResponse
// @Success 401
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/sign_in [post]
func (t *Transport) SignIn(w http.ResponseWriter, r *http.Request) {
	handlers.SignIn(&t.core, w, r)
}

// SignUp ...  SignUp
// @Summary Sign Up
// @Description User registration
// @Tags Auth
// @Param request body domain.SignUpRequest true "body params"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/sign_up [post]
func (t *Transport) SignUp(w http.ResponseWriter, r *http.Request) {
	handlers.SignUp(&t.core, w, r)
}

// Auth ...  Auth
// @Summary Auth
// @Description Auth
// @Tags Auth
// @Param request body domain.AuthRequest true "query params"
// @Success 200 {object} domain.AuthResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/auth [get]
func (t *Transport) Auth(w http.ResponseWriter, r *http.Request) {
	handlers.Auth(&t.core, w, r)
}

// GetProviderUri ...  GetProviderUri
// @Summary Get Provider Uri
// @Description Get Provider Uri
// @Tags Auth
// @Param request body domain.GetProviderUriRequest true "body params"
// @Success 200 {object} domain.GetProviderUriResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/auth/provider [post]
func (t *Transport) GetProviderUri(w http.ResponseWriter, r *http.Request) {
	handlers.GetProviderUri(&t.core, w, r)
}

// PROFILE

// GetProfile ...  Get Profile
// @Summary Get Profile
// @Description Getting user data
// @Tags Profile
// @Param request body domain.GetProfileRequest true "query params"
// @Success 200 {object} domain.GetProfileResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/profile [get]
func (t *Transport) GetProfile(w http.ResponseWriter, r *http.Request) {
	handlers.GetProfile(&t.core, w, r)
}

// UpdateProfile ...  Update Profile
// @Summary Update Profile
// @Description Updating user data
// @Tags Profile
// @Param request body domain.UpdateProfileRequest true "body params"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/profile [put]
func (t *Transport) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateProfile(&t.core, w, r)
}

// USERS

// GetUsers ...  Get Users
// @Summary Get Users
// @Description Get Users
// @Tags Users
// @Param request body domain.GetUsersRequest true "query params"
// @Success 200 {object} domain.GetUsersResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/users [get]
func (t *Transport) GetUsers(w http.ResponseWriter, r *http.Request) {
	handlers.GetUsers(&t.core, w, r)
}

// CALENDAR EVENTS

// GetCalendarEvents ...  Get Calendar Events
// @Summary Get Calendar Events
// @Description Get Calendar Events
// @Tags Calendar
// @Param request body domain.GetCalendarEventsRequest true "query params"
// @Success 200 {object} domain.GetCalendarEventsResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/calendar [get]
func (t *Transport) GetCalendarEvents(w http.ResponseWriter, r *http.Request) {
	handlers.GetCalendarEvents(&t.core, w, r)
}

// REPORTS

// GetReport ... Get report
// @Summary Get report
// @Description get report
// @Tags Reports
// @Param   report_id   path      string  true  "report_id"
// @Success 200 {object} domain.GetReportResponse
// @Failure 404 {object} httperror.ErrorResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/reports/{report_id} [get]
func (t *Transport) GetReport(w http.ResponseWriter, r *http.Request) {
	handlers.GetReport(&t.core, w, r)
}

// GetReports ... Get all reports
// @Summary Get all reports
// @Description get all reports
// @Tags Reports
// @Param request body domain.GetReportsRequest true "query params"
// @Success 200 {object} domain.GetReportsResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/reports [get]
func (t *Transport) GetReports(w http.ResponseWriter, r *http.Request) {
	handlers.GetReports(&t.core, w, r)
}

// CreateReport ...  Create report
// @Summary Create report
// @Description Create report
// @Tags Reports
// @Param request body domain.CreateReportRequest true "body params"
// @Success 201 {object} domain.CreateReportResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/reports [post]
func (t *Transport) CreateReport(w http.ResponseWriter, r *http.Request) {
	handlers.CreateReport(&t.core, w, r)
}

// UpdateReport ...  Update report
// @Summary Update report
// @Description Update report
// @Tags Reports
// @Param   id   path      string  true  "report_id"
// @Param request body domain.UpdateReportRequest true "body params"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/reports/{report_id} [put]
func (t *Transport) UpdateReport(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateReport(&t.core, w, r)
}

// DeleteReport ...  Delete report
// @Summary Delete report
// @Description Delete report
// @Tags Reports
// @Param   id   path      string  true  "report_id"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/reports/{report_id} [delete]
func (t *Transport) DeleteReport(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteReport(&t.core, w, r)
}

// GetSickLeave ... Get sick leave
// @Summary Get sick leave
// @Description get sick leave
// @Tags SickLeaves
// @Param   sick_leave_id   path      string  true  "sick_leave_id"
// @Success 200 {object} domain.GetSickLeaveResponse
// @Failure 403 {object} httperror.ErrorResponse
// @Failure 404 {object} httperror.ErrorResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/sick_leaves/{sick_leave_id} [get]
func (t *Transport) GetSickLeave(w http.ResponseWriter, r *http.Request) {
	handlers.GetSickLeave(&t.core, w, r)
}

// CreateSickLeave ...  Create sick leave
// @Summary Create sick leave
// @Description Create sick leave
// @Tags SickLeaves
// @Param request body domain.CreateSickLeaveRequest true "body params"
// @Success 201 {object} domain.CreateSickLeaveResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/sick_leaves [post]
func (t *Transport) CreateSickLeave(w http.ResponseWriter, r *http.Request) {
	handlers.CreateSickLeave(&t.core, w, r)
}

// DeleteSickLeave ...  Delete sick leave
// @Summary Delete sick leave
// @Description Delete sick leave
// @Tags SickLeaves
// @Param   id   path      string  true  "sick_leave_id"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/sick_leaves/{sick_leave_id} [delete]
func (t *Transport) DeleteSickLeave(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteSickLeave(&t.core, w, r)
}

//-------------

// GetDayOff ... Get day off
// @Summary Get day off
// @Description get day off
// @Tags DayOffs
// @Param   day_off_id   path      string  true  "day_off_id"
// @Success 200 {object} domain.GetDayOffResponse
// @Failure 403 {object} httperror.ErrorResponse
// @Failure 404 {object} httperror.ErrorResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/day_offs/{day_off_id} [get]
func (t *Transport) GetDayOff(w http.ResponseWriter, r *http.Request) {
	handlers.GetDayOff(&t.core, w, r)
}

// CreateDayOff ...  Create day off
// @Summary Create day off
// @Description Create day off
// @Tags DayOffs
// @Param request body domain.CreateDayOffRequest true "body params"
// @Success 201 {object} domain.CreateDayOffResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/day_offs [post]
func (t *Transport) CreateDayOff(w http.ResponseWriter, r *http.Request) {
	handlers.CreateDayOff(&t.core, w, r)
}

// DeleteDayOff ...  Delete day off
// @Summary Delete day off
// @Description Delete day off
// @Tags DayOffs
// @Param   id   path      string  true  "day_off_id"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/day_offs/{day_off_id} [delete]
func (t *Transport) DeleteDayOff(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteDayOff(&t.core, w, r)
}

//-------------

// GetVacationPaid ... Get vacation paid
// @Summary Get vacation paid
// @Description get vacation paid
// @Tags VacationPaid
// @Param   vacation_paid_id   path      string  true  "vacation_paid_id"
// @Success 200 {object} domain.GetVacationPaidResponse
// @Failure 403 {object} httperror.ErrorResponse
// @Failure 404 {object} httperror.ErrorResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/vacations_paid/{vacations_paid_id} [get]
func (t *Transport) GetVacationPaid(w http.ResponseWriter, r *http.Request) {
	handlers.GetVacationPaid(&t.core, w, r)
}

// CreateVacationPaid ...  Create vacation paid
// @Summary Create vacation paid
// @Description Create vacation paid
// @Tags VacationPaid
// @Param request body domain.CreateVacationPaidRequest true "body params"
// @Success 201 {object} domain.CreateVacationPaidResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/vacations_paid [post]
func (t *Transport) CreateVacationPaid(w http.ResponseWriter, r *http.Request) {
	handlers.CreateVacationPaid(&t.core, w, r)
}

// DeleteVacationPaid ...  Delete vacation paid
// @Summary Delete vacation paid
// @Description Delete vacation paid
// @Tags VacationPaid
// @Param   id   path      string  true  "vacation_paid_id"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/vacations_paid/{vacation_paid_id} [delete]
func (t *Transport) DeleteVacationPaid(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteVacationPaid(&t.core, w, r)
}

// GetVacationUnpaid ... Get vacation unpaid
// @Summary Get vacation unpaid
// @Description get vacation unpaid
// @Tags VacationUnpaid
// @Param   vacation_id   path      string  true  "vacations_unpaid_id"
// @Success 200 {object} domain.GetVacationUnpaidResponse
// @Failure 403 {object} httperror.ErrorResponse
// @Failure 404 {object} httperror.ErrorResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/vacations_unpaid/{vacations_unpaid_id} [get]
func (t *Transport) GetVacationUnpaid(w http.ResponseWriter, r *http.Request) {
	handlers.GetVacationUnpaid(&t.core, w, r)
}

// CreateVacationUnpaid ...  Create vacation unpaid
// @Summary Create vacation unpaid
// @Description Create vacation unpaid
// @Tags VacationUnpaid
// @Param request body domain.CreateVacationUnpaidRequest true "body params"
// @Success 201 {object} domain.CreateVacationUnpaidResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/vacations_unpaid [post]
func (t *Transport) CreateVacationUnpaid(w http.ResponseWriter, r *http.Request) {
	handlers.CreateVacationUnpaid(&t.core, w, r)
}

// DeleteVacationUnpaid ...  Delete vacation unpaid
// @Summary Delete vacation unpaid
// @Description Delete vacation unpaid
// @Tags VacationUnpaid
// @Param   id   path      string  true  "vacation_unpaid_id"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/vacations_unpaid/{vacation_unpaid_id} [delete]
func (t *Transport) DeleteVacationUnpaid(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteVacationUnpaid(&t.core, w, r)
}

// ExportReportsToCsv ... Get all reports
// @Summary Export reports to csv
// @Description Export reports to csv
// @Tags Reports
// @Param request body domain.GetReportsRequest true "query params"
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/export/reports [get]
func (t *Transport) ExportReportsToCsv(w http.ResponseWriter, r *http.Request) {
	handlers.ExportReportsToCsv(&t.core, w, r)
}

// SendEmail ...  Sending a report by mail
// @Summary Send email
// @Description Send email
// @Tags Email
// @Param request body domain.SendEmailRequest true "body params"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/send_mail [post]
func (t *Transport) SendEmail(w http.ResponseWriter, r *http.Request) {
	handlers.SendEmail(&t.core, w, r)
}

// GetStatistics ...  Get Statistics
// @Summary Get Statistics
// @Description Get Statistics
// @Tags Statistics
// @Param request body domain.GetStatisticsRequest true "query params"
// @Success 200 {object} domain.GetStatisticsResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/stats [get]
func (t *Transport) GetStatistics(w http.ResponseWriter, r *http.Request) {
	handlers.GetStatistics(&t.core, w, r)
}

// AddObjectToUserPermission ...  Give the user permission to read an object
// @Summary Add object to user permission
// @Description Give the user permission to read an object
// @Tags Permission
// @Param request body domain.AddObjectToUserPermissionRequest true "body params"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/users/{user_id}/permissions [post]
func (t *Transport) AddObjectToUserPermission(w http.ResponseWriter, r *http.Request) {
	handlers.AddObjectToUserPermission(&t.core, w, r)
}

// RemoveObjectFromUserPermission ...  Remove the user's permission to read an object
// @Summary Remove object from user permission
// @Description Remove the user's permission to read an object
// @Tags Permission
// @Param request body domain.AddObjectToUserPermissionRequest true "body params"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/users/{user_id}/permissions [delete]
func (t *Transport) RemoveObjectFromUserPermission(w http.ResponseWriter, r *http.Request) {
	handlers.RemoveObjectFromUserPermission(&t.core, w, r)
}
