package main

import (
	"ci_cd/config"
	"ci_cd/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8080")
}
