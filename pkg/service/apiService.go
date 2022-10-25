package service

import (
	"L3WB"
	"L3WB/pkg/repository"
	"encoding/json"
	"sort"
	"time"

	"github.com/sirupsen/logrus"
)

type apiService struct {
	repo repository.PostgresQueries
}

func NewApiService(repo repository.PostgresQueries) *apiService {
	return &apiService{repo: repo}
}

func (a *apiService) GetCityNameAndIdListFromDb() ([]string, error) {
	cityNameAndIdList, err := a.repo.GetCityNameAndIdList()
	if err != nil {
		logrus.Error("Querying coordinate info for city list error (GetCityNameAndIdListFromDb func): %s", err.Error())
		return nil, err
	}

	var cityNameList []string
	for _, v := range cityNameAndIdList {
		cityNameList = append(cityNameList, v.Name)
	}

	sort.Slice(cityNameList, func(i, j int) bool {
		return cityNameList[i] < cityNameList[j]
	})

	return cityNameList, nil
}

func (a *apiService) GetShortCityWeatherDataByName(cityName string) (L3WB.ShortCityWeatherData, error) {
	var shortCityWeatherData L3WB.ShortCityWeatherData

	shortCityWeatherData.CityName = cityName

	curentTime := time.Now().Format("02-Jan-2006 15:04:05")
	cityId, err := a.repo.GetCityIdByName(cityName)

	if err != nil {
		logrus.Error("Getting city id by name error (GetShortCityWeatherDataByName func): %s", err.Error())
		return shortCityWeatherData, err
	}

	cityWeatherDataListForTimeAfterDate, err := a.repo.GetCityWeatherDataListForTimeAfterDate(cityId, curentTime)

	var avgTemp float64
	var dateList []string

	for _, v := range cityWeatherDataListForTimeAfterDate {
		avgTemp = avgTemp + v.Temp
		dateList = append(dateList, v.Date)
	}

	avgTemp = avgTemp / float64(len(cityWeatherDataListForTimeAfterDate))

	sort.Slice(dateList, func(i, j int) bool {
		return dateList[i] < dateList[j]
	})

	var fullCityGeoAndWeatherData L3WB.FullCityGeoAndWeatherData
	err = json.Unmarshal(cityWeatherDataListForTimeAfterDate[0].Full_info, &fullCityGeoAndWeatherData)
	if err != nil {
		logrus.Error("JSON decoding error (GetShortCityWeatherDataByName func): %s", err.Error())
	}

	shortCityWeatherData.AvgTemp = avgTemp
	shortCityWeatherData.Date = dateList
	shortCityWeatherData.Country = fullCityGeoAndWeatherData.CityGeoData.Country

	return shortCityWeatherData, nil
}

func (a *apiService) GetFullCityWeatherData(cityName string, date string) (L3WB.FullCityGeoAndWeatherData, error) {
	var fullCityGeoAndWeatherData L3WB.FullCityGeoAndWeatherData

	cityId, err := a.repo.GetCityIdByName(cityName)

	if err != nil {
		logrus.Error("Getting city id by name error (GetFullCityWeatherData func): %s", err.Error())
		return fullCityGeoAndWeatherData, err
	}

	cityWeatherDataByDate, err := a.repo.GetCityWeatherDataByDate(cityId, date)
	if err != nil {
		logrus.Error("Getting full city info error (GetFullCityWeatherData func): %s", err.Error())
		return fullCityGeoAndWeatherData, err
	}

	err = json.Unmarshal(cityWeatherDataByDate.Full_info, &fullCityGeoAndWeatherData)
	if err != nil {
		logrus.Error("JSON decoding error (GetFullCityWeatherData func): %s", err.Error())
		return fullCityGeoAndWeatherData, err
	}

	return fullCityGeoAndWeatherData, nil
}
