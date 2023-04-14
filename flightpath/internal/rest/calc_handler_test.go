package rest_test

import (
	"flightpath/internal/rest"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFindRoute(t *testing.T) {

	var (
		requests = []string{
			``,
			`[["SAN", "LAS"],["LAS", "SFO"],["sfo", "san"]]`, // cycle
			`[["sfo", "EWR"]`,                  // bad json
			`[["sfo", "EWR"]]`,                 // trivial
			`[["sfo", "ewr"], ["ewr", "sfo"]]`, // trivial roundtrip
			`[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`,
		}
		expected  = `["SFO","EWR"]`
		wantError = map[int]bool{0: true, 1: true, 2: true, 3: false, 4: false, 5: false, 6: false}
	)
	for idx, request := range requests {
		e := rest.Setup()
		req := httptest.NewRequest(http.MethodPost, rest.RouteEndpoint, strings.NewReader(request))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		// SUT - our handle
		err := rest.FindRoute(c)

		v, ok := wantError[idx]
		if ok && v {
			require.Error(t, err)
			//require.Equal(t, http.StatusBadRequest, rec.Code)
		} else {
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, rec.Code)
			actual := strings.TrimSpace(rec.Body.String())
			assert.Equal(t, expected, actual)
		}
	}
}
