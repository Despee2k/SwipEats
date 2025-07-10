import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { Router, RouterLink } from '@angular/router';

@Component({
  selector: 'app-landing',
  imports: [CommonModule],
  templateUrl: './landing.html',
  styleUrl: './landing.css'
})
export class Landing {
  moved = false;

  constructor(private router: Router) {}

  ngAfterViewInit() {
    setTimeout(() => {
      this.moved = true;
      setTimeout(() => {
        this.router.navigate(['/login'])
      }, 300)
    }, 3000);
  }

  ngOnInit() {
    const [navEntry] = performance.getEntriesByType("navigation") as PerformanceNavigationTiming[];
    if (navEntry?.type === 'reload') {
      // This was a browser reload
      this.moved = false;
    }
  }

}
