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

type syncCitiesOpenweathermapData struct {
	mu   sync.Mutex
	list []L3WB.CityWeatherDataForFiveDays
}

func NewGeocodingService(repo repository.PostgresQueries) *geocodingService {
	return &geocodingService{repo: repo}
}

func (g *geocodingService) GetCitiesGeoData(cityNameAndIdList []L3WB.CityNameAndId) []L3WB.CityGeoData {

	var cityGeoDataList []L3WB.CityGeoData

	for _, cityNameAndId := range cityNameAndIdList {
		query := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", cityNameAndId.Name, os.Getenv("OPEN_WEATHER_API_KEY"))

		req, err := http.Get(query)
		if err != nil {
			logrus.Error("Querying coordinate info for city list error (GetCitiesGeoData func): %s", err.Error())
		}
		defer req.Body.Close()
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			logrus.Error("Getting coordinate info for city list error (GetCitiesGeoData func): %s", err.Error())
		}

		var cityGeoData []L3WB.CityGeoData
		err = json.Unmarshal([]byte(data), &cityGeoData)
		if err != nil {
			logrus.Error("JSON decoding error (GetCitiesGeoData func): %s", err.Error())
		}

		if len(cityGeoData) == 0 {
			logrus.Error("Answer is empty (happend when city name are wrong (GetCitiesGeoData func))")
			continue
		}

		//Add city id to struct
		cityGeoData[0].Id = cityNameAndId.Id

		cityGeoDataList = append(cityGeoDataList, cityGeoData[0])
	}

	return cityGeoDataList
}

func (g *geocodingService) GetCitiesOpenweathermapData(cityGeoDataList []L3WB.CityLatAndLon) []L3WB.CityWeatherDataForFiveDays {

	var syncCitiesOpenweathermapData syncCitiesOpenweathermapData

	var wg sync.WaitGroup
	wg.Add(len(cityGeoDataList))

	for _, v := range cityGeoDataList {
		go syncCitiesOpenweathermapData.GetCityOpenweathermapData(v, &wg)
	}

	wg.Wait()
	return syncCitiesOpenweathermapData.list
}

func (s *syncCitiesOpenweathermapData) GetCityOpenweathermapData(cityLatAndLon L3WB.CityLatAndLon, wg *sync.WaitGroup) {
	query := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s", cityLatAndLon.Lat, cityLatAndLon.Lon, os.Getenv("OPEN_WEATHER_API_KEY"))

	req, err := http.Get(query)
	if err != nil {
		logrus.Error("Getting city temperature info error: %s", err.Error())
	}
	defer req.Body.Close()
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logrus.Error("Getting city temperature info error (ioutil.ReadAll(req.Body)): %s", err.Error())
	}

	var cityWeatherDataForFiveDays L3WB.CityWeatherDataForFiveDays
	err = json.Unmarshal([]byte(data), &cityWeatherDataForFiveDays)
	if err != nil {
		logrus.Error("JSON decoding error (Getting city temperature info): %s", err.Error())
	}

	cityWeatherDataForFiveDays.CitiId = cityLatAndLon.Id

	s.mu.Lock()
	defer s.mu.Unlock()
	s.list = append(s.list, cityWeatherDataForFiveDays)
	wg.Done()
}

func (g *geocodingService) BackgroundUpdatingProcess() {
	for {
		/* Getting cities geo from db*/
		CitiesGeoList, _ := g.repo.GetCitiesLatAndLonList()

		/* Getting temperature info for every city from geoApi. Parallel execution!!!*/
		CitiesTemperatureInfo := g.GetCitiesOpenweathermapData(CitiesGeoList)

		/* Save temperature for every city and time to db*/
		g.repo.InsertOrUpdateCitiesWeatherData(CitiesTemperatureInfo)

		time.Sleep(60 * time.Second)
	}
}

/* Collect geo data from openweathermap for every city (coordinates, state, country etc.) */
func (g *geocodingService) GetGeoAboutAllCities() {

	/* Getting city's names from db */
	CityNameList, err := g.repo.GetCityNameAndIdList()
	if err != nil {
		logrus.Error("JSON decoding error (Getting city temperature info): %s", err.Error())
	}

	/* Getting Coordinates for every city from geoApi*/
	CitiesGeoList := g.GetCitiesGeoData(CityNameList)

	/* Save Coordinates for every city to db*/
	g.repo.UpdateCitiesGeoData(CitiesGeoList)
}
