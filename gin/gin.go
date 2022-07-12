package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	//模式
	gin.SetMode(gin.ReleaseMode)
	//gin.DebugMode
	//gin.ReleaseMode
	//gin.TestMode
	//日志
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//路由引擎
	engine := gin.Default()
	//新路由引擎
	newEngine := gin.New()
	//路由
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	engine.POST("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	engine.PUT("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	engine.DELETE("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	engine.HEAD("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	engine.Any("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	
	//路游参数
	engine.GET("/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.JSON(200, gin.H{
			"msg":  "OK",
			"name": name,
		})
	})
	//路由组
	v1 := engine.Group("v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"msg": "OK",
			})
		})
	}
	//声明HTML
	engine.LoadHTMLGlob("templates/*")
	engine.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//模板
	//使用模板
	//{{template "header" .}}
	//定义模板开始
	//{{define "header"}}
	//定义模板结束
	//{{end}}
	//判断
	//{{if .condition1}}
	//{{else if .condition2}}
	//{{end}}
	//and、or、not
	//{{if and .condition1 .condition2}}
	//{{end}}
	//{{if or .condition1 .condition2}}
	//{{end}}
	//{{if not .condition}}
	//{{end}}
	//等于、不等于、小于、大于
	//{{if eq .var1 .var2}}
	//{{end}}
	//{{if ne .var1 .var2}}
	//{{end}}
	//{{if lt .var1 .var2}}
	//{{end}}
	//{{if gt .var1 .var2}}
	//{{end}}
	//循环
	//{{range $i,$v := .slice}}
	//{{end}}
	//{{range .slice}}
	//  {{$.GlobalVar}}
	//  {{.LocalVar}}
	//  {{.field}}
	//{{end}}
	//响应内容
	//context.JSON(200, gin.H{
	//	"msg": "OK"})
	//context.HTML(200, "index.tmpl", gin.H{
	//	"title": "Main website"})
	//context.Redirect(301, "http://www.google.com/")
	//参数
	engine.GET("/", func(context *gin.Context) {
		name := context.Query("name")
		ago := context.DefaultQuery("ago", "18")
		home, ok := context.GetQuery("home")
		if ok == false {
			context.JSON(404, gin.H{
				"error": "404",
			})
		}
		context.JSON(200, gin.H{
			"msg":  "OK",
			"name": name,
			"ago":  ago,
			"home": home,
		})
	})
	engine.GET("/", func(context *gin.Context) {
		ScoreArr := context.QueryArray("score")
		NameArr, ok := context.GetQueryArray("name")
		if ok == false {
			context.JSON(404, gin.H{
				"error": "404",
			})
		}
		context.JSON(200, gin.H{
			"msg":   "OK",
			"score": ScoreArr,
			"name":  NameArr,
		})
	})
	engine.GET("/", func(context *gin.Context) {
		ScoreMap := context.QueryMap("score")
		NameMap, ok := context.GetQueryMap("name")
		if ok == false {
			context.JSON(404, gin.H{
				"error": "404",
			})
		}
		context.JSON(200, gin.H{
			"msg":   "OK",
			"score": ScoreMap,
			"name":  NameMap,
		})
	})
	engine.POST("/", func(context *gin.Context) {
		name := context.PostForm("name")
		age := context.DefaultQuery("age", "22")
		home, ok := context.GetPostForm("home")
		if ok == false {
			context.JSON(404, gin.H{
				"error": "404",
			})
		}
		context.JSON(200, gin.H{
			"msg":  "OK",
			"name": name,
			"ago":  age,
			"home": home,
		})
	})
	engine.POST("/", func(context *gin.Context) {
		ScoreArr := context.PostFormArray("score")
		NameArr, ok := context.GetPostFormArray("name")
		if ok == false {
			context.JSON(404, gin.H{
				"error": "404",
			})
		}
		context.JSON(200, gin.H{
			"msg":   "OK",
			"score": ScoreArr,
			"name":  NameArr,
		})
	})
	engine.POST("/", func(context *gin.Context) {
		ScoreMap := context.PostFormMap("score")
		NameMap, ok := context.GetPostFormMap("name")
		if ok == false {
			context.JSON(404, gin.H{
				"error": "404",
			})
		}
		context.JSON(200, gin.H{
			"msg":   "OK",
			"score": ScoreMap,
			"name":  NameMap,
		})
	})
	//参数绑定
	type User struct {
		name string `form:"name" binding:"required"`
		ago  int8   `form:"ago" binding:"required"`
	}
	engine.GET("/", func(context *gin.Context) {
		var user1 User
		err := context.BindQuery(&user1)
		if err != nil {
			context.JSON(500, gin.H{
				"error": "500",
			})
		}
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	
	engine.GET("/", func(context *gin.Context) {
		var user2 User
		err := context.ShouldBindQuery(&user2)
		if err != nil {
			context.JSON(500, gin.H{
				"error": "500",
			})
		}
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	//context.BindJSON(&user)
	//context.ShouldBindJSON(&user)
	//context.BindXML(&user)
	//context.ShouldBindXML(&user)
	//context其他
	//获取路径
	engine.GET("/", func(context *gin.Context) {
		fmt.Println(context.Request.URL.Path)
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	//IP
	engine.GET("/", func(context *gin.Context) {
		fmt.Println(context.ClientIP())
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	//任务
	engine.GET("/", func(context *gin.Context) {
		done := context.Done()
		//业务代码
		<-done
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	//请求头
	engine.GET("/", func(context *gin.Context) {
		head := context.Request.Header
		fmt.Println(head)
		context.JSON(200, gin.H{
			"msg": "ok",
		})
	})
	//cookie
	//获取cookie
	engine.GET("/", func(context *gin.Context) {
		name, err := context.Cookie("name")
		if err != nil {
			return
		}
		fmt.Println("name", name)
		ago, err := context.Cookie("name")
		if err != nil {
			context.JSON(401, gin.H{
				"msg": "err",
			})
		}
		fmt.Println("ago", ago)
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	//设置cookie
	engine.GET("/", func(context *gin.Context) {
		context.SetCookie("name", "sky", 604800, "/", "localhost", false, true)
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	//中间件
	engine.GET("/", middle(), func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	v3 := engine.Group("v3").Use(middle())
	v3.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	newEngine.Use(gin.Logger())
	//文件
	engine.StaticFile("/favicon.ico", "./favicon.icon")
	engine.StaticFS("/static", http.Dir("/static"))
	//错误
	//404
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{
			"error": "404",
		})
	})
	//500
	engine.Use(func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.JSON(500, gin.H{
					"error": "500",
				})
			}
		}()
		context.Next()
	})
	//运行
	if engine.Run("0.0.0.0:80") != nil {
		return
	}
	
}

func middle() gin.HandlerFunc {
	return func(context *gin.Context) {
		//业务代码
		context.Next()
	}
}
