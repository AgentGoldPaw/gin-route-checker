# gin-route-checker
simple route tester for a gin api

## Install 

```bash 
go get github.com/golden-protocol/gin-route-checker
```

## Usage

### Test Multiple Routes

```go
package mypackage_test

import (
	gin_route_checker "gin-route-checker"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
	// setup server 
	server := gin.New()
	server.GET("/test", func(context *gin.Context) {
		context.AbortWithStatus(http.StatusOK)
	})
	// configure the routes to check 
	err := gin_route_checker.SetRoutes(map[string]*gin_route_checker.RouteOptions{
		"/test": &gin_route_checker.RouteOptions{
			Method:     http.MethodGet,
			StatusCode: http.StatusOK,
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	// actually check
	err = gin_route_checker.CheckRoutes(server, t)
	// handle error
	assert.Nil(t, err)
}
```

### Test Single Route

```go
package mypackage_test

import (
	gin_route_checker "gin-route-checker"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
	// setup server 
	server := gin.New()
	server.GET("/test", func(context *gin.Context) {
		context.AbortWithStatus(http.StatusOK)
	})
	// configure the route to check 
	gin_route_checker.SetRoute(http.MethodGet, "/test", http.StatusOK)
	// actually check
	err := gin_route_checker.CheckRoutes(server, t)
	// handle error
	assert.Nil(t, err)
}
```