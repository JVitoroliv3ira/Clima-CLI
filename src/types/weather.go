package types

type Weather struct {
	City    string `json:"name"`
	Country struct {
		Name string `json:"country"`
	} `json:"sys"`
	Info []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Climate struct {
		Temp      float64 `json:"temp"`
		Sensation float64 `json:"feels_like"`
		Max       float64 `json:"temp_max"`
		Min       float64 `json:"temp_min"`
	} `json:"main"`
}
