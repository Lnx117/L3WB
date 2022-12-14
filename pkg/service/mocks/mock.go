// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	L3WB "L3WB"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGeocoding is a mock of Geocoding interface.
type MockGeocoding struct {
	ctrl     *gomock.Controller
	recorder *MockGeocodingMockRecorder
}

// MockGeocodingMockRecorder is the mock recorder for MockGeocoding.
type MockGeocodingMockRecorder struct {
	mock *MockGeocoding
}

// NewMockGeocoding creates a new mock instance.
func NewMockGeocoding(ctrl *gomock.Controller) *MockGeocoding {
	mock := &MockGeocoding{ctrl: ctrl}
	mock.recorder = &MockGeocodingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeocoding) EXPECT() *MockGeocodingMockRecorder {
	return m.recorder
}

// BackgroundUpdatingProcess mocks base method.
func (m *MockGeocoding) BackgroundUpdatingProcess() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BackgroundUpdatingProcess")
}

// BackgroundUpdatingProcess indicates an expected call of BackgroundUpdatingProcess.
func (mr *MockGeocodingMockRecorder) BackgroundUpdatingProcess() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackgroundUpdatingProcess", reflect.TypeOf((*MockGeocoding)(nil).BackgroundUpdatingProcess))
}

// GetCitiesGeoData mocks base method.
func (m *MockGeocoding) GetCitiesGeoData(arg0 []L3WB.CityNameAndId) []L3WB.CityGeoData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCitiesGeoData", arg0)
	ret0, _ := ret[0].([]L3WB.CityGeoData)
	return ret0
}

// GetCitiesGeoData indicates an expected call of GetCitiesGeoData.
func (mr *MockGeocodingMockRecorder) GetCitiesGeoData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCitiesGeoData", reflect.TypeOf((*MockGeocoding)(nil).GetCitiesGeoData), arg0)
}

// GetCitiesOpenweathermapData mocks base method.
func (m *MockGeocoding) GetCitiesOpenweathermapData(arg0 []L3WB.CityLatAndLon) []L3WB.CityWeatherDataForFiveDays {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCitiesOpenweathermapData", arg0)
	ret0, _ := ret[0].([]L3WB.CityWeatherDataForFiveDays)
	return ret0
}

// GetCitiesOpenweathermapData indicates an expected call of GetCitiesOpenweathermapData.
func (mr *MockGeocodingMockRecorder) GetCitiesOpenweathermapData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCitiesOpenweathermapData", reflect.TypeOf((*MockGeocoding)(nil).GetCitiesOpenweathermapData), arg0)
}

// GetGeoAboutAllCities mocks base method.
func (m *MockGeocoding) GetGeoAboutAllCities() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetGeoAboutAllCities")
}

// GetGeoAboutAllCities indicates an expected call of GetGeoAboutAllCities.
func (mr *MockGeocodingMockRecorder) GetGeoAboutAllCities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeoAboutAllCities", reflect.TypeOf((*MockGeocoding)(nil).GetGeoAboutAllCities))
}

// MockApi is a mock of Api interface.
type MockApi struct {
	ctrl     *gomock.Controller
	recorder *MockApiMockRecorder
}

// MockApiMockRecorder is the mock recorder for MockApi.
type MockApiMockRecorder struct {
	mock *MockApi
}

// NewMockApi creates a new mock instance.
func NewMockApi(ctrl *gomock.Controller) *MockApi {
	mock := &MockApi{ctrl: ctrl}
	mock.recorder = &MockApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApi) EXPECT() *MockApiMockRecorder {
	return m.recorder
}

// GetCityNameAndIdListFromDb mocks base method.
func (m *MockApi) GetCityNameAndIdListFromDb() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCityNameAndIdListFromDb")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCityNameAndIdListFromDb indicates an expected call of GetCityNameAndIdListFromDb.
func (mr *MockApiMockRecorder) GetCityNameAndIdListFromDb() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCityNameAndIdListFromDb", reflect.TypeOf((*MockApi)(nil).GetCityNameAndIdListFromDb))
}

// GetFullCityWeatherData mocks base method.
func (m *MockApi) GetFullCityWeatherData(arg0, arg1 string) (L3WB.FullCityGeoAndWeatherData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullCityWeatherData", arg0, arg1)
	ret0, _ := ret[0].(L3WB.FullCityGeoAndWeatherData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullCityWeatherData indicates an expected call of GetFullCityWeatherData.
func (mr *MockApiMockRecorder) GetFullCityWeatherData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullCityWeatherData", reflect.TypeOf((*MockApi)(nil).GetFullCityWeatherData), arg0, arg1)
}

// GetShortCityWeatherDataByName mocks base method.
func (m *MockApi) GetShortCityWeatherDataByName(arg0 string) (L3WB.ShortCityWeatherData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShortCityWeatherDataByName", arg0)
	ret0, _ := ret[0].(L3WB.ShortCityWeatherData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShortCityWeatherDataByName indicates an expected call of GetShortCityWeatherDataByName.
func (mr *MockApiMockRecorder) GetShortCityWeatherDataByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShortCityWeatherDataByName", reflect.TypeOf((*MockApi)(nil).GetShortCityWeatherDataByName), arg0)
}
