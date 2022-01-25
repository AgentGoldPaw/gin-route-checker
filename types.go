package gin_route_checker

type RouteOptions struct {
	Method     string
	StatusCode int
	Headers    map[string]string
}

type Routes map[string][]*RouteOptions
