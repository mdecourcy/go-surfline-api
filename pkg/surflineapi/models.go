package surflineapi

import "time"

// General types
type LatLng struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Units struct {
	Temperature string `json:"temperature"`
	TideHeight  string `json:"tideHeight"`
	SwellHeight string `json:"swellHeight"`
	WaveHeight  string `json:"waveHeight"`
	WindSpeed   string `json:"windSpeed"`
	Pressure    string `json:"pressure"`
}

type Associated struct {
	Units                      Units  `json:"units"`
	UtcOffset                  int    `json:"utcOffset"`
	Location                   LatLng `json:"location"`
	ForecastLocation           LatLng `json:"forecastLocation"`
	OffshoreLocation           LatLng `json:"offshoreLocation"`
	RunInitializationTimestamp int64  `json:"runInitializationTimestamp"`
}

// Region Forecast Conditions
type RegionsForecastConditionsResponse struct {
	Associated  Associated     `json:"associated"`
	Data        ForecastData   `json:"data"`
	Permissions UserPermission `json:"permissions"`
}

type ForecastData struct {
	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	Timestamp   int64      `json:"timestamp"`
	ForecastDay string     `json:"forecastDay"`
	Forecaster  Forecaster `json:"forecaster"`
	Human       bool       `json:"human"`
	Observation string     `json:"observation"`
	AM          TimeData   `json:"am"`
	PM          TimeData   `json:"pm"`
	UtcOffset   int        `json:"utcOffset"`
}

type Forecaster struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type TimeData struct {
	Timestamp        int64  `json:"timestamp"`
	Observation      string `json:"observation"`
	Rating           string `json:"rating"`
	MinHeight        int    `json:"minHeight"`
	MaxHeight        int    `json:"maxHeight"`
	Plus             bool   `json:"plus"`
	HumanRelation    string `json:"humanRelation"`
	OccasionalHeight *int   `json:"occasionalHeight"`
}

// Spot Forecast Rating
type SpotForecastRatingResponse struct {
	Associated Associated `json:"associated"`
	Data       struct {
		Rating []struct {
			Timestamp int64 `json:"timestamp"`
			UtcOffset int   `json:"utcOffset"`
			Rating    struct {
				Key   string  `json:"key"`
				Value float32 `json:"value"`
			} `json:"rating"`
		} `json:"rating"`
	} `json:"data"`
}

// User Permission and Violation
type UserPermission struct {
	Data       []interface{} `json:"data"`
	Violations []Violation   `json:"violations"`
}

type Violation struct {
	Code       int        `json:"code"`
	Message    string     `json:"message"`
	Permission Permission `json:"permission"`
}

type Permission struct {
	Name     string `json:"name"`
	Required bool   `json:"required"`
}

// Wave Forecast
type WaveForecastResponse struct {
	Associated Associated `json:"associated"`
	Data       struct {
		Wave []WaveData `json:"wave"`
	} `json:"data"`
	Permissions UserPermission `json:"permissions"`
}

type WaveData struct {
	Timestamp   int64   `json:"timestamp"`
	Probability float64 `json:"probability"`
	UtcOffset   float64 `json:"utcOffset"`
	Surf        struct {
		Min           float64 `json:"min"`
		Max           float64 `json:"max"`
		OptimalScore  int     `json:"optimalScore"`
		Plus          bool    `json:"plus"`
		HumanRelation string  `json:"humanRelation"`
		Raw           struct {
			Min float64 `json:"min"`
			Max float64 `json:"max"`
		} `json:"raw"`
	} `json:"surf"`
	Power  float64 `json:"power"`
	Swells []struct {
		Height       float64 `json:"height"`
		Period       int     `json:"period"`
		Impact       float64 `json:"impact"`
		Power        float64 `json:"power"`
		Direction    float64 `json:"direction"`
		DirectionMin float64 `json:"directionMin"`
		OptimalScore int     `json:"optimalScore"`
	} `json:"swells"`
}

// Wind Forecast
type WindForecastResponse struct {
	Associated  WindAssociated `json:"associated"`
	Data        WindData       `json:"data"`
	Permissions UserPermission `json:"permissions"`
}

type WindAssociated struct {
	Units                      Units  `json:"units"`
	UtcOffset                  int    `json:"utcOffset"`
	Location                   LatLng `json:"location"`
	RunInitializationTimestamp int64  `json:"runInitializationTimestamp"`
}

