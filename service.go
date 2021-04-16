package main

import (
	"SnowballDataGenerator/structs"
	"log"
)

func main() {
	log.Println(generateResponse())
}

func generateResponse() float64 {
	sb := new(structs.Snowball)
	return sb.Generate()
}
