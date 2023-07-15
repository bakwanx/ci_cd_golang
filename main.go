package main

import (
	"ci_cd_golang/config"
	"ci_cd_golang/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8080")//
}
