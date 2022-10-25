package repository

import (
	"L3WB"
	"database/sql"
)

type PostgresQueries interface {
	GetCityNameList() ([]L3WB.CityList, error)
	UpdateCitiesGeo(cityInfoList []L3WB.CityInfo)
	GetCitiesGeoList() ([]L3WB.CityGeo, error)
	InsertOrUpdateCitiesTemperatureInfo(CityTempMain []L3WB.CityTempInfo)
	ReturnCityIdByName(cityName string) (int, error)
	GetAllCityTempRowsByCityNameAfterDate(cityId int, date string) ([]L3WB.CityTemp, error)
	GetCityTempRowByDate(cityId int, date string) (L3WB.CityTemp, error)
}

type Repository struct {
	PostgresQueries
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		PostgresQueries: NewpostgresQueries(db),
	}
}
