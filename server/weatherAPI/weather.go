package weatherAPI

import (
	"encoding/json"
	"fmt"
	"heimigo/server/helpers"
	"io/ioutil"
	"net/http"
	"time"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float32 `json:"temp"`
	} `json:"main"`
	Sys struct {
		Sunrise int32 `json:"sunrise"`
		Sunset  int32 `json:"sunset"`
	} `json:"sys"`
	Weather       []map[string]interface{} `json:"weather"`
	ActualWeather string
}

func (w WeatherData) Print() string {
	return fmt.Sprintf(`\nTemperature %f 
	city: %s 
	sunset %d 
	sunrise %d 
	Weather %v
	Weatheract %s`,
		w.Main.Temp, w.Name, w.Sys.Sunset, w.Sys.Sunrise, w.Weather[0]["main"], w.ActualWeather)
}
func (w *WeatherData) getActualWeather() {
	// Extract the weatherdata eg what weather is it in cleartext
	w.ActualWeather = w.Weather[0]["main"].(string)
}

func (w WeatherData) CheckSunIsUp() bool {
	t := time.Now().Unix()
	if t > int64(w.Sys.Sunrise) && t < int64(w.Sys.Sunset) {
		return true
	} else {
		return false
	}
}

func ReadWeather() WeatherData {
	var API_key string = "79793771a515cf5b843ba5652129affa"
	var BASE_WEATHER_API_URL string = "http://api.openweathermap.org/data/2.5/weather"
	var url = buildQuery("Kalmar", API_key, BASE_WEATHER_API_URL)
	w := getWeather(url)
	return w
}
func buildQuery(city string, api string, url string) string {
	return fmt.Sprintf("%s?q=%s&units=%s&appid=%s", url, city, "metric", api)
}

func getWeather(url string) WeatherData {
	var w WeatherData

	req, err := http.Get(url)
	helpers.CheckErr(err)
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	fmt.Printf("Here is body %s\n", body)
	helpers.CheckErr(err)

	err = json.Unmarshal(body, &w)
	helpers.CheckErr(err)

	w.getActualWeather()

	// fmt.Printf("Here is w %s\n", w.Print())
	// fmt.Println(w.Main.Temp)
	//fmt.Println(w["name"])
	// fmt.Println(w.Name)
	return w

}
