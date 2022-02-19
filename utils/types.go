package utils

type Any = interface{}
type AnyStruct = map[string]Any

type MessageResponse struct {
	Message string `json:"message"`
}

type TokenResponse struct {
	MessageResponse
	Token string `json:"token"`
}
