import { CommonModule, isPlatformBrowser } from '@angular/common';
import { Component, Inject, PLATFORM_ID } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-landing',
  imports: [CommonModule],
  templateUrl: './landing.html',
  styleUrl: './landing.css'
})
export class Landing {
  moved = false;

  constructor(private router: Router, @Inject(PLATFORM_ID) private platformId: Object) {}

  ngAfterViewInit() {
    if (isPlatformBrowser(this.platformId)) {
      requestAnimationFrame(() => {
        setTimeout(() => {
          this.moved = true;
          setTimeout(() => {
            this.router.navigate(['/login'])
          }, 300)
        }, 3000);
      });
    }
  }

  ngOnInit() {
    const [navEntry] = performance.getEntriesByType("navigation") as PerformanceNavigationTiming[];
    if (navEntry?.type === 'reload') {
      // This was a browser reload
      this.moved = false;
    }
  }

}
