package attire_test

import (
	"testing"
	"whistlestop.com/attire"
	"whistlestop.com/weather"
)

type weatherMock struct{}

func (wm weatherMock) GetWeather(location string) (weather.Weather, error) {
	if location == "Manchester" {
		return weather.Weather{
			Weather: "Rain",
			Temp:    21.0,
		}, nil
	}
	return weather.Weather{}, nil
}

func Test_Recommend(t *testing.T) {
	t.Run("should return umbrella with rain", func(t *testing.T) {
		attireService, err := attire.New(weatherMock{})
		if err != nil {
			t.Fatal(err)
		}

		someLocation := "Manchester"

		attire, err := attireService.Recommend(someLocation)
		if err != nil {
			t.Fatal(err)
		}

		if attire.Umbrella != true {
			t.Fatal("no umbrella")
		}
	})
	t.Run("should return jacket in correct conditions", func(t *testing.T) {
		tt := []struct {
			location string
			expected attire.Attire
		}{
			{
				location: "Manchester",
				expected: attire.Attire{
					Jacket: struct {
						Waterproof bool
						Windproof  bool
					}{
						Waterproof: true,
						Windproof:  false,
					},
					Pants:    "shorts",
					Umbrella: true,
				},
			},
		}
		attireService, err := attire.New(weatherMock{})
		if err != nil {
			t.Fatal(err)
		}
		for _, tc := range tt {
			attire, err := attireService.Recommend(tc.location)
			if err != nil {
				t.Fatal(err)
			}

			if attire != tc.expected {
				t.Fatal("mismatch")
			}
		}
	})
}
