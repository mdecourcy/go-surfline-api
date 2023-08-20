# SurflineAPI Go Package

## Overview

The `surflineapi` Go package facilitates seamless interactions with the Surfline API, allowing developers to fetch various forecast data including wave, wind, sunlight, tides, weather, and nearby buoys' information. It's designed to simplify the process of integrating Surfline data into your Go applications.

## Features

- **Direct API Interactions**: Easily fetch data directly from Surfline with the provided methods.
- **Built-in HTTP Client**: Inbuilt HTTP client to manage requests, with the ability to use your custom client if required.
- **Strongly Typed Responses**: Enjoy type-safe responses for all the API interactions.
- **Comprehensive Coverage**: Covering a vast range of forecasts from regions, spots, and buoy data.

## Installation

Fetch the package using `go get`:

```bash
go get github.com/mdecourcy/go-surfline-api
```

## Usage

1. Import the package:

```go
import "github.com/mdecourcy/go-surfline-api"
```

2. Initialize the API client and make calls:

```go
apiClient := &surflineapi.SurflineAPI{
    HTTPClient: http.DefaultClient,
}

// Example: Fetching Wave Forecast for a spot
forecast, err := apiClient.GetWaveForecast("spotId", 7, 1)
if err != nil {
    log.Fatal(err)
}
// Process the forecast data...
```

## Available Methods

- `GetRegionsForecastConditions(subregionId string, days int)`: Fetches regional forecast conditions.
- `GetSpotForecastRating(spotId string, days int, intervalHours int)`: Fetches forecast ratings for specific spots.
- `GetWaveForecast(spotId string, days int, intervalHours int)`: Retrieves wave forecast data.
- `GetWindForecast(spotId string, days int, intervalHours int, corrected bool, cacheEnabled bool)`: Fetches wind forecast data.
- `GetSunlightForecast(spotId string, days int, intervalHours int)`: Obtains sunlight forecast data.
- `GetTideForecast(spotId string, days int)`: Grabs tide forecast data.
- `GetWeatherForecast(spotId string, days int)`: Retrieves weather forecast data.
- `GetNearbyBuoys(latitude, longitude float64, limit, distance int)`: Fetches nearby buoys based on given location parameters.

## Contributing

Contributions are highly encouraged! Open an issue, first, to start a discussion.


---

May your code ride the perfect digital wave! 🌊🏄‍♂️🔧