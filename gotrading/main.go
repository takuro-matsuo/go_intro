package main

import (
	"go_intro_mrsakai/gotrading/app/controllers"
	"go_intro_mrsakai/gotrading/config"
	"go_intro_mrsakai/gotrading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
