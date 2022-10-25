package repository

import (
	"L3WB"
	"database/sql"
)

type PostgresQueries interface {
	GetCityNameAndIdList() ([]L3WB.CityNameAndId, error)
	UpdateCitiesGeoData([]L3WB.CityGeoData)
	GetCitiesLatAndLonList() ([]L3WB.CityLatAndLon, error)
	InsertOrUpdateCitiesWeatherData([]L3WB.CityWeatherDataForFiveDays)
	GetCityIdByName(string) (int, error)
	GetCityWeatherDataListForTimeAfterDate(int, string) ([]L3WB.CityWeatherDataTable, error)
	GetCityWeatherDataByDate(int, string) (L3WB.CityWeatherDataTable, error)
}

type Repository struct {
	PostgresQueries
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		PostgresQueries: NewpostgresQueries(db),
	}
}
