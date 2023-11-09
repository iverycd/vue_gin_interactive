package router

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"vue_gin_interactive/handlers"
)

func Init() {
	r := gin.Default()
	// 用来给前端展示的响应数据
	r.GET("test", func(c *gin.Context) {
		// 设置服务端的请求头,否则前端无法获取后端的数据
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{"data": "backend,hello i am backend message!"})
		//c.HTML(http.StatusOK, "index.html", gin.H{"title": "hello 我是内容"})
	})
	// v1分组示例路由
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", handlers.HelloPage)
		v1.GET("/line", func(c *gin.Context) {
			legendData := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
			xAxisData := []int{120, 240, rand.Intn(500), rand.Intn(500), 150, 230, 180}
			c.JSON(200, gin.H{
				"legend_data": legendData,
				"xAxis_data":  xAxisData,
			})
		})
	}
	// 加载html静态页面
	r.LoadHTMLGlob("templates/*")
	// v2分组示例路由 使用/v2/index 访问index.html静态页面
	v2 := r.Group("/v2")
	{
		v2.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{"title": "hello Gin this is content ."})
		})
	}
	// 404 NotFound
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}
