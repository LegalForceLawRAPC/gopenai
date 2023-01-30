package constants

type endpoints map[string]RequestData

// openAiEndpoints is a map of all the endpoints for the OpenAI API
var openAiEndpoints = endpoints{
	// To list all models
	"listModels": {
		Endpoint: "/v1/models",
		Method:   "GET",
		Body:     nil,
	},
}

// dalleEndpoints is a map of all the endpoints for the Dalle API
var dalleEndpoints = endpoints{
	// To generate images using a prompt
	"generateImages": {
		Endpoint: "v1/images/generations",
		Method:   "POST",
		Body:     nil,
	},
	// To edit images using a mask and an image with a prompt
	"editImages": {
		Endpoint:    "v1/images/edits",
		Method:      "POST",
		Body:        nil,
		ContentType: "multipart/form-data",
	},
}

// RequestData is a struct that contains the data required to make a request
type RequestData struct {
	// Endpoint is the endpoint of the request
	Endpoint string
	// Method is the method of the request GET/POST/PATCH etc
	Method string
	// Body accepts an interface{} as the body of the request
	Body interface{}
	// ContentType of the request is optional, if not provided it will be set to application/json
	ContentType string
}

// GetOpenAIEndpoint returns the endpoint data for the OpenAI API
func GetOpenAIEndpoint(endpointType string) RequestData {
	return openAiEndpoints[endpointType]
}

// GetDalleEndpoint returns the endpoint data for the Dalle API
func GetDalleEndpoint(endpointType string) RequestData {
	return dalleEndpoints[endpointType]
}
