package http

import (
	"fmt"
	"github.com/Ovsienko023/reporter/app/core"
	_ "github.com/Ovsienko023/reporter/docs"
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func RegisterHTTPEndpoints(router chi.Router, c core.Core, apiConfig *configuration.Api) http.Handler {
	swaggerUrl := fmt.Sprintf("http://%s:%s/docs/doc.json", apiConfig.Doc.Host, apiConfig.Doc.Port)
	h := NewTransport(c)

	router.Get("/echo", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("{'status': 'ok', 'version: '1.1.22'}"))
	})
	// DOCS
	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(swaggerUrl)))

	// AUTH
	router.Post("/api/v1/sign_in", h.SignIn)
	router.Post("/api/v1/sign_up", h.SignUp)

	router.Get("/api/v1/auth", h.Auth)
	router.Post("/api/v1/auth/provider", h.GetProviderUri)

	// PROFILE
	router.Get("/api/v1/profile", h.GetProfile)
	router.Put("/api/v1/profile", h.UpdateProfile)

	// USERS
	router.Get("/api/v1/users", h.GetUsers)

	// CALENDAR EVENTS
	router.Get("/api/v1/calendar", h.GetCalendarEvents)

	//REPORTS
	router.Get("/api/v1/reports", h.GetReports)
	router.Get("/api/v1/reports/{report_id}", h.GetReport)
	router.Post("/api/v1/reports", h.CreateReport)
	router.Put("/api/v1/reports/{report_id}", h.UpdateReport)
	router.Delete("/api/v1/reports/{report_id}", h.DeleteReport)
	router.Get("/api/v1/export/reports", h.ExportReportsToCsv)

	// DayOffs
	router.Get("/api/v1/day_offs/{day_off_id}", h.GetDayOff)
	router.Post("/api/v1/day_offs", h.CreateDayOff)
	router.Delete("/api/v1/day_offs/{day_off_id}", h.DeleteDayOff)

	// SICK LEAVES
	router.Get("/api/v1/sick_leaves/{sick_leave_id}", h.GetSickLeave)
	router.Post("/api/v1/sick_leaves", h.CreateSickLeave)
	router.Delete("/api/v1/sick_leaves/{sick_leave_id}", h.DeleteSickLeave)

	// VACATION PAID
	router.Get("/api/v1/vacations_paid/{vacation_paid_id}", h.GetVacationPaid)
	router.Post("/api/v1/vacations_paid", h.CreateVacationPaid)
	router.Delete("/api/v1/vacations_paid/{vacation_paid_id}", h.DeleteVacationPaid)

	// VACATION UNPAID
	router.Get("/api/v1/vacations_unpaid/{vacations_unpaid_id}", h.GetVacationUnpaid)
	router.Post("/api/v1/vacations_unpaid", h.CreateVacationUnpaid)
	router.Delete("/api/v1/vacations_unpaid/{vacation_unpaid_id}", h.DeleteVacationUnpaid)

	// STATS
	router.Get("/api/v1/stats", h.GetStatistics)

	// OTHERS
	router.Post("/api/v1/send_email", h.SendEmail)

	// FOR ADMINISTRATORS
	// PERMISSIONS
	router.Post("/api/v1/users/{user_id}/permissions", h.AddObjectToUserPermission)
	router.Delete("/api/v1/users/{user_id}/permissions", h.RemoveObjectFromUserPermission)

	return router
}
