package models

type InputData struct {
	Day    int  `json:"day"`
	Month  int  `json:"month"`
	Year   int  `json:"year"`
	Hour   int  `json:"hour"`
	Gender int  `json:"gender"`
	IsSun  bool `json:"isSun"`
}
