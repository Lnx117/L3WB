package service

import (
	"L3WB"
	"L3WB/pkg/repository"
)

type Geocoding interface {
	GetCitiesGeoList(cityList []L3WB.CityList) []L3WB.CityInfo
	GetCitiesTemperatureInfo(CityGeo []L3WB.CityGeo) []L3WB.CityTempInfo
	BackgroundUpdatingProcess()
	GetGeoAboutAllCities()
}

type Api interface {
	GetApiCityList() ([]string, error)
	GetShortCityInfo(cityName string) (L3WB.ShortCityInfoApiAnswer, error)
	GetFullCityInfo(cityName string, date string) (L3WB.AllCityInfoJson, error)
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
