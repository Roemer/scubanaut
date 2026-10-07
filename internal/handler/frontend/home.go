package frontend

import (
	"net/http"

	"github.com/roemer/scubanaut/internal/component"
)

func (h *FrontendHandler) Home(w http.ResponseWriter, r *http.Request) {
	type pageData struct {
		Breadcrumbs component.Breadcrumb
	}
	data := pageData{
		Breadcrumbs: component.NewBreadcrumb(),
	}
	h.renderPage(w, r, "page/home.html", data)
}
