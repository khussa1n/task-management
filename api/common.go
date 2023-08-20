package api

type Error struct {
	Message string `json:"message"`
}

type OK struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
