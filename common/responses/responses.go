package responses

type Success[T any] struct {
	Status bool `json:"status"` // true
	Data   T    `json:"data"`
}

type Error struct {
	Status  bool   `json:"status"` // false
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Empty struct {
	Status bool `json:"status"` // true
}
