package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"saferoute/pkg/config"
)

type Route struct {
	Geometry struct {
		Coordinates [][]float64 `json:"coordinates"`
	} `json:"geometry"`
}

func GetSafeRoute(start, end string) (Route, error) {
	cfg := config.Load()
	baseURL := "https://api.mapbox.com/directions/v5/mapbox/walking"

	params := url.Values{}
	params.Add("access_token", cfg.MapboxToken)
	params.Add("geometries", "geojson")
	params.Add("annotations", "congestion")
	params.Add("avoid_features", "poor_lighting")

	url := fmt.Sprintf("%s/%s;%s?%s", baseURL, start, end, params.Encode())

	resp, err := http.Get(url)
	if err != nil {
		return Route{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Route{}, err
	}

	var route Route
	if err := json.Unmarshal(body, &route); err != nil {
		return Route{}, err
	}

	return route, nil
}

func GetSafeRoute(start, end string) (Route, error) {
	cfg := config.Load()
	baseURL := "https://api.mapbox.com/directions/v5/mapbox/walking"

	params := url.Values{}
	params.Add("access_token", cfg.MapboxToken)
	params.Add("geometries", "geojson")
	params.Add("annotations", "congestion")
	params.Add("avoid_features", "poor_lighting")

	url := fmt.Sprintf("%s/%s;%s?%s", baseURL, start, end, params.Encode())

	resp, err := http.Get(url)
	if err != nil {
		return Route{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Route{}, err
	}

	var route Route
	if err := json.Unmarshal(body, &route); err != nil {
		return Route{}, err
	}

	return route, nil
}
