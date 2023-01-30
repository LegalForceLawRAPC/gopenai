# GopenAI (Go Wrapper to work with OpenAI APIs)

## Structuring

```
|--- gopenapi.go
|--- errors
|       |--- http_errors.go
|       |--- openai_errors.go
|       |--- structs.go
|--- endpoints.go
|--- structs.go
|--- README.md
|--- Dalle2
|
|
|
|
|--- GPT


```

## Installation

Initialize a go module and run `go get -u github.com/LegalForceLawRAPC/gopenai`.

## Examples

1. Listing models
```
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

```

2. Generating Image with Dalle
```
func main() {
	cl := gopenai.NewClient()
	err := cl.Connect("YOUR_API_KEY", "YOUR_ORG_ID")
	if err != nil {
		log.Panicln(err)
	}

	res, err := cl.Dalle().GenerateImages("YOUR_PROMPT", 1, "1024x1024", "")
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res.Data)
}
```