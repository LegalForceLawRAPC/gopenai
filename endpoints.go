package gopenai

type endpoints map[string]RequestData

var openAiEndpoints = endpoints{
	"listModels": {
		endpoint: "/v1/models",
		method:   "GET",
		body:     nil,
	},
}
