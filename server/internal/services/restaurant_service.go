package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/SwipEats/SwipEats/server/internal/constants"
	"github.com/SwipEats/SwipEats/server/internal/dtos"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"github.com/SwipEats/SwipEats/server/internal/repositories"
)

func fetchImageUrl(cuisine string) string {
	query := fmt.Sprintf("orientation=landscape&query=%s%%20restaurant%%20food&client_id=%s", url.QueryEscape(cuisine), constants.UNSPLASHED_KEY)
	urlWithQuery := "https://api.unsplash.com/photos/random?" + query

	resp, err := http.Get(urlWithQuery)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var imageResponse struct {
		Urls struct {
			Raw string `json:"raw"`
		} `json:"urls"`
	}

	if err := json.Unmarshal(body, &imageResponse); err != nil {
		log.Fatal(err)
		return ""
	}

	return imageResponse.Urls.Raw
}

func FetchRestaurantsNearby(lat, long float64, radius int) ([]models.Restaurant, error) {
	query := fmt.Sprintf(`[out:json];node(around:%d,%f,%f)[amenity=restaurant][cuisine][name];out body 10;`, radius, lat, long)
	urlWithQuery := "https://overpass-api.de/api/interpreter?data=" + url.QueryEscape(query)

	resp, err := http.Get(urlWithQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	var overpassResponse dtos.OverpassResponseDto
	if err := json.Unmarshal(body, &overpassResponse); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var restaurants []models.Restaurant
	for _, element := range overpassResponse.Elements {
		existingRestaurant, err := repositories.GetRestaurantByName(element.Tags.Name)
		if err != nil {
			log.Printf("Error checking existing restaurant: %v", err)
			continue
		}
		if existingRestaurant != nil {
			log.Printf("Restaurant with name '%s' already exists, skipping.", element.Tags.Name)
			continue
		}

		restaurant := models.Restaurant{
			Name:        element.Tags.Name,
			LocationLat:  element.Lat,
			LocationLong: element.Lon,
			Cuisine:    element.Tags.Cuisine,
			PhotoURL:   fetchImageUrl(element.Tags.Cuisine),
		}
		if err := repositories.AddRestaurant(&restaurant); err != nil {
			log.Printf("Error adding restaurant: %v", err)
			continue
		}

		restaurants = append(restaurants, restaurant)
	}

	return restaurants, nil
}