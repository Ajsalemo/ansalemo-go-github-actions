package main

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	fmt.Println("beego")
	web.Run()
}