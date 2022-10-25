package service

import (
	"L3WB"
	"L3WB/pkg/repository"
)

type Geocoding interface {
	GetCitiesGeoData([]L3WB.CityNameAndId) []L3WB.CityGeoData
	GetCitiesOpenweathermapData([]L3WB.CityLatAndLon) []L3WB.CityWeatherDataForFiveDays
	BackgroundUpdatingProcess()
	GetGeoAboutAllCities()
}

type Api interface {
	GetCityNameAndIdListFromDb() ([]string, error)
	GetShortCityWeatherDataByName(string) (L3WB.ShortCityWeatherData, error)
	GetFullCityWeatherData(string, string) (L3WB.FullCityGeoAndWeatherData, error)
}

type Service struct {
	Geocoding
	Api
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Geocoding: NewGeocodingService(repos),
		Api:       NewApiService(repos),
	}
}
