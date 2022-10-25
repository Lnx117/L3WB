package handler

import (
	"net/http"
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
	apiCityList, err := h.services.Api.GetApiCityList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, apiCityList)
}

// shortInfo godoc
// @Summary ShortCityInfo
// @Description Show short weather info
// @ID shortInfo
// @Param cityName path string true "City name"
// @Accept  json
// @Produce  json
// @Success 200 {object} L3WB.ShortCityInfoApiAnswer
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/shortInfo/{cityName} [get]
func (h *Handler) shortInfo(c *gin.Context) {
	cityName := c.Param("cityName")
	shortCityInfo, err := h.services.GetShortCityInfo(cityName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, shortCityInfo)
}

// fullInfo godoc
// @Summary FullInfo
// @Description Show full weather info
// @ID fullInfo
// @Param cityName path string true "City name"
// @Param date path string true "Date"
// @Accept  json
// @Produce  json
// @Success 200 {object} L3WB.AllCityInfo
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/full_info/{cityName}/{date} [get]
func (h *Handler) fullInfo(c *gin.Context) {
	cityName := c.Param("cityName")
	stringDate := c.Param("date")
	t, err := time.Parse(time.RFC3339, stringDate)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fullCityInfo, err := h.services.GetFullCityInfo(cityName, t.Format("2006-01-02 15:04:05"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fullCityInfo)
}
