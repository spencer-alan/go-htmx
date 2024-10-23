package routes

import (
	// "database/sql"
	"ghost/internal/common"
	"ghost/pkg/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, settings common.AppSettings) {
	// Routes for Layout
	indexHandler := handlers.HomeHandler{Settings: settings}
	issueHandler := handlers.IssueHandler{}
	e.GET("/", indexHandler.HandleIndex)
	e.GET("/issues/all", issueHandler.HandleAll)

	// Test for the card to see how HTMX is working etc.
	e.GET("/card-content", func(c echo.Context) error {
		return c.String(http.StatusOK, "More content loaded!")
	})
	e.GET("/card-action", func(c echo.Context) error {
		return c.String(http.StatusOK, "<button>Action Completed!</button>")
	})
}
