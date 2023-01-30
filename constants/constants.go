package constants

var BaseURL = "https://api.openai.com"
var apiKey = ""
var orgId = ""

// SetApiKey sets the API key for the OpenAI API
func SetApiKey(key string) {
	apiKey = key
}

// SetOrgId sets the organization ID for the OpenAI API
func SetOrgId(id string) {
	orgId = id
}

// GetApiKey returns the API key for the OpenAI API
func GetApiKey() string {
	return apiKey
}

// GetOrgId returns the organization ID for the OpenAI API
func GetOrgId() string {
	return orgId
}

// GetToken returns the Bearer token for the OpenAI API
func GetToken() string {
	return "Bearer " + apiKey
}
