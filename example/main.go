package main

import (
	"github.com/LegalForceLawRAPC/gopenai"
	"log"
)

func main() {
	cl := gopenai.NewClient()
	err := cl.Connect("YOUR_API_KEY",
		"YOUR_ORG_ID")
	if err != nil {
		log.Panicln(err)
	}
	data := cl.ListModels()

	for _, el := range data.Data {
		log.Println(el)
	}
}
