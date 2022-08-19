package controller

import (
	"github.com/labstack/echo"
	"go-test/src/service"
	"net/http"
)

func FileController(group *echo.Group) {

	g := group.Group("")

	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// AWS S3 파일 업로드 테스트
	g.POST("upload", service.FileUpload())
}
