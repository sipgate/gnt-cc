package router_test

import (
	"bytes"
	"encoding/json"

	"gnt-cc/auth"
	"gnt-cc/config"
	"gnt-cc/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "gnt-cc/docs"

	"github.com/gin-gonic/gin"
)

func Login(username string, password string) (int, string) {
	r := gin.New()
	router.Routes(r, true)

	recorder := httptest.NewRecorder()
	login, _ := json.Marshal(auth.Login{
		Username: username,
		Password: password,
	})
	req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(login))
	req.Header.Add("Content-type", "application/json")

	r.ServeHTTP(recorder, req)

	var m map[string]string
	json.Unmarshal(recorder.Body.Bytes(), &m)

	return recorder.Code, m["token"]
}

func TestLogin_can_successfully_login(t *testing.T) {
	config.Parse()

	status, token := Login("admin", "admin")

	assert.NotEqual(t, "", token)
	assert.Equal(t, http.StatusOK, status)
}
func TestLogin_cannot_login_with_bad_credentials(t *testing.T) {
	config.Parse()

	status, token := Login("admin", "not-the-password")

	assert.Equal(t, "", token)
	assert.Equal(t, http.StatusUnauthorized, status)
}
