package main

import (
	"fmt"
	"logistics/constants"
	"logistics/libs/conf"
	"logistics/libs/databases"
	"logistics/libs/logs"
	mymiddleware "logistics/libs/middleware"
	"logistics/libs/redis"
	"logistics/libs/wechat"
	"logistics/service/common"
	service_wechat "logistics/service/wechat"
	"os"
	"os/signal"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func syncRedis() bool {
	return true
}

func Initialize() {
	// todo: 初始化配置
	if _pwd, err := os.Getwd(); err != nil {
		panic("get current work dir error")
	} else {
		fmt.Println("Current working directory:", _pwd)
	}
	isDebug := os.Getenv("DEBUG") == "true"
	var configFile string
	if isDebug {
		configFile = "./conf/dev.toml"
	} else {
		configFile = "./conf/prod.toml"
	}
	// 设置运行配置文件
	os.Setenv("runConfig", configFile)
	conf.Init()
	// 读取环境变量
	constants.Init()
	// todo: 初始化日志
	logs.Init(conf.Get("logger"))
	// todo: 初始化Mysql
	databases.Init(conf.Get("mysql"))
	// todo: 同步数据库结构
	//if ok := common.DatabaseSync(); !ok {
	//	panic("database sync error")
	//}
	// todo: 初始化redis
	redis.Init(conf.Get("redis"))
	// todo: 初始化wxClient
	wechat.Init(conf.Get("wechat"))

	// todo: 同步数据到redis
	if !syncRedis() {
		panic("redis init error")
	}
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		127.0.0.1:10001
// @BasePath	/v1
func RunServer(e *echo.Echo, wg *sync.WaitGroup) {
	defer wg.Done()

	e.Validator = &common.CustomValidator{Validator: validator.New()}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Recover())
	e.Use(mymiddleware.Cors()) // 开发环境，允许跨域
	basePath := e.Group("/v1")

	r := basePath.Group("/test")
	r.Use(mymiddleware.Request()) // 使用日志中间件
	r.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Hello, World!"})
	})

	wx := basePath.Group("/wx")
	r.Use(mymiddleware.Request())
	wx.POST("/code2session", service_wechat.WxCode2Session)

	user := basePath.Group("/users")
	user.Use(mymiddleware.Request(), mymiddleware.Access(), middleware.CORS())
	//user.POST("/getAddressBooks", service_users.GetAddressBooks)

	//打印路由信息
	for _, route := range e.Routes() {
		println(route.Method, route.Path)
	}

	addr := fmt.Sprintf("%s:%s", os.Getenv("WEB_HOST"), os.Getenv("WEB_PORT"))
	e.Logger.Fatal(e.Start(addr))
}

func RunTask(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			// 这里可以记录日志或执行其他操作
			fmt.Println("Recovered from panic in RunTask:", r)
		}
	}()
}

func main() {
	// 初始化中间件
	Initialize()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	s := echo.New()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		// 启动web server
		go RunServer(s, wg)
		// 启动任务
		go RunTask(wg)
	}()

	<-sigCh // 等待中断信号
	fmt.Println("Received interrupt signal, shutting down...")
	// todo:这里可以执行清理操作
	wg.Wait()
}
