package main

import (
	"fmt"
	"net/http"

	surflineapi "github.com/mdecourcy/go-surfline-api/pkg/surflineapi"
)

func main() {
	client := &http.Client{}
	api := &surflineapi.SurflineAPI{
		HTTPClient: client,
	}

	conditions, err := api.GetRegionsForecastConditions("58581a836630e24c4487900d", 5)
	if err != nil {
		fmt.Println("Error fetching forecast conditions:", err)
		return
	}
	fmt.Println(conditions)

	ratings, err := api.GetSpotForecastRating("5842041f4e65fad6a7708841", 5, 1)
	if err != nil {
		fmt.Println("Error fetching spot forecast rating:", err)
		return
	}
	fmt.Println(ratings)

	waveForcast, err := api.GetWaveForecast("5842041f4e65fad6a7708841", 5, 1)
	if err != nil {
		fmt.Println("Error fetching wave forecast:", err)
		return
	}
	fmt.Println(waveForcast)

	weatherForecast, err := api.GetWeatherForecast("5842041f4e65fad6a7708841", 5)
	if err != nil {
		fmt.Println("Error fetching weather forecast:", err)
		return
	}
	fmt.Println(weatherForecast)

	tideForecast, err := api.GetTideForecast("5842041f4e65fad6a7708841", 5)
	if err != nil {
		fmt.Println("Error fetching tide forecast:", err)
		return
	}
	fmt.Println(tideForecast)

	GetNearbyBuoys, err := api.GetNearbyBuoys(32.820224805799526, -117.2627870480912, 10, 200)
	if err != nil {
		fmt.Println("Error fetching nearby buoys:", err)
		return
	}
	fmt.Println(GetNearbyBuoys)

}
