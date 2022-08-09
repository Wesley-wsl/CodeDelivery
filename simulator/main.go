package main

import (
	// "fmt"
	"log"

	// route2 "github.com/codeedu/imersaofsfc2-simulator/application/route"
	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}

func main() {

	producer := kafka.NewKafkaProducer()

	kafka.Publish("Hello World!", "readtest", producer)

	for {
		_ = 1
	}

	// route := route2.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringjson, _ := route.ExportJsonPositions()

	// fmt.Println(stringjson[0])
}
