package lib

type Response struct {
	Message string `json:"message"`
	Number  string `json:"tribonacci"`
}

type ConfigStructure struct {
	Port   string `json:"port"`
	Router string `json:"router"`
	Redis string `json:"redis"`
}
