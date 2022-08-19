package main

import (
	"go-test/src/controller"
	"log"
	"net"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // 적용하면 자동으로 콘솔에 context정보 로깅 하는듯
	e.Use(middleware.Recover()) // 패닉으로 프로그램 종료되는것 방지 해주는듯
	setHandler(e)
	// Route => handler
	//e.GET("/", func(c echo.Context) error {
	//	//fmt.Println("111")
	//	//panic("PANIC")
	//	//fmt.Println("222")
	//	return c.String(http.StatusOK, "Hello, World!\n")
	//})

	// Start server
	//e.Logger.Fatal(e.Start(":1323"))
	// Http server (echo)
	listener, err := net.Listen("tcp", ":"+"1323")
	if err != nil {
		log.Fatal(err)
	}
	e.Server.ReadTimeout = 20 * time.Minute
	e.Server.WriteTimeout = 20 * time.Minute
	e.Server.IdleTimeout = 120 * time.Second

	e.Listener = listener
	e.Logger.Fatal(e.StartServer(e.Server))
}

func setHandler(e *echo.Echo) {
	g := e.Group("/")

	controller.FileController(g)
}
