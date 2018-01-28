package cors

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type requestOptions struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    io.Reader
}

func init() {
	gin.SetMode(gin.TestMode)
}

func newTestRouter(opts Options) *gin.Engine {
	router := gin.New()
	router.Use(Middleware(opts))
	router.GET("/", func(c *gin.Context) {
		c.String(200, "get")
	})
	router.POST("/", func(c *gin.Context) {
		c.String(200, "post")
	})
	router.PUT("/", func(c *gin.Context) {
		c.String(200, "put")
	})
	router.PATCH("/", func(c *gin.Context) {
		c.String(200, "patch")
	})
	return router
}

func request(server *gin.Engine, opts requestOptions) *httptest.ResponseRecorder {

	w := httptest.NewRecorder()
	req, err := http.NewRequest(opts.Method, opts.URL, opts.Body)
	if err != nil {
		panic(err)
	}

	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	server.ServeHTTP(w, req)
	return w
}
