import { CommonModule } from '@angular/common';
import { Component, HostListener, OnInit } from '@angular/core';

@Component({
  selector: 'app-group-started',
  imports: [CommonModule],
  templateUrl: './group-started.html',
  styleUrl: './group-started.css'
})
export class GroupStarted implements OnInit {
  restaurantCards = [
    {
      name: 'Bella Italia',
      cuisine: 'Italian',
      rating: 4.5,
      price: '$$$',
      distance: '0.3 kilometers',
      description: 'Authentic Italian cuisine with fresh pasta and wood-fired pizzas',
    },
    {
      name: 'Sushi Zen',
      cuisine: 'Japanese',
      rating: 4.7,
      price: '$$$',
      distance: '1.2 kilometers',
      description: 'Fresh sushi and sashimi straight from the coast',
    },
    {
      name: 'Spicy Tandoori',
      cuisine: 'Indian',
      rating: 4.3,
      price: '$$',
      distance: '0.9 kilometers',
      description: 'Traditional Indian dishes with rich spices and aroma',
    },
  ];

  currentIndex = 0;
  startX = 0;
  currentDeltaX = 0;
  isDragging = false;
  transforms: string[] = [];

  ngOnInit(): void {
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
    setTimeout(() => this.nextCard(), 300);
  }

  swipeRight() {
    this.transforms[this.currentIndex] = 'translateX(120vw)';
    setTimeout(() => this.nextCard(), 300);
  }

  nextCard() {
    this.currentIndex++;
  }

  get isFinished(): boolean {
    return this.currentIndex >= this.restaurantCards.length;
  }
}