type WindData struct {
	Wind []WindDetail `json:"wind"`
}

type WindDetail struct {
	Timestamp     int64   `json:"timestamp"`
	UtcOffset     int     `json:"utcOffset"`
	Speed         float64 `json:"speed"`
	Direction     float64 `json:"direction"`
	DirectionType string  `json:"directionType"`
	Gust          float64 `json:"gust"`
	OptimalScore  int     `json:"optimalScore"`
}

// Sunlight Forecast
type SunlightForecastResponse struct {
	Associated  Associated     `json:"associated"`
	Data        SunlightData   `json:"data"`
	Permissions UserPermission `json:"permissions"`
}

type SunlightData struct {
	Sunlight []SunlightInfo `json:"sunlight"`
}

type SunlightInfo struct {
	Midnight          int64 `json:"midnight"`
	MidnightUTCOffset int   `json:"midnightUTCOffset"`
	Dawn              int64 `json:"dawn"`
	DawnUTCOffset     int   `json:"dawnUTCOffset"`
	Sunrise           int64 `json:"sunrise"`
	SunriseUTCOffset  int   `json:"sunriseUTCOffset"`
	Sunset            int64 `json:"sunset"`
	SunsetUTCOffset   int   `json:"sunsetUTCOffset"`
	Dusk              int64 `json:"dusk"`
	DuskUTCOffset     int   `json:"duskUTCOffset"`
}

// Tides Forecast

type TideForecastResponse struct {
	Associated  TideAssociated `json:"associated"`
	Data        TideData       `json:"data"`
	Permissions UserPermission `json:"permissions"`
}

type TideAssociated struct {
	UtcOffset    int          `json:"utcOffset"`
	Units        Units        `json:"units"`
	TideLocation TideLocation `json:"tideLocation"`
}

type TideLocation struct {
	Name string  `json:"name"`
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
	Mean float64 `json:"mean"`
}

type TideData struct {
	Tides []TideInfo `json:"tides"`
}

type TideInfo struct {
	Timestamp int64   `json:"timestamp"`
	UtcOffset int     `json:"utcOffset"`
	Type      string  `json:"type"`
	Height    float64 `json:"height"`
}

// Weather Forecast

type WeatherForecastResponse struct {
	Associated struct {
		Units struct {
			Temperature string `json:"temperature"`
		} `json:"units"`
		UTCOffset                  int    `json:"utcOffset"`
		WeatherIconPath            string `json:"weatherIconPath"`
		RunInitializationTimestamp int    `json:"runInitializationTimestamp"`
	} `json:"associated"`
	Data struct {
		SunlightTimes []struct {
			Midnight          int `json:"midnight"`
			MidnightUTCOffset int `json:"midnightUTCOffset"`
			Dawn              int `json:"dawn"`
			DawnUTCOffset     int `json:"dawnUTCOffset"`
			Sunrise           int `json:"sunrise"`
			SunriseUTCOffset  int `json:"sunriseUTCOffset"`
			Sunset            int `json:"sunset"`
			SunsetUTCOffset   int `json:"sunsetUTCOffset"`
			Dusk              int `json:"dusk"`
			DuskUTCOffset     int `json:"duskUTCOffset"`
		} `json:"sunlightTimes"`
		Weather []struct {
			Timestamp   int     `json:"timestamp"`
			UTCOffset   int     `json:"utcOffset"`
			Temperature float64 `json:"temperature"`
			Condition   string  `json:"condition"`
			Pressure    int     `json:"pressure"`
		} `json:"weather"`
	} `json:"data"`
	Permissions struct {
		Data       []interface{} `json:"data"`
		Violations []struct {
			Code       int    `json:"code"`
			Message    string `json:"message"`
			Permission struct {
				Name     string `json:"name"`
				Required bool   `json:"required"`
			} `json:"permission"`
		} `json:"violations"`
	} `json:"permissions"`
}

