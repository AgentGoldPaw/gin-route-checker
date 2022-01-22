package gin_route_checker

import (
	"errors"
	"github.com/gin-gonic/gin"
	unitTest "github.com/golden-protocol/gin_unit_test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	routes = make(Routes)
)

func SetRoute(method, route string, statusCode int) {
	routes[route] = &RouteOptions{
		Method: method,
		StatusCode: statusCode,
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
			opts.Method, route, "json", nil, nil)
		if err != nil {
			return err
		}
		assert.Equal(test, opts.StatusCode, resp.StatusCode)
	}
	return nil
}