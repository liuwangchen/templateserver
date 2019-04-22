package main

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"templateserver/pkg/web"
)

func main() {
	err := gracehttp.Serve(web.As, web.Bs)
	if err != nil {
		fmt.Println(err)
	}
}
