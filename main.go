package main

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/todaychiji/demo/routers"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	beego.BConfig.Listen.HTTPPort = port
	beego.Run()
}
