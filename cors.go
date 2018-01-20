package cors

import (
	"time"
)

// Options ...
type Options struct {
	AllowOrigins  []string
	AllowMethods  []string
	AllowHeaders  []string
	ExposeHeaders []string
	MaxAge        time.Duration
}
