export interface Restaurant {
    id: number;
    name: string;
    cuisine: string;
    photo_url: string;
}

export interface GroupRestaurant {
    id: number;
    group_id: number;
    restaurant: Restaurant;
    distance_in_km: number;
}