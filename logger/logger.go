package logger

import (
	"log"

	"github.com/Rus203/shop/config"
)

func Log(message any) {
	logType := configs.Env.Log


	if logType == "prod" {
		return
	}

	if logType == "debug" {
		log.Println(message)
		return
	}

	log.Fatal("unknow log type")

}

func Panic (message any) {
	log.Panicln(message)
}