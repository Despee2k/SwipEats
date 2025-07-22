package dtos

type OverpassTagsResponseDto struct {
	Cuisine string `json:"cuisine"`
	Name    string `json:"name"`
}

type OverpassRestaurantResponseDto struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Tags    OverpassTagsResponseDto `json:"tags"`
}

type OverpassResponseDto struct {
	Elements []OverpassRestaurantResponseDto `json:"elements"`
}