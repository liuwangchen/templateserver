package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"templateserver/config"
	"templateserver/pkg/web/controller"
	"time"
)

var (
	As *http.Server
	Bs *http.Server
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	ar := gin.Default()
	//初始化admin route
	controller.InitAdminRoute(ar)
	As = &http.Server{
		Addr:           fmt.Sprintf(":%v", config.APort),
		Handler:        ar,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	br := gin.Default()
	//初始化业务route
	controller.InitBusRoute(br)
	Bs = &http.Server{
		Addr:           fmt.Sprintf(":%v", config.BPort),
		Handler:        br,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
