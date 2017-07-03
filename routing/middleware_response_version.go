package routing

import (
	"github.com/labstack/echo"
)

var (
	_version string = "NO_VERSION"
	_gitSha1 string = "NO_SHA1"
)

//MiddlewareResponseVersion is a middleware to add X-VERSION and X-SHA1 to response headers
func MiddlewareResponseVersion() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := c.Response()
			res.Header().Add("X-VERSION", _version)
			res.Header().Add("X-SHA1", _gitSha1)
			return next(c)
		}
	}
}

//InitResponseVersions sets the local version and gitSha1
func InitResponseVersions(version, gitSha1 string) {
	_version = version
	_gitSha1 = gitSha1
}
