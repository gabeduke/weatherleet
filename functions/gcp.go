package functions

import (
	"fmt"
	"github.com/gabeduke/weatherleet/pkg/weatherleet"
	"net/http"
)

type WeatherLeetServer struct {
	*weatherleet.WeatherLeet
}

func New(w *weatherleet.WeatherLeet) *WeatherLeetServer {
	return &WeatherLeetServer{w}
}

func (s *WeatherLeetServer) CurrentWeather(w http.ResponseWriter, r *http.Request) {

	oc, err := s.Current()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, _ = fmt.Fprint(w, oc.String())
}

func (s *WeatherLeetServer) MinutelyWeather(w http.ResponseWriter, r *http.Request) {

	oc, err := s.Minutely()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, _ = fmt.Fprint(w, oc.String())
}

func (s *WeatherLeetServer) HourlyWeather(w http.ResponseWriter, r *http.Request) {

	oc, err := s.Hourly()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, _ = fmt.Fprint(w, oc.String())
}

func (s *WeatherLeetServer) DailyWeather(w http.ResponseWriter, r *http.Request) {

	oc, err := s.Daily()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, _ = fmt.Fprint(w, oc.String())
}
