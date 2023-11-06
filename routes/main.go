package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiKey = "bbf2679e1f694a0594a140321230311&q"

type WeatherResponse struct {
	Location LocationData `json:"location"`
	Current  CurrentData  `json:"current"`
}

type LocationData struct {
	Name string `json:"name"`
}

type CurrentData struct {
	temp float64 `json:"temp_c"`
}

func getWeatherForecast(apiKey, location string) (*WeatherResponse, error) {
	url := "https://api.weatherapi.com/v1/forecast.json?key=" + apiKey + "&q=" + location

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var response WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func GetWeatherHandler(c *gin.Context) {
	location := c.Param("location")
	weatherData, err := getWeatherForecast(apiKey, location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := struct {
		City        string
		Temperature float64
	}{
		City:        weatherData.Location.Name,
		Temperature: weatherData.Current.temp,
	}

	// parsing data to template
	err = template.Must(template.ParseFiles("templates/index.html")).Execute(c.Writer, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

type Route struct {
	Method  string
	Path    string
	Handler func(c *gin.Context)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	routes := []Route{
		{
			Method:  "GET",
			Path:    "/weather/:location",
			Handler: GetWeatherHandler,
		},
	}
	for _, route := range routes {
		r.Handle(route.Method, route.Path, route.Handler)
	}

	fmt.Println("Server is running on port 8080...")
	r.Run(":8080")
}
