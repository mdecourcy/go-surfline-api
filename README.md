# SurflineAPI Go Package

## Overview

The `surflineapi` Go package provides type definitions to work with the Surfline API. It includes type structures for various forecast data, including wave, wind, sunlight, tides, and more. Whether you're building a surf forecast app or analyzing surf data, this package will help streamline your integration with the Surfline API.

## Features

- **Comprehensive Type Definitions**: Covers all the expected responses from the Surfline API, making data parsing and usage a breeze.
- **Clean & Organized**: Structs are organized based on their respective forecast data type to make it intuitive and easy to navigate.
- **Flexibility**: With clearly defined types, it's easy to extend or modify the package based on unique requirements.

## Installation

Use `go get` to fetch the package:

```bash
go get github.com/mdecourcy/go-surfline-api
```

## Usage

1. Import the package:

```go
import "github.com/yourusername/surflineapi"
```

2. Call the provided functions:
```go
conditions, err := api.GetRegionsForecastConditions("58581a836630e24c4487900d", 5)
	if err != nil {
		fmt.Println("Error fetching forecast conditions:", err)
		return
	}
	fmt.Println(conditions)
```

3. Use the fetched data as per your application's needs.

## Structs Included

- `LatLng`: General type for latitude and longitude.
- `Units`: Measurement units used in forecasts.
- `RegionsForecastConditionsResponse`: Regional forecast conditions.
- `SpotForecastRatingResponse`: Rating for specific spots.
- `WaveForecastResponse`: Detailed wave forecast data.
- `WindForecastResponse`: Wind-related forecast data.
- `SunlightForecastResponse`: Sunlight timings and details.
- `TideForecastResponse`: Tide-related forecast data.

... and many more for detailed data parsing.

## Contributing

Contributions are always welcome! Open an issue first, to start a discussion.

---

Happy coding and may you always find the perfect wave! 🌊🏄‍♂️