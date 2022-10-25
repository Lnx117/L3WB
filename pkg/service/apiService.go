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

func (a *apiService) GetApiCityList() ([]string, error) {
	cityNameList, err := a.repo.GetCityNameList()
	if err != nil {
		logrus.Error("Querying coordinate info for city list error: %s", err.Error())
		return nil, err
	}

	var list []string
	for _, v := range cityNameList {
		list = append(list, v.Name)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	return list, nil
}

func (a *apiService) GetShortCityInfo(cityName string) (L3WB.ShortCityInfoApiAnswer, error) {
	var shortCityInfo L3WB.ShortCityInfoApiAnswer
	shortCityInfo.CityName = cityName
	curentTime := time.Now().Format("02-Jan-2006 15:04:05")
	cityId, err := a.repo.ReturnCityIdByName(cityName)

	if err != nil {
		logrus.Error("Getting city id by name error: %s", err.Error())
		return shortCityInfo, err
	}

	list, err := a.repo.GetAllCityTempRowsByCityNameAfterDate(cityId, curentTime)

	var avgTemp float64
	var dateList []string

	for _, v := range list {
		avgTemp = avgTemp + v.Temp
		dateList = append(dateList, v.Date)
	}

	avgTemp = avgTemp / float64(len(list))

	sort.Slice(dateList, func(i, j int) bool {
		return dateList[i] < dateList[j]
	})

	var fullInfo L3WB.AllCityInfoJson
	err = json.Unmarshal(list[0].Full_info, &fullInfo)
	if err != nil {
		logrus.Error("JSON decoding error (getting full_info in GetAllCityTempRowsByCityNameAfterDate): %s", err.Error())
	}

	shortCityInfo.AvgTemp = avgTemp
	shortCityInfo.Date = dateList
	shortCityInfo.Country = fullInfo.City.Country

	return shortCityInfo, nil
}

func (a *apiService) GetFullCityInfo(cityName string, date string) (L3WB.AllCityInfoJson, error) {
	var allCityInfo L3WB.AllCityInfoJson

	cityId, err := a.repo.ReturnCityIdByName(cityName)

	if err != nil {
		logrus.Error("Getting city id by name error (in getting full temp info route): %s", err.Error())
		return allCityInfo, err
	}

	cutyTempRow, err := a.repo.GetCityTempRowByDate(cityId, date)
	if err != nil {
		logrus.Error("Getting full city info error (in getting full temp info route): %s", err.Error())
		return allCityInfo, err
	}

	err = json.Unmarshal(cutyTempRow.Full_info, &allCityInfo)
	if err != nil {
		logrus.Error("JSON decoding error (getting full_info in GetFullCityInfo): %s", err.Error())
		return allCityInfo, err
	}

	return allCityInfo, nil
}
