package main

import (
	_ "divination/config"
	_ "divination/database"
	"divination/routes"
)

func main() {

	r := routes.SetupRouter()

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/*")

	r.Run(":5050")
}
