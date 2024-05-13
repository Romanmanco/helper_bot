package client

type WeatherData struct {
	Name    string      `json:"name"`
	Main    MainWeather `json:"main"`
	Wind    Wind        `json:"wind"`
	Weather []Weather   `json:"weather"`
}

type MainWeather struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
