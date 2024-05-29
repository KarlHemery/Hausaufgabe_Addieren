package models

type AddRequest struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

type AddResponse struct {
	Sum float64 `json:"sum"`
}
