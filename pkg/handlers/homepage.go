package handlers

import (
	"ghost/internal/common"
	"ghost/internal/pages"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	Settings common.AppSettings
}

func (h HomeHandler) HandleIndex(c echo.Context) error {
	return render(c, pages.Base(h.Settings.AppNavbar.Links))
}
