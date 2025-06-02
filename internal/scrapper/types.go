package scrapper

type Task struct {
	Id     int64  `json:"id"`
	Url    string `json:"url"`
	Method string `json:"method"`
}

type Result struct {
	Id         int64  `json:"id"`
	Url        string `json:"url"`
	BodyLength int64  `json:"body_length"`
	WorkerId   int64  `json:"worker_id"`
	Err        string `json:"err"`
}
