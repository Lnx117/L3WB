package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// apiGetCityList godoc
// @Summary GetCityList
// @Description Get all cities avaliable for forecast
// @ID apiGetCityList
// @Accept  json
// @Produce  json
// @Success 200 {array}	string
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/cityList [get]
func (h *Handler) apiGetCityList(c *gin.Context) {
	cityList, err := h.services.Api.GetCityNameAndIdListFromDb()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cityList)
}

// shortInfo godoc
// @Summary ShortCityInfo
// @Description Show short weather info
// @ID shortInfo
// @Param cityName path string true "City name in any case"
// @Accept  json
// @Produce  json
// @Success 200 {object} L3WB.ShortCityWeatherData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/shortInfo/{cityName} [get]
func (h *Handler) shortInfo(c *gin.Context) {
	cityName := c.Param("cityName")
	shortCityWeatherData, err := h.services.GetShortCityWeatherDataByName(strings.ToLower(cityName))
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			newErrorResponse(c, http.StatusInternalServerError, "Make shure you use right city name. Just copy it from /cityList answer! /n"+err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, shortCityWeatherData)
}

// fullInfo godoc
// @Summary FullInfo
// @Description Show full weather info
// @ID fullInfo
// @Param cityName path string true "City name in any case"
// @Param date path string true "Date in 2022-10-26T12:00:00Z format"
// @Accept  json
// @Produce  json
// @Success 200 {object} L3WB.FullCityGeoAndWeatherData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/full_info/{cityName}/{date} [get]
func (h *Handler) fullInfo(c *gin.Context) {
	cityName := c.Param("cityName")
	stringDate := c.Param("date")
	t, err := time.Parse(time.RFC3339, stringDate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Make shure you use 2022-10-26T12:00:00Z date format. Just copy it from /shortInfo answer!")
		return
	}

	fullCityWeatherData, err := h.services.GetFullCityWeatherData(strings.ToLower(cityName), t.Format("2006-01-02 15:04:05"))
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			newErrorResponse(c, http.StatusInternalServerError, "Make shure you use 2022-10-26T12:00:00Z date format and right city name. Just copy it from /shortInfo and /cityList answers! /n"+err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fullCityWeatherData)
}
