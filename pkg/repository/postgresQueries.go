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

func (p *postgresQueries) GetCityNameList() ([]L3WB.CityList, error) {
	var list []L3WB.CityList

	query := fmt.Sprintf("SELECT id, name FROM %s", cityTable)
	rows, err := p.db.Query(query)

	if err != nil {
		logrus.Error("Getting city name and id list error: %s", err.Error())
	}

	var cityItem L3WB.CityList

	for rows.Next() {
		err := rows.Scan(&cityItem.Id, &cityItem.Name)
		if err != nil {
			logrus.Error("Getting city name and id list parsing error: %s", err.Error())
		}

		list = append(list, cityItem)
	}
	return list, err
}

func (p *postgresQueries) UpdateCitiesGeo(cityInfoList []L3WB.CityInfo) {

	for _, v := range cityInfoList {

		query := fmt.Sprintf("UPDATE %s SET lat=%f, lon=%f, state='%s', country='%s' WHERE id = %d",
			cityTable, v.Lat, v.Lon, v.State, v.Country, v.Id)

		_, err := p.db.Exec(query)

		if err != nil {
			logrus.Error("Update lat and lon error: %s", err.Error())
		}
	}
}

func (p *postgresQueries) GetCitiesGeoList() ([]L3WB.CityGeo, error) {

	var list []L3WB.CityGeo

	query := fmt.Sprintf("SELECT id, lat, lon FROM %s", cityTable)
	rows, err := p.db.Query(query)

	if err != nil {
		logrus.Error("Getting city id, lat, lon list error: %s", err.Error())
	}

	var cityItem L3WB.CityGeo

	for rows.Next() {
		err := rows.Scan(&cityItem.Id, &cityItem.Lat, &cityItem.Lon)
		if err != nil {
			logrus.Error("Getting city id, lat, lon list parsing error: %s", err.Error())
		}

		list = append(list, cityItem)
	}

	return list, err
}

func (p *postgresQueries) InsertOrUpdateCitiesTemperatureInfo(CityTempInfo []L3WB.CityTempInfo) {

	for _, cityInfo := range CityTempInfo {
		for i, v := range cityInfo.List {
			query := fmt.Sprintf("INSERT INTO %s (city_id, temp, date, Full_info) VALUES (%d, %f, '%s', $1) ON CONFLICT (temp, date) DO NOTHING",
				cityTemp, cityInfo.CitiId, v.Main.Temp, v.DtTxt)

			var jsonInfo L3WB.AllCityInfoJson
			jsonInfo.City, jsonInfo.CityTempInfo = cityInfo.City, cityInfo.List[i]
			jsonData, err := json.Marshal(jsonInfo)

			if err != nil {
				logrus.Error("Marshalling city temp data to json error: %s", err.Error())
			}

			_, err = p.db.Exec(query, jsonData)

			if err != nil {
				logrus.Error("Update (or insert) temperature data error: %s", err.Error())
			}
		}
	}
}

func (p *postgresQueries) ReturnCityIdByName(cityName string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT ID FROM %s WHERE NAME = '%s'", cityTable, cityName)

	row := p.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		logrus.Error("Select city id by name error: %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (p *postgresQueries) GetAllCityTempRowsByCityNameAfterDate(cityId int, date string) ([]L3WB.CityTemp, error) {

	var list []L3WB.CityTemp

	query := fmt.Sprintf("SELECT * FROM %s WHERE CITY_ID = %d and DATE > '%s'", cityTemp, cityId, date)
	rows, err := p.db.Query(query)

	if err != nil {
		logrus.Error("Select all cityTEmp rows by date and city name error: %s", err.Error())
	}

	var cityTempItem L3WB.CityTemp

	for rows.Next() {
		err := rows.Scan(&cityTempItem.CityId, &cityTempItem.Temp, &cityTempItem.Date, &cityTempItem.Full_info)
		if err != nil {
			logrus.Error("Getting city id, lat, lon list parsing error: %s", err.Error())
		}

		list = append(list, cityTempItem)
	}

	fmt.Println(list)
	return list, err
}

func (p *postgresQueries) GetCityTempRowByDate(cityId int, date string) (L3WB.CityTemp, error) {

	var fullCityInfo L3WB.CityTemp
	query := fmt.Sprintf("SELECT * FROM %s WHERE DATE = '%s' AND CITY_ID = %d", cityTemp, date, cityId)

	row := p.db.QueryRow(query)
	if err := row.Scan(&fullCityInfo.CityId, &fullCityInfo.Temp, &fullCityInfo.Date, &fullCityInfo.Full_info); err != nil {
		logrus.Error("Select city id by name error: %s", err.Error())
		return fullCityInfo, err
	}

	return fullCityInfo, nil
}
