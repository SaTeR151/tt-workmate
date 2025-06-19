package dto

type HTTPError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type ResponceId struct {
	Id string `json:"id"`
}

type TaskInfo struct {
	Status      string `json:"status"`
	CreateDate  string `json:"create_date"`
	ProcessTime string `json:"process_time"`
}
