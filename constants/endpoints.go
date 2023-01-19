package constants

type endpoints map[string]RequestData

var openAiEndpoints = endpoints{
	"listModels": {
		Endpoint: "/v1/models",
		Method:   "GET",
		Body:     nil,
	},
}

var dalleEndpoints = endpoints{
	"generateImages": {
		Endpoint: "v1/images/generations",
		Method:   "POST",
		Body:     nil,
	},
	"editImages": {
		Endpoint:    "v1/images/edits",
		Method:      "POST",
		Body:        nil,
		ContentType: "multipart/form-data",
	},
}

type RequestData struct {
	Endpoint    string
	Method      string
	Body        interface{}
	ContentType string
}

func GetOpenAIEndpoint(endpointType string) RequestData {
	return openAiEndpoints[endpointType]
}

func GetDalleEndpoint(endpointType string) RequestData {
	return dalleEndpoints[endpointType]
}
