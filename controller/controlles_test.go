package controller

import (

	"echoServise/store/filestore"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouteGetUser(t *testing.T) {
	e := echo.New()
	s := filestore.NewStore()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec :=  httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := GetUser(s)(c)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rec.Body.String())

}
func TestAddUser(t *testing.T) {
	e := echo.New()
	s := filestore.NewStore()
	req := httptest.NewRequest(http.MethodPost, "/",strings.NewReader(`{"name":"Test Controller"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := AddUser(s)(c)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rec.Body.String())
}
func TestGetUsersList(t *testing.T) {
	e := echo.New()
	s := filestore.NewStore()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec :=  httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := GetUsersList(s)(c)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rec.Body.String())
}
func TestRemoveUser(t *testing.T) {
	e := echo.New()
	s := filestore.NewStore()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec :=  httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := RemoveUser(s)(c)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rec.Body.String())
}
func TestUpdateUser(t *testing.T) {
	e := echo.New()
	s := filestore.NewStore()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{"name":"Update Name"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec :=  httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := UpdateUser(s)(c)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rec.Body.String())
}