package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RequestInfos() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		method := c.Request.Method
		fmt.Printf("请求路径为: %s,请求协议为: %s\n", path, method)
	}
}
