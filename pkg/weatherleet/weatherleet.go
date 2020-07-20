package weatherleet

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

const (
	currentExcludes  = "minutely,hourly,daily"
	minutelyExcludes = "current,hourly,daily"
	hourlyExcludes   = "current,minutely,daily"
	dailyExcludes    = "current,minutely,hourly"
)

func New(key, lat, lon string) (*WeatherLeet, error) {
	w := &WeatherLeet{
		key:     key,
		lat:     lat,
		lon:     lon,
		onecall: &onecall{},
	}

	w.url = &url.URL{
		Scheme: "https",
		Host:   "api.openweathermap.org",
		Path:   "data/2.5/onecall",
	}

	q := w.url.Query()
	q.Set("lat", w.lat)
	q.Set("lon", w.lon)
	q.Set("appid", w.key)

	w.url.RawQuery = q.Encode()

	return w, nil
}

func (oc onecall) String() string {
	jsonData, err := json.MarshalIndent(oc, "", " ")
	if err != nil {
		return err.Error()
	}

	return string(jsonData)
}

func getWeather(u url.URL) (*onecall, error) {
	o := &onecall{}

	log.WithField("url", u.String()).Debug("Get weather")

	resp, err := http.Get(u.String())
	if err != nil {
		return o, err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(o)

	return o, nil
}

func (w *WeatherLeet) getUrlExcluding(excludes string) url.URL {

	u := *w.url

	u.RawQuery = u.Query().Encode()

	q := w.url.Query()
	q.Set("exclude", excludes)

	u.RawQuery = q.Encode()

	return u
}

func (w *WeatherLeet) Current() (*onecall, error) {
	oc, err := getWeather(w.getUrlExcluding(currentExcludes))
	if err != nil {
		return nil, err
	}

	return oc, nil
}

func (w *WeatherLeet) Minutely() (*onecall, error) {
	oc, err := getWeather(w.getUrlExcluding(minutelyExcludes))
	if err != nil {
		return nil, err
	}

	return oc, nil
}

func (w *WeatherLeet) Hourly() (*onecall, error) {
	oc, err := getWeather(w.getUrlExcluding(hourlyExcludes))
	if err != nil {
		return nil, err
	}

	return oc, nil
}

func (w *WeatherLeet) Daily() (*onecall, error) {
	oc, err := getWeather(w.getUrlExcluding(dailyExcludes))
	if err != nil {
		return nil, err
	}

	return oc, nil
}
