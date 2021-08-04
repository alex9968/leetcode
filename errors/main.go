package main

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/gin-gonic/gin"
)

var ErrNotFound = errors.New("record not found")

func main() {
	err := f1()
	if err !=  nil && errors.Is(ErrNotFound, errors.Cause(err)) {
		fmt.Println("没找到")
		fmt.Printf("trace:%+v", err)
	}


}

func f1() error {

	err := errors.Wrap(ErrNotFound,"f1 failed")
	return errors.WithMessage(err,"with message")
}


type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func AccessLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		fmt.Sprintf("url=%s, status=%d, resp=%s", c.Request.URL, c.Writer.Status(), blw.body.String())
	}
}
