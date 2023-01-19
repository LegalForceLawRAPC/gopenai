package constants

var BaseURL = "https://api.openai.com"
var apiKey = ""
var orgId = ""

func SetApiKey(key string) {
	apiKey = key
}

func SetOrgId(id string) {
	orgId = id
}

func GetApiKey() string {
	return apiKey
}

func GetOrgId() string {
	return orgId
}

func GetToken() string {
	return "Bearer " + apiKey
}