// Nearby Buoys
type NearbyBuoysResponse struct {
	Associated struct {
		Units struct {
			SwellHeight string `json:"swellHeight"`
		} `json:"units"`
	} `json:"associated"`
	Data []struct {
		ID           string  `json:"id"`
		Name         string  `json:"name"`
		SourceID     string  `json:"sourceId"`
		Latitude     float64 `json:"latitude"`
		Longitude    float64 `json:"longitude"`
		Status       string  `json:"status"`
		AbbrTimezone string  `json:"abbrTimezone"`
		LatestData   struct {
			Timestamp int     `json:"timestamp"`
			Height    float64 `json:"height"`
			Period    int     `json:"period"`
			Direction *int    `json:"direction,omitempty"` // *int since direction can be null
			Swells    []struct {
				Height    float64 `json:"height"`
				Period    int     `json:"period"`
				Direction int     `json:"direction"`
			} `json:"swells"`
			UTCOffset int `json:"utcOffset"`
		} `json:"latestData"`
	} `json:"data"`
}

type TaxonomyDetailsResponse struct {
	ID        string   `json:"_id"`
	GeonameID int      `json:"geonameId"`
	Type      string   `json:"type"`
	LiesIn    []string `json:"liesIn"`
	Geonames  struct {
		Population  int    `json:"population"`
		Fcode       string `json:"fcode"`
		Fcl         string `json:"fcl"`
		Lat         string `json:"lat"`
		AdminName1  string `json:"adminName1"`
		FcodeName   string `json:"fcodeName"`
		ToponymName string `json:"toponymName"`
		FclName     string `json:"fclName"`
		Name        string `json:"name"`
		GeonameID   int    `json:"geonameId"`
		Lng         string `json:"lng"`
	} `json:"geonames"`
	Location struct {
		Coordinates []float64 `json:"coordinates"`
		Type        string    `json:"type"`
	} `json:"location"`
	EnumeratedPath string `json:"enumeratedPath"`
	Name           string `json:"name"`
	Category       string `json:"category"`
	HasSpots       bool   `json:"hasSpots"`
	Associated     struct {
		Links []any `json:"links"`
	} `json:"associated"`
	In []struct {
		ID        string `json:"_id"`
		GeonameID int    `json:"geonameId"`
		Type      string `json:"type"`
		Geonames  struct {
			Population  int64  `json:"population"`
			Fcode       string `json:"fcode"`
			Fcl         string `json:"fcl"`
			Lat         string `json:"lat"`
			AdminName1  string `json:"adminName1"`
			FcodeName   string `json:"fcodeName"`
			ToponymName string `json:"toponymName"`
			FclName     string `json:"fclName"`
			Name        string `json:"name"`
			GeonameID   int    `json:"geonameId"`
			Lng         string `json:"lng"`
		} `json:"geonames"`
		Location struct {
			Coordinates []int  `json:"coordinates"`
			Type        string `json:"type"`
		} `json:"location"`
		EnumeratedPath string `json:"enumeratedPath"`
		Name           string `json:"name"`
		Category       string `json:"category"`
		HasSpots       bool   `json:"hasSpots"`
		LiesIn         []any  `json:"liesIn"`
		Depth          int    `json:"depth"`
		Associated     struct {
			Links []any `json:"links"`
		} `json:"associated"`
	} `json:"in"`
	Contains []struct {
		ID        string   `json:"_id"`
		GeonameID int      `json:"geonameId"`
		Type      string   `json:"type"`
		LiesIn    []string `json:"liesIn"`
		Geonames  struct {
			Fcode       string `json:"fcode"`
			Lat         string `json:"lat"`
			AdminName1  string `json:"adminName1"`
			FcodeName   string `json:"fcodeName"`
			CountryName string `json:"countryName"`
			FclName     string `json:"fclName"`
			Name        string `json:"name"`
			CountryCode string `json:"countryCode"`
			Population  int    `json:"population"`
			Fcl         string `json:"fcl"`
			CountryID   string `json:"countryId"`
			ToponymName string `json:"toponymName"`
			GeonameID   int    `json:"geonameId"`
			Lng         string `json:"lng"`
			AdminCode1  string `json:"adminCode1"`
		} `json:"geonames"`
		Location struct {
			Coordinates []float64 `json:"coordinates"`
			Type        string    `json:"type"`
		} `json:"location"`
		EnumeratedPath string `json:"enumeratedPath"`
		Name           string `json:"name"`
		Category       string `json:"category"`
		HasSpots       bool   `json:"hasSpots"`
		Depth          int    `json:"depth"`
		Associated     struct {
			Links []struct {
				Key  string `json:"key"`
				Href string `json:"href"`
			} `json:"links"`
		} `json:"associated"`
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
		CreatedAt time.Time `json:"createdAt,omitempty"`
	} `json:"contains"`
}
