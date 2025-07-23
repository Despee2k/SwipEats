import { Component, Input, OnInit } from '@angular/core';
import { GroupRestaurant } from '../../types/restaurants';
import { CommonModule } from '@angular/common';
import { ToastrService } from 'ngx-toastr';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-group-match',
  standalone: true,
  imports: [CommonModule, RouterLink],
  templateUrl: './group-match.html',
  styleUrl: './group-match.css'
})
export class GroupMatch implements OnInit {
  @Input() finalRestaurant!: GroupRestaurant | null;

  // Only required fields based on type structure
  imageUrl: string = '';
  distance: string = '';
  name: string = '';
  cuisine: string = '';

  constructor(private toastr: ToastrService) {}

  ngOnInit(): void {
    if (!this.finalRestaurant) {
      this.toastr.error('Final restaurant data is not provided', 'Error');
      return;
    }

    const { restaurant, distance_in_km } = this.finalRestaurant;

    this.imageUrl = restaurant.photo_url || '';
    this.distance = distance_in_km < 1
      ? `${distance_in_km.toFixed(1)} km`
      : `${Math.round(distance_in_km)} km`;

    this.name = restaurant.name;
    this.cuisine = restaurant.cuisine?.split(',')[0] || '';
  }
}
