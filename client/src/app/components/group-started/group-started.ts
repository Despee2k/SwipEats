import { CommonModule } from '@angular/common';
import { Component, HostListener, Input, OnInit } from '@angular/core';
import { GroupRestaurant } from '../../types/restaurants';
import { trimToFirstDelimiter } from '../../utils/general';

@Component({
  selector: 'app-group-started',
  imports: [CommonModule],
  templateUrl: './group-started.html',
  styleUrl: './group-started.css'
})
export class GroupStarted implements OnInit {
  @Input() groupRestaurants!: GroupRestaurant[]
  @Input() alreadyVoted!: boolean;
  @Input() isOwner!: boolean;

  @Input() handleVote!: (groupRestaurantId: number, vote: boolean) => void;
  @Input() handleSubmit!: () => void;
  @Input() endGroup!: () => void;

  restaurantCards: { name: string; cuisine: string; distance: string, photo_url: string }[] = [];

  currentIndex = 0;
  startX = 0;
  currentDeltaX = 0;
  isDragging = false;
  transforms: string[] = [];

  ngOnInit(): void {
    this.restaurantCards = this.groupRestaurants.map(groupRestaurant => {
      const { restaurant, distance_in_km } = groupRestaurant;

      const displayed_distance = (distance_in_km < 1) ? distance_in_km.toFixed(1) + ' km' : Math.round(distance_in_km) + ' km';

      return {
        name: restaurant.name,
        cuisine: trimToFirstDelimiter(restaurant.cuisine),
        distance: displayed_distance,
        photo_url: restaurant.photo_url
      };
    });

    this.transforms = this.restaurantCards.map(() => 'translateX(0px)');
  }

  // TOUCH HANDLING
  @HostListener('touchstart', ['$event']) onTouchStart(event: TouchEvent) {
    this.startDrag(event.touches[0].clientX);
  }

  @HostListener('touchmove', ['$event']) onTouchMove(event: TouchEvent) {
    this.moveDrag(event.touches[0].clientX);
  }

  @HostListener('touchend') onTouchEnd() {
    this.endDrag();
  }

  // MOUSE HANDLING
  @HostListener('mousedown', ['$event']) onMouseDown(event: MouseEvent) {
    if (this.isFinished) return;
    this.startDrag(event.clientX);
    document.addEventListener('mousemove', this.mouseMove);
    document.addEventListener('mouseup', this.mouseUp);
  }

  mouseMove = (event: MouseEvent) => {
    this.moveDrag(event.clientX);
  };

  mouseUp = () => {
    this.endDrag();
    document.removeEventListener('mousemove', this.mouseMove);
    document.removeEventListener('mouseup', this.mouseUp);
  };

  // SHARED DRAG METHODS
  startDrag(x: number) {
    this.startX = x;
    this.isDragging = true;
  }

  moveDrag(x: number) {
    if (!this.isDragging || this.isFinished) return;
    this.currentDeltaX = x - this.startX;
    this.transforms[this.currentIndex] = `translateX(${this.currentDeltaX}px)`;
  }

  endDrag() {
    if (!this.isDragging || this.isFinished) return;
    this.isDragging = false;

    const threshold = 100;
    if (this.currentDeltaX > threshold) {
      this.swipeRight();
    } else if (this.currentDeltaX < -threshold) {
      this.swipeLeft();
    } else {
      this.transforms[this.currentIndex] = 'translateX(0px)';
    }

    this.currentDeltaX = 0;
  }

  swipeLeft() {
    this.transforms[this.currentIndex] = 'translateX(-120vw)';
    this.handleVote(this.groupRestaurants[this.currentIndex].id, false);
    setTimeout(() => this.nextCard(), 300);
  }

  swipeRight() {
    this.transforms[this.currentIndex] = 'translateX(120vw)';
    this.handleVote(this.groupRestaurants[this.currentIndex].id, true);
    setTimeout(() => this.nextCard(), 300);
  }

  nextCard() {
    this.currentIndex++;

    if (this.currentIndex >= this.restaurantCards.length) {
      this.handleSubmit();
      return;
    }
  }

  get isFinished(): boolean {
    return this.alreadyVoted || this.currentIndex >= this.restaurantCards.length;
  }
}
