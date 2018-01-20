package cors

import (
	"time"

	"github.com/gin-gonic/gin"
)

var (
	defaultAllowHeaders = []string{"Origin", "Accept", "Content-Type", "Authorization"}
	defaultAllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"}
)

// Options ...
type Options struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           time.Duration
}

// Middleware sets CORS headers for every request
func Middleware(opts Options) gin.HandlerFunc {
	if opts.AllowHeaders == nil {
		opts.AllowHeaders = defaultAllowHeaders
	}

	if opts.AllowMethods == nil {
		opts.AllowMethods = defaultAllowMethods
	}
}
