package login

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// correct loginid and password
func TestLogin(test *testing.T) {
	e := echo.New()

	userJSON := `{"loginId":"lahu.r.gunjal1996@gmail.com","password":"lahugunjal"}`
	req, _ := http.NewRequest(echo.POST, "o/loginUser", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, loginRoute(c)) {
		assert.Equal(test, http.StatusOK, rec.Code)
	}
}

//Blank Password and LoginId
func Test2Login(test *testing.T) {
	e := echo.New()

	userJSON := `{"loginId":"","passWord":""}`
	req, _ := http.NewRequest(echo.POST, "o/loginUser", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, loginRoute(c)) {
		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
	}
}
