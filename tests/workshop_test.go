package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func TestPackageValidityChecker(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/workshops/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
}

func TestGetWorkshopDetails(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/workshops/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
}
