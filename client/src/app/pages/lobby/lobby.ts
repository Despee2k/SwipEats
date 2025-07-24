import { Component, inject, PLATFORM_ID } from '@angular/core';
import { NavigationBar } from '../../components/navigation-bar/navigation-bar';
import { GroupService } from '../../services/group/group';
import { AuthService } from '../../services/auth/auth';
import { GetGroupResponse } from '../../types/group';
import { ToastrService } from 'ngx-toastr';
import { CommonModule, isPlatformBrowser, ɵnormalizeQueryParams } from '@angular/common';
import { RouterLink } from '@angular/router';
import { GroupRestaurant } from '../../types/restaurants';
import { MatchService } from '../../services/match/match';

@Component({
  selector: 'app-lobby',
  imports: [NavigationBar, CommonModule, RouterLink, ],
  templateUrl: './lobby.html',
  styleUrl: './lobby.css'
})
export class Lobby {
  private platformId = inject(PLATFORM_ID);
  getGroupResponse: GetGroupResponse[] = [];
  recentMatches: GroupRestaurant[] = [];

  constructor(private groupService: GroupService, private matchService: MatchService, private authService: AuthService, private toastr: ToastrService) {}

  ngOnInit() {
    if (isPlatformBrowser(this.platformId)) {
      const token = this.authService.getToken();
      if (token) {
        this.groupService.fetchUserGroups(token).subscribe({
          next: (res) => {
            if (!res) {
              this.toastr.error('Failed to fetch user groups.', 'Error');
              return;
            }

            if (res.data) {
              this.getGroupResponse = res.data.reverse();
            }
          },
          error: (error) => {
            this.toastr.error('Error fetching user groups:', error.message);
          }
        });

        this.matchService.fetchUserRecentMatches(token).subscribe({
          next: (res) => {
            if (!res || !res.data) {
              this.toastr.error('Failed to fetch recent matches.', 'Error');
              return;
            }
            this.recentMatches = res.data;
          },
          error: (error) => {
            this.toastr.error('Error fetching recent matches:', error.message);
          }
        });
      } else {
        this.toastr.error('No authentication token found.', 'Error');
        // Logout
      }
    }
  }
}
