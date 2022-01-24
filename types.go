package github.com/golden-protocol/gin-route-checker

type RouteOptions struct {
	Method     string
	StatusCode int
	Headers    map[string]string
}

type Routes map[string]*RouteOptions
