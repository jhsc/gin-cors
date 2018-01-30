package cors

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

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

func TestDefault(t *testing.T) {
	r := newTestRouter(Options{})
	assert := assert.New(t)

	req := request(r, requestOptions{
		URL: "/",
		Headers: map[string]string{
			"Origin": "http://test.com",
		},
	})

	assert.Equal("http://test.com", req.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal("get", req.Body.String())
}

func TestAllowOrigins(t *testing.T) {
	r := newTestRouter(Options{
		AllowOrigins: []string{"http://test.com"},
	})
	assert := assert.New(t)

	req := request(r, requestOptions{
		URL: "/",
		Headers: map[string]string{
			"Origin": "http://test.com",
		},
	})
	assert.Equal("http://test.com", req.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal("get", req.Body.String())
}

func TestAllowHeaders(t *testing.T) {
	r := newTestRouter(Options{
		ExposeHeaders: []string{"Foo", "Bar"},
	})
	assert := assert.New(t)

	req := request(r, requestOptions{
		URL: "/",
		Headers: map[string]string{
			"Origin": "http://test.com",
		},
	})
	assert.Equal("Foo,Bar", req.Header().Get("Access-Control-Expose-Headers"))
	assert.Equal("get", req.Body.String())
}
