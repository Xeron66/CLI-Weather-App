package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Weather struct {
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentUnits         struct {
		TemperatureUnit         string `json:"temperature_2m"`
		RelativeHumidityUnit    string `json:"relative_humidity_2m"`
		ApparentTemperatureUnit string `json:"apparent_temperature"`
		PrecipitationUnit       string `json:"precipitation"`
		RainUnit                string `json:"rain"`
		WindSpeedUnit           string `json:"wind_speed_10m"`
		WindDirectionUnit       string `json:"wind_direction_10m"`
	} `json:"current_units"`
	Dailyunits struct {
		TemperatureUnit      string `json:"temperature_2m_max"`
		DayLightDurationUnit string `json:"daylight_duration"`
		WindSpeedUnit        string `json:"wind_speed_10m_max"`
		PrecipitationUnit    string `json:"precipitation_probability_max"`
	} `json:"daily_units"`
	CurrentWeather struct {
		Time                string  `json:"time"`
		Temperature         float64 `json:"temperature_2m"`
		RelativeHumidity    float64 `json:"relative_humidity_2m"`
		ApparentTemperature float64 `json:"apparent_temperature"`
		Precipitation       float64 `json:"precipitation"`
		RainProbability     float64 `json:"rain"`
		WindSpeed           float64 `json:"wind_speed_10m"`
		WindDirection       float64 `json:"wind_direction_10m"`
	} `json:"current"`
	DailyWeather struct {
		Date                    []string  `json:"time"`
		Temperature             []float64 `json:"temperature_2m_max"`
		SunRise                 []string  `json:"sunrise"`
		SunSet                  []string  `json:"sunset"`
		DayLightDuration        []float64 `json:"daylight_duration"`
		UvIndex                 []float64 `json:"uv_index_max"`
		WindSpeed               []float64 `json:"wind_speed_10m_max"`
		PrecipitationProbabiliy []float64 `json:"precipitation_probability_max"`
	} `json:"daily"`
}

func main() {
	res, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=23.7104&longitude=90.4074&daily=temperature_2m_max,sunrise,sunset,daylight_duration,uv_index_max,wind_speed_10m_max,precipitation_probability_max&current=temperature_2m,relative_humidity_2m,apparent_temperature,is_day,precipitation,rain,wind_speed_10m,wind_direction_10m&timezone=auto&forecast_days=1")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API Not Available!")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))

	var w Weather
	err = json.Unmarshal(body, &w)
	if err != nil {
		panic(err)
	}
	// Units
	timezone := w.Timezone
	time_zone_abb := w.TimezoneAbbreviation
	elevation := w.Elevation
	temp_unit := w.CurrentUnits.TemperatureUnit
	rel_humidity_unit := w.CurrentUnits.RelativeHumidityUnit
	precipitation_unit := w.CurrentUnits.PrecipitationUnit
	rain_unit := w.CurrentUnits.RainUnit
	wind_speed_unit := w.CurrentUnits.WindSpeedUnit
	wind_direction_unit := w.CurrentUnits.WindDirectionUnit

	// Values
	curr_time := strings.Split(w.CurrentWeather.Time, "T")[1]
	curr_temp := w.CurrentWeather.Temperature
	curr_rel_humid := w.CurrentWeather.RelativeHumidity
	curr_app_temp := w.CurrentWeather.ApparentTemperature
	curr_precip := w.CurrentWeather.Precipitation
	curr_rain_prob := w.CurrentWeather.RainProbability
	curr_wind_speed := w.CurrentWeather.WindSpeed
	curr_wind_dir := w.CurrentWeather.WindDirection

	d_date := w.DailyWeather.Date

	fmt.Printf("\nTimeZone: %q, %q\nElevation: %.2f", timezone, time_zone_abb, elevation)
	fmt.Printf("\nDate: %q, Time: %q", d_date[0], curr_time)
	fmt.Printf("\nCurrent Temp: %.2f%q", curr_temp, temp_unit)
	fmt.Printf("\nRelative Humidity: %.2f%q", curr_rel_humid, rel_humidity_unit)
	fmt.Printf("\nApparent Temp: %.2f%q", curr_app_temp, temp_unit)
	fmt.Printf("\nPrecipitation: %.2f%q", curr_precip, precipitation_unit)
	fmt.Printf("\nRain Probability: %.2f%q", curr_rain_prob, rain_unit)
	fmt.Printf("\nWind Speed: %.2f%q", curr_wind_speed, wind_speed_unit)
	fmt.Printf("\nWind Direction: %.2f%q", curr_wind_dir, wind_direction_unit)

}
