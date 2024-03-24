package weather

import (
	"fmt"
	"io"
	"net/http"
)

func GetCurrentWeather(city string) (string, error) {
	resp, err := http.Get(getWeatherServiceURL(city))
	if err != nil {
		return "", fmt.Errorf("doing http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("got unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body: %w", err)
	}

	return string(body), nil
}

func getWeatherServiceURL(city string) string {
	return fmt.Sprintf("https://wttr.in/%s?format=2", city)
}
