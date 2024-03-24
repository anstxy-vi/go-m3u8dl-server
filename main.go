package main

import (
	"fmt"
	"gorman/m3u8dl/config"
	"gorman/m3u8dl/initial"
	"gorman/m3u8dl/routes"
)

func main() {

	initial.InitLogger()

	router := routes.Routes()

	router.Run(fmt.Sprintf("0.0.0.0:%d", config.Global.Port))
}
