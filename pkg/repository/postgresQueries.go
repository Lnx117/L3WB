package repository

import (
	"L3WB"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

type postgresQueries struct {
	db *sql.DB
}

func NewpostgresQueries(db *sql.DB) *postgresQueries {
	return &postgresQueries{db}
}

func (p *postgresQueries) GetCityNameAndIdList() ([]L3WB.CityNameAndId, error) {
	var cityNameAndIdList []L3WB.CityNameAndId

	query := fmt.Sprintf("SELECT id, name FROM %s", cityTable)
	rows, err := p.db.Query(query)

	if err != nil {
		logrus.Fatal("Getting city name and id list error (GetCityNameList func): %s", err.Error())
	}

	var cityNameAndId L3WB.CityNameAndId

	for rows.Next() {
		err := rows.Scan(&cityNameAndId.Id, &cityNameAndId.Name)
		if err != nil {
			logrus.Fatal("Getting city name and id list parsing error (GetCityNameList func): %s", err.Error())
		}

		cityNameAndIdList = append(cityNameAndIdList, cityNameAndId)
	}
	return cityNameAndIdList, err
}

func (p *postgresQueries) UpdateCitiesGeoData(cityGeoDataList []L3WB.CityGeoData) {

	for _, v := range cityGeoDataList {

		query := fmt.Sprintf("UPDATE %s SET lat=%f, lon=%f, country='%s' WHERE id = %d",
			cityTable, v.Lat, v.Lon, v.Country, v.Id)

		_, err := p.db.Exec(query)

		if err != nil {
			logrus.Error("DB update lat and lon error (UpdateCitiesGeoData func): %s", err.Error())
		}
	}
}

func (p *postgresQueries) GetCitiesLatAndLonList() ([]L3WB.CityLatAndLon, error) {

	var CityLatAndLonList []L3WB.CityLatAndLon

	query := fmt.Sprintf("SELECT id, lat, lon FROM %s", cityTable)
	rows, err := p.db.Query(query)

	if err != nil {
		logrus.Error("Getting city id, lat, lon list error (GetCitiesLatAndLonList func): %s", err.Error())
	}

	var cityLatAndLon L3WB.CityLatAndLon

	for rows.Next() {
		err := rows.Scan(&cityLatAndLon.Id, &cityLatAndLon.Lat, &cityLatAndLon.Lon)
		if err != nil {
			logrus.Error("Getting city id, lat, lon list parsing error (GetCitiesLatAndLonList func): %s", err.Error())
		}

		CityLatAndLonList = append(CityLatAndLonList, cityLatAndLon)
	}

	return CityLatAndLonList, err
}

func (p *postgresQueries) InsertOrUpdateCitiesWeatherData(citiesWeatherDataForFiveDaysList []L3WB.CityWeatherDataForFiveDays) {

	for _, cityWeatherDataForFiveDays := range citiesWeatherDataForFiveDaysList {
		for i, v := range cityWeatherDataForFiveDays.List {
			query := fmt.Sprintf("INSERT INTO %s (city_id, temp, date, Full_info) VALUES (%d, %f, '%s', $1) ON CONFLICT (temp, date) DO NOTHING",
				cityTemp, cityWeatherDataForFiveDays.CitiId, v.Main.Temp, v.DtTxt)

			var fullCityGeoAndWeatherData L3WB.FullCityGeoAndWeatherData

			fullCityGeoAndWeatherData.CityGeoData, fullCityGeoAndWeatherData.CityWeatherData = cityWeatherDataForFiveDays.City, cityWeatherDataForFiveDays.List[i]
			fullCityGeoAndWeatherDataJson, err := json.Marshal(fullCityGeoAndWeatherData)

			if err != nil {
				logrus.Error("Marshalling city temp data to json error (InsertOrUpdateCitiesWeatherData func): %s", err.Error())
			}

			_, err = p.db.Exec(query, fullCityGeoAndWeatherDataJson)

			if err != nil {
				logrus.Error("Update (or insert) temperature data error (InsertOrUpdateCitiesWeatherData func): %s", err.Error())
			}
		}
	}
}

func (p *postgresQueries) GetCityIdByName(cityName string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT ID FROM %s WHERE NAME = '%s'", cityTable, cityName)

	row := p.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		logrus.Error("Select city id by name error (GetCityIdByName func): %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (p *postgresQueries) GetCityWeatherDataListForTimeAfterDate(cityId int, date string) ([]L3WB.CityWeatherDataTable, error) {

	var cityWeatherDataTableList []L3WB.CityWeatherDataTable

	query := fmt.Sprintf("SELECT * FROM %s WHERE CITY_ID = %d and DATE > '%s'", cityTemp, cityId, date)
	rows, err := p.db.Query(query)

	if err != nil {
		logrus.Error("Select all cityTEmp rows by date and city name error (GetCityWeatherDataListForTimeAfterDate func): %s", err.Error())
	}

	var cityWeatherDataTable L3WB.CityWeatherDataTable

	for rows.Next() {
		err := rows.Scan(&cityWeatherDataTable.CityId, &cityWeatherDataTable.Temp, &cityWeatherDataTable.Date, &cityWeatherDataTable.Full_info)
		if err != nil {
			logrus.Error("Getting city id, lat, lon list parsing error (GetCityWeatherDataListForTimeAfterDate func): %s", err.Error())
		}

		cityWeatherDataTableList = append(cityWeatherDataTableList, cityWeatherDataTable)
	}

	return cityWeatherDataTableList, err
}

func (p *postgresQueries) GetCityWeatherDataByDate(cityId int, date string) (L3WB.CityWeatherDataTable, error) {

	var cityWeatherDataTable L3WB.CityWeatherDataTable
	query := fmt.Sprintf("SELECT * FROM %s WHERE DATE = '%s' AND CITY_ID = %d", cityTemp, date, cityId)

	row := p.db.QueryRow(query)
	if err := row.Scan(&cityWeatherDataTable.CityId, &cityWeatherDataTable.Temp, &cityWeatherDataTable.Date, &cityWeatherDataTable.Full_info); err != nil {
		logrus.Error("Select city id by name error: %s", err.Error())
		return cityWeatherDataTable, err
	}

	return cityWeatherDataTable, nil
}
