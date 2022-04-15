package controller

import (
	"belajar-go-echo/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = "{\"data\":[{\"id\":1,\"username\":\"fahmy\"},{\"id\":2,\"username\":\"abida\"}]}\n"
)

func TestGetAllUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImZhaG15Iiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJmYWhteWFiaWRhIiwiZXhwaXJlZF9hdCI6IjMwMzAtMTItMzFUMjM6NTk6NTkrMDc6MDAiLCJjcmVhdGVkX2F0IjoiMjAwMC0xMi0zMVQyMzo1OTo1OSswNzowMCJ9.lx6t9uQhTY6-hg9gf6BW9d3V8l6g0QJG-ZdJmRMf1V0")
	c.SetPath("/users")

	ut := repository.NewTestUnit()
	userController := NewUserController(ut.IFaceUserRepo)

	// // Assertions
	if assert.NoError(t, userController.GetAllData(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetAllUserNoHeaderAuth(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImZhaG15Iiwicm9sZSI6ImFkbWluIiwidXNlcm5hbWUiOiJmYWhteWFiaWRhIiwiZXhwaXJlZF9hdCI6IjMwMzAtMTItMzFUMjM6NTk6NTkrMDc6MDAiLCJjcmVhdGVkX2F0IjoiMjAwMC0xMi0zMVQyMzo1OTo1OSswNzowMCJ9.lx6t9uQhTY6-hg9gf6BW9d3V8l6g0QJG-ZdJmRMf1V0")
	c.SetPath("/users")

	ut := repository.NewTestUnit()
	userController := NewUserController(ut.IFaceUserRepo)

	// // Assertions
	if assert.NoError(t, userController.GetAllData(c)) {
		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, "{\"error\":\"header Authorization invalid value\"}\n", rec.Body.String())
	}
}
