package main

import (
	"fmt"
	"gindemo/middleware"
	_ "gindemo/middleware"
	"gindemo/models"
	_ "gindemo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	//engine.Use(middleware.RequestInfos())
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

	/*engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		username := context.PostForm("username")
		password := context.PostForm("password")
		fmt.Printf("%s -> %s\n", username, password)
		context.Writer.WriteString("login success,Username: " + username + ",password: " + password)
	})*/

	/*engine.DELETE("/user/:id", func(context *gin.Context) {
		id := context.Param("id")
		fmt.Println("delete id: ", id)
		context.Writer.WriteString("delete id " + id + "success")
	})*/

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
		fmt.Printf("%#v\n", user)
		context.Writer.WriteString("bind user success! info: " + fmt.Sprint(user))
	})

	// 返回json数据
	engine.GET("/json", func(context *gin.Context) {
		mmap := map[string]interface{}{
			"code":    1,
			"message": "ok",
			"data":    context.FullPath(),
		}

		context.JSON(200, mmap)
	})

	// 返回结构体数据
	engine.GET("/struct", func(context *gin.Context) {
		resp := models.Response{
			1, "ok", "response data",
		}

		context.JSON(200, &resp)
	})

	// 需要设置全局的静态文件目录 否则无法正常执行
	engine.LoadHTMLGlob("./view/*")
	// 返回html文件
	engine.GET("/html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"fullPath": context.FullPath(),
		})
	})

	//加载静态资源
	engine.Static("/img", "./img")

	// 使用group来管理请求
	group := engine.Group("/user")
	group.POST("/register", registerHandler)

	group.POST("/login", loginHandler)

	group.DELETE("/:id", deleteHandler)

	// 为单独的某个请求设置中间件.
	engine.GET("/query", middleware.RequestInfos(), func(context *gin.Context) {
		fmt.Println("hello world")
		context.JSON(200, map[string]interface{}{
			"code": 200,
			"data": "ok",
		})
	})
	// 设置使用的端口号为8090
	// engine.Run(":8090")
	if _, err := models.Db.Exec("CREATE TABLE person(" +
		"id int AUTO_INCREMENT PRIMARY KEY," +
		"name varchar(12) NOT NULL," +
		"age int DEFAULT 1" +
		");"); err != nil {
		log.Fatal(err.Error())
		return
	}
	if err := engine.Run(); err != nil {
		//log.Fatal(err.Error())
		fmt.Println(err.Error())
	}
}

func deleteHandler(context *gin.Context) {
	id := context.Param("id")
	context.Writer.WriteString("删除ID为: " + id + "的用户.")
}

func loginHandler(context *gin.Context) {
	fullPath := "用户登录: " + context.FullPath()
	context.Writer.WriteString(fullPath)
}

func registerHandler(context *gin.Context) {
	fullPath := "用户注册: " + context.FullPath()
	context.Writer.WriteString(fullPath)
}
