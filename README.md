# gin-cors

CORS middleware for Gin.

## Installation
Download and install:
```sh
$ go get github.com/jhsc/gin-cors
```

Import it in your code:

```go
import "github.com/jhsc/gin-cors"
```

### Example:
```go
import (
  "github.com/jhsc/gin-cors"
  "github.com/gin-gonic/gin"
)

func main() {
  g := gin.New()
  g.Use(cors.Options{
		AllowOrigins: []string{"http://test.com"},
	})
}
```