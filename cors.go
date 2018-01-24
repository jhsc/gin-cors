package cors

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

	return func(c *gin.Context) {
		req := c.Request
		resp := c.Writer
		origin := req.Header.Get("Origin")
		reqMethod := req.Header.Get("Access-Control-Request-Method")
		reqHeaders := req.Header.Get("Access-Control-Request-Headers")

		fmt.Printf("request :%v", req)
		if len(opts.AllowOrigins) > 0 {
			for _, ao := range opts.AllowOrigins {
				if ao == origin {
					//req origin match our trusted origin list
					resp.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
		} else {
			resp.Header().Set("Access-Control-Allow-Origin", origin)
		}

		if opts.AllowCredentials {
			resp.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if len(opts.ExposeHeaders) > 0 {
			resp.Header().Set("Access-Control-Expose-Headers", strings.Join(opts.ExposeHeaders, ","))
		}

		if req.Method == "OPTIONS" {
			if len(opts.AllowMethods) > 0 {
				resp.Header().Set("Access-Control-Allow-Methods", strings.Join(opts.AllowMethods, ","))
			} else if reqMethod != "" {
				resp.Header().Set("Access-Control-Allow-Methods", reqMethod)
			}

			if len(opts.AllowHeaders) > 0 {
				resp.Header().Set("Access-Control-Allow-Headers", strings.Join(opts.AllowHeaders, ","))
			} else if reqHeaders != "" {
				resp.Header().Set("Access-Control-Allow-Headers", reqHeaders)
			}

			if opts.MaxAge > time.Duration(0) {
				resp.Header().Set("Access-Control-Max-Age", strconv.FormatInt(int64(opts.MaxAge/time.Second), 10))
			}

			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
}
