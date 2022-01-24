package github.com/golden-protocol/gin-route-checker

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	unitTest "github.com/golden-protocol/gin_unit_test"
	"github.com/stretchr/testify/assert"
)

var (
	routes = make(Routes)
)

func SetRoute(method, route string, statusCode int, headers map[string]string) {
	routes[route] = &RouteOptions{
		Method:     method,
		StatusCode: statusCode,
		Headers:    headers,
	}
}

func SetRoutes(r map[string]*RouteOptions) error {
	for _, opts := range r {
		switch opts.Method {
		case http.MethodGet:
		case http.MethodPost:
		case http.MethodDelete:
		case http.MethodPatch:
		case http.MethodPut:
			continue
		default:
			return errors.New("invalid method")
		}
	}
	routes = r
	return nil
}

func CheckRoutes(server *gin.Engine, test *testing.T) error {
	unitTest.SetRouter(server)
	for route, opts := range routes {
		_, resp, err := unitTest.TestOrdinaryHandler(
			opts.Method, route, "json", nil, opts.Headers)
		if err != nil {
			return err
		}
		assert.Equal(test, opts.StatusCode, resp.StatusCode)
	}
	return nil
}
