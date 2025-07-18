import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Location } from '@angular/common';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-navigation-bar',
  imports: [RouterLink],
  templateUrl: './navigation-bar.html',
  styleUrl: './navigation-bar.css'
})
export class NavigationBar implements OnInit {
  isLobbyPage = false;

  constructor(private router: Router, private location: Location) {}

  ngOnInit(): void {
    this.isLobbyPage = this.router.url.includes('/lobby');
  }

  goBack(): void {
    this.location.back();
  }
}