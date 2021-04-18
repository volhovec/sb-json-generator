package main

import (
	"SnowballDataGenerator/structs"
	"log"
)

func main() {
	log.Println(generateResponse())
}

func generateResponse() string {
	sb := new(structs.Snowball)

	return sb.Generate()
}
