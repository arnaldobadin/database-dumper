package storage

type Connection struct {
	Strg string `json:"storage"`
	Host string `json:"host"`
	User string `json:"user"`
	Pwd string `json:"password"`
	Db string `json:"database"`
}