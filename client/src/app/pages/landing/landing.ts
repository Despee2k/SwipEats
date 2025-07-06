import { CommonModule } from '@angular/common';
import { Component, ElementRef, ViewChild } from '@angular/core';

@Component({
  selector: 'app-landing',
  imports: [CommonModule],
  templateUrl: './landing.html',
  styleUrl: './landing.css'
})
export class Landing {
  moved = false;

    ngAfterViewInit() {
      setTimeout(() => {
        this.moved = true;
      }, 3000);
    }
}
