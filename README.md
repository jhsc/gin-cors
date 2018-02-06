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
    "github.com/gin-gonic/gin"
    "github.com/jhsc/gin-cors"
)

func main(){
    g := gin.New()
    g.Use(cors.Middleware(cors.Options{
      AllowOrigins: []string{"http://test.com"},
    }))
}
```