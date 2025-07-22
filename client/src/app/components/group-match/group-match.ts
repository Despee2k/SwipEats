import { Component, Input } from '@angular/core';
import { GroupRestaurant } from '../../types/restaurants';
import { CommonModule } from '@angular/common';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-group-match',
  imports: [CommonModule],
  templateUrl: './group-match.html',
  styleUrl: './group-match.css'
})
export class GroupMatch {
  @Input() finalRestaurant!: GroupRestaurant | null;

  imageUrl: string = '';

  constructor(private toastr: ToastrService) {}

  ngOnInit(): void {
    if (!this.finalRestaurant) {
      this.toastr.error('Final restaurant data is not provided', 'Error');
      return;
    }

    this.imageUrl = this.finalRestaurant.restaurant.photo_url || '';
  }
}
