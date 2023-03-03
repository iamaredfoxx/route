package arah

import "github.com/labstack/echo/v4"

// Route Path Implementation
type RoutePathInterface interface {
	Group(group string, f func(rg RoutePathInterface), m ...echo.MiddlewareFunc)
	Get(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath
	Post(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath
	Put(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath
	Patch(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath
	Delete(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *routePath
	Any(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) []*routePath
}

// Route Implementation
type RouteInterface interface {
	Create(RoutePathInterface)
}

// Start Implementation
type StartInterface interface {
	Start(e *echo.Echo) error
}
