package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Weather provides accurate weather details
type Weather struct {
	Temp    float32
	Weather string
}

type client struct {
	url string
}

// New creates a new weather client
func New(url string) (*client, error) {
	return &client{
		url: url,
	}, nil
}

// GetWeather returns the weather information for the provided location
func (c *client) GetWeather(location string) (Weather, error) {
	url := fmt.Sprintf("%s?location=%s", c.url, location)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Weather{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Weather{}, err
	}
	defer resp.Body.Close()

	b := struct {
		Main          string  `json:"main"`
		Description   string  `json:"description"`
		Location      string  `json:"location"`
		TempMin       float32 `json:"temp_min"`
		TempMax       float32 `json:"temp_max"`
		WindSpeed     int     `json:"wind_speed"`
		WindDirection string  `json:"wind_direction"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&b); err != nil {
		return Weather{}, nil
	}

	return Weather{
		Temp:    (b.TempMin / b.TempMax) * 100,
		Weather: b.Main,
	}, nil
}
