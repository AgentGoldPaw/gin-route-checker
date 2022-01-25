package gin_route_checker

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

func SetRoute(methods []string, route string, statusCode int, headers map[string]string) {
	routes[route] = &RouteOptions{
		Methods:    methods,
		StatusCode: statusCode,
		Headers:    headers,
	}
}

func SetRoutes(r map[string]*RouteOptions) error {
	for _, opts := range r {
		for _, method := range opts.Methods {
			switch method {
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

	}
	routes = r
	return nil
}

func CheckRoutes(server *gin.Engine, test *testing.T) error {
	unitTest.SetRouter(server)
	for route, opts := range routes {
		for _, method := range opts.Methods {
			_, resp, err := unitTest.TestOrdinaryHandler(
				method, route, "json", nil, opts.Headers)
			if err != nil {
				return err
			}
			assert.Equal(test, opts.StatusCode, resp.StatusCode)

		}

	}
	return nil
}
