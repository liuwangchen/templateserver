package web

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"templateserver/pkg/web/controller"
	"time"
)

var (
	As *http.Server
	Bs *http.Server
)

func init() {
	aport := flag.Int("aport", 8080, "admin port")
	bport := flag.Int("bport", 9090, "business port")
	//gin.SetMode(gin.ReleaseMode)
	ar := gin.Default()
	controller.InitAdminRoute(ar)
	As = &http.Server{
		Addr:           fmt.Sprintf(":%v", *aport),
		Handler:        ar,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	br := gin.Default()
	controller.InitBusRoute(br)
	Bs = &http.Server{
		Addr:           fmt.Sprintf(":%v", *bport),
		Handler:        br,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
