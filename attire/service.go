package attire

import (
	"errors"
	"whistlestop.com/weather"
)

// Forecaster predicts the weather
type Forecaster interface {
	GetWeather(location string) (weather.Weather, error)
}

// Attire is the recommended attire
type Attire struct {
	Jacket struct {
		Waterproof bool
		Windproof  bool
	}
	Pants    string
	Umbrella bool
}

type Service struct {
	f Forecaster
}

func New(f Forecaster) (*Service, error) {
	if f == nil {
		return nil, errors.New("invalid forecaster")
	}

	return &Service{
		f: f,
	}, nil
}

func (s *Service) Recommend(location string) (Attire, error) {
	w, err := s.f.GetWeather(location)
	if err != nil {
		return Attire{}, err
	}

	// Set some defaults
	a := Attire{
		Umbrella: false,
	}

	if w.Weather == "Rain" {
		a.Umbrella = true
	}

	if w.Temp < 20.0 {
		a.Pants = "trousers"
		a.Jacket = struct {
			Waterproof bool
			Windproof  bool
		}{
			Waterproof: true,
			Windproof:  true,
		}
	} else {
		a.Pants = "shorts"
		a.Jacket = struct {
			Waterproof bool
			Windproof  bool
		}{
			Waterproof: true,
			Windproof:  false,
		}
	}

	return a, nil
}
