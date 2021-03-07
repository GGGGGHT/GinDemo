package main

import (
	"fmt"
	"gindemo/models"
	_ "gindemo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径: ", context.FullPath())
		context.Writer.WriteString("world")
	})

	engine.Handle("GET", "/test", func(c *gin.Context) {
		fullPath := c.FullPath()
		fmt.Println("fullpath: ", fullPath)
		name := c.DefaultQuery("name", "hello")
		fmt.Println("name: ", name)
		c.Writer.WriteString("hello" + name)
	})

	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		username := context.PostForm("username")
		password := context.PostForm("password")
		fmt.Printf("%s -> %s\n", username, password)
		context.Writer.WriteString("login success,Username: " + username + ",password: " + password)
	})

	engine.DELETE("/user/:id", func(context *gin.Context) {
		id := context.Param("id")
		fmt.Println("delete id: ", id)
		context.Writer.WriteString("delete id " + id + "success")
	})

	// 绑定参数
	engine.GET("/stu", func(context *gin.Context) {
		stu := models.Student{}
		if err := context.ShouldBindQuery(&stu); err == nil {
			fmt.Printf("%#v\n", stu)
		}
	})

	engine.POST("/reg", func(context *gin.Context) {
		var user models.Register
		// 使用json绑定
		//if err := context.BindJSON(&user); err != nil {
		//	fmt.Errorf("%#v\n", err)
		//}
		context.ShouldBind(&user)
		fmt.Printf("%#v\n",user)
		context.Writer.WriteString("bind user success! info: " + fmt.Sprint(user))
	})

	if err := engine.Run(); err != nil {
		//log.Fatal(err.Error())
		fmt.Println(err.Error())
	}
}
