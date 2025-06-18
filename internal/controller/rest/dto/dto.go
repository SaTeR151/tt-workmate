package dto

type HTTPError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type ResponceId struct {
	Id string `json:"id"`
}
