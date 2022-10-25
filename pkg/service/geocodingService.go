package service

import (
	"L3WB"
	"L3WB/pkg/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type geocodingService struct {
	repo repository.PostgresQueries
}

type syncCytiesTempList struct {
	mu   sync.Mutex
	list []L3WB.CityTempInfo
}

func NewGeocodingService(repo repository.PostgresQueries) *geocodingService {
	return &geocodingService{repo: repo}
}

func (g *geocodingService) GetCitiesGeoList(cityList []L3WB.CityList) []L3WB.CityInfo {

	var list []L3WB.CityInfo

	for _, v := range cityList {
		query := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", v.Name, os.Getenv("OPEN_WEATHER_API_KEY"))

		req, err := http.Get(query)
		if err != nil {
			logrus.Error("Querying coordinate info for city list error: %s", err.Error())
		}
		defer req.Body.Close()
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			logrus.Error("Getting coordinate info for city list error: %s", err.Error())
		}

		var cityInfo []L3WB.CityInfo
		err = json.Unmarshal([]byte(data), &cityInfo)
		if err != nil {
			logrus.Error("JSON decoding error (getting coordinates for cityList): %s", err.Error())
		}

		if len(cityInfo) == 0 {
			logrus.Error("Answer is empty (happend when city name are wrong)")
			continue
		}

		//Add city id to struct
		cityInfo[0].Id = v.Id

		list = append(list, cityInfo[0])
	}

	return list
}

func (g *geocodingService) GetCitiesTemperatureInfo(CityGeo []L3WB.CityGeo) []L3WB.CityTempInfo {

	var syncCytiesTempList syncCytiesTempList

	var wg sync.WaitGroup
	wg.Add(len(CityGeo))

	for _, v := range CityGeo {
		go syncCytiesTempList.GetOneCityTemperatureInfo(v, &wg)
	}

	wg.Wait()
	return syncCytiesTempList.list
}

func (s *syncCytiesTempList) GetOneCityTemperatureInfo(v L3WB.CityGeo, wg *sync.WaitGroup) {
	query := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s", v.Lat, v.Lon, os.Getenv("OPEN_WEATHER_API_KEY"))

	req, err := http.Get(query)
	if err != nil {
		logrus.Error("Getting city temperature info error: %s", err.Error())
	}
	defer req.Body.Close()
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logrus.Error("Getting city temperature info error (ioutil.ReadAll(req.Body)): %s", err.Error())
	}

	var CityTempInfo L3WB.CityTempInfo
	err = json.Unmarshal([]byte(data), &CityTempInfo)
	if err != nil {
		logrus.Error("JSON decoding error (Getting city temperature info): %s", err.Error())
	}

	CityTempInfo.CitiId = v.Id

	s.mu.Lock()
	defer s.mu.Unlock()
	s.list = append(s.list, CityTempInfo)
	wg.Done()
}

func (g *geocodingService) BackgroundUpdatingProcess() {
	for {
		/* Getting cities geo from db*/
		CitiesGeoList, _ := g.repo.GetCitiesGeoList()

		/* Getting temperature info for every city from geoApi. Parallel execution!!!*/
		CitiesTemperatureInfo := g.GetCitiesTemperatureInfo(CitiesGeoList)

		/* Save temperature for every city and time to db*/
		g.repo.InsertOrUpdateCitiesTemperatureInfo(CitiesTemperatureInfo)

		time.Sleep(60 * time.Second)
	}
}

/* Collect geo data from openweathermap for every city (coordinates, state, country etc.) */
func (g *geocodingService) GetGeoAboutAllCities() {

	/* Getting city's names from db */
	CityNameList, err := g.repo.GetCityNameList()
	if err != nil {
		logrus.Error("JSON decoding error (Getting city temperature info): %s", err.Error())
	}

	/* Getting Coordinates for every city from geoApi*/
	CitiesGeoList := g.GetCitiesGeoList(CityNameList)

	/* Save Coordinates for every city to db*/
	g.repo.UpdateCitiesGeo(CitiesGeoList)
}
