package response

import (
	"github.com/gin-gonic/gin"
	"sync"
)

const (
	CodeSuccess = 0
	CodeFail    = -1
)

type Gin struct {
	c      *gin.Context
	resp   *Response
	status int
}

type Response struct {
	Meta    Meta        `json:"meta"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Meta map[string]interface{}

var (
	l = &sync.RWMutex{}
)

func (m Meta) Set(key string, val interface{}) {
	l.Lock()
	m[key] = val
	l.Unlock()
}

// NewGinResponse
func NewGinResponse(c *gin.Context) *Gin {
	return &Gin{
		c,
		&Response{
			Meta: Meta{},
		},
		200,
	}
}

func (g *Gin) Fail(code int, message string) *Gin {
	g.resp.Code = code
	g.resp.Message = message
	return g
}

func (g *Gin) SetStatus(status int) *Gin {
	g.status = status
	return g
}

func (g *Gin) SetMeta(key string, val interface{}) *Gin {
	g.resp.Meta.Set(key, val)
	return g
}

func (g *Gin) Success(data interface{}) *Gin {
	g.resp.Code = CodeSuccess
	g.resp.Message = "success"
	g.resp.Data = data
	return g
}

// Response setting gin.JSON
func (g *Gin) Response() {
	g.c.JSON(g.status, g.resp)
	g.c.Abort()
	return
}
