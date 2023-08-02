package api

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type OK struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
