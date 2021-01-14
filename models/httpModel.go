package models

type Data struct {
	UniqueId string `json:"unique_id"`
	Content  string `json:"content"`
	Callback string `json:"callback"`
}

type Check struct {
	TaskId string   `json:"task_id"`
	Data   []Data   `json:"data"`
	Model  []string `json:"model"`
	Sync   string   `json:"sync"`
}

type ContentResult struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}
