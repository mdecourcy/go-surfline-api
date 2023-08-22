package surflineapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://services.surfline.com/kbyg"

type SurflineAPI struct {
	HTTPClient *http.Client
}

// Core HTTP fetcher
func (api *SurflineAPI) fetch(url string, v interface{}) error {
	resp, err := api.HTTPClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}

// Regions Forecast Conditions
func (api *SurflineAPI) GetRegionsForecastConditions(subregionId string, days int) (*RegionsForecastConditionsResponse, error) {
	url := fmt.Sprintf("%s/regions/forecasts/conditions?subregionId=%s&days=%d", baseURL, subregionId, days)
	var forecastResponse RegionsForecastConditionsResponse
	err := api.fetch(url, &forecastResponse)
	return &forecastResponse, err
}

// Spot Forecast Rating
func (api *SurflineAPI) GetSpotForecastRating(spotId string, days int, intervalHours int) (*SpotForecastRatingResponse, error) {
	url := fmt.Sprintf("%s/spots/forecasts/rating?spotId=%s&days=%d&intervalHours=%d&cacheEnabled=true", baseURL, spotId, days, intervalHours)
	var ratingResponse SpotForecastRatingResponse
	err := api.fetch(url, &ratingResponse)
	return &ratingResponse, err
}

// Wave Forecast
func (api *SurflineAPI) GetWaveForecast(spotId string, days int, intervalHours int) (*WaveForecastResponse, error) {
	url := fmt.Sprintf("%s/spots/forecasts/wave?spotId=%s&days=%d&intervalHours=%d&cacheEnabled=true&units[swellHeight]=FT&units[waveHeight]=FT", baseURL, spotId, days, intervalHours)
	var forecastResponse WaveForecastResponse
	err := api.fetch(url, &forecastResponse)
	return &forecastResponse, err
}

// Wind Forecast
func (api *SurflineAPI) GetWindForecast(spotId string, days int, intervalHours int, corrected bool, cacheEnabled bool) (*WindForecastResponse, error) {
	url := fmt.Sprintf("%s/spots/forecasts/wind?spotId=%s&days=%d&intervalHours=%d&corrected=%t&cacheEnabled=%t&units%%5BwindSpeed%%5D=KTS",
		baseURL, spotId, days, intervalHours, corrected, cacheEnabled)
	var windForecastResponse WindForecastResponse
	err := api.fetch(url, &windForecastResponse)
	return &windForecastResponse, err
}

// Sunlight Forecast
func (api *SurflineAPI) GetSunlightForecast(spotId string, days int, intervalHours int) (*SunlightForecastResponse, error) {
	url := fmt.Sprintf("%s/spots/forecasts/sunlight?spotId=%s&days=%d&intervalHours=%d", baseURL, spotId, days, intervalHours)
	var sunlightResponse SunlightForecastResponse
	err := api.fetch(url, &sunlightResponse)
	return &sunlightResponse, err
}

// Tide Forecast
func (api *SurflineAPI) GetTideForecast(spotId string, days int) (*TideForecastResponse, error) {
	url := fmt.Sprintf("%s/spots/forecasts/tides?spotId=%s&days=%d&cacheEnabled=true&units%%5BtideHeight%%5D=FT", baseURL, spotId, days)
	var tideForecastResponse TideForecastResponse
	err := api.fetch(url, &tideForecastResponse)
	return &tideForecastResponse, err
}

// Weather Forecast
func (api *SurflineAPI) GetWeatherForecast(spotId string, days int) (*WeatherForecastResponse, error) {
	url := fmt.Sprintf("%s/spots/forecasts/weather?spotId=%s&days=%d&intervalHours=1&cacheEnabled=true&units%%5Btemperature%%5D=F", baseURL, spotId, days)
	var weatherForecastResponse WeatherForecastResponse
	err := api.fetch(url, &weatherForecastResponse)
	return &weatherForecastResponse, err
}

// Nearby Buoys
func (api *SurflineAPI) GetNearbyBuoys(latitude, longitude float64, limit, distance int) (*NearbyBuoysResponse, error) {
	url := fmt.Sprintf("%s/buoys/nearby?cacheEnabled=true&units%%5BswellHeight%%5D=FT&latitude=%f&longitude=%f&limit=%d&distance=%d", baseURL, latitude, longitude, limit, distance)
	var buoysResponse NearbyBuoysResponse
	err := api.fetch(url, &buoysResponse)
	return &buoysResponse, err
}

func (api *SurflineAPI) GetTaxonomyDetails(taxonomyId string) (*TaxonomyDetailsResponse, error) {
	url := fmt.Sprintf("https://services.surfline.com/taxonomy?type=taxonomy&id=%s&maxDepth=0", taxonomyId)
	var taxonomyResponse TaxonomyDetailsResponse
	err := api.fetch(url, &taxonomyResponse)
	return &taxonomyResponse, err
}
