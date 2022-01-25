package ginroutechecker

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

func SetRoutes(r map[string][]*RouteOptions) error {
	for _, routes := range r {
		for _, opts := range routes {
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
	}
	routes = r
	return nil
}

func CheckRoutes(server *gin.Engine, test *testing.T) error {
	unitTest.SetRouter(server)
	for route, routeOptions := range routes {
		for _, opts := range routeOptions {
			_, resp, err := unitTest.TestOrdinaryHandler(
				opts.Method, route, "json", nil, opts.Headers)
			if err != nil {
				return err
			}
			assert.Equal(test, opts.StatusCode, resp.StatusCode)
		}
	}
	return nil
}
