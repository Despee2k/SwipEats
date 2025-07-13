import { Component } from '@angular/core';
import { NavigationBar } from '../../components/navigation-bar/navigation-bar';
import { GroupService } from '../../services/group/group';
import { AuthService } from '../../services/auth/auth';
import { ToastrService } from 'ngx-toastr';
import { Router, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-group',
  imports: [NavigationBar, FormsModule, RouterLink],
  templateUrl: './group.html',
  styleUrl: './group.css'
})
export class Group {
  activeModal: 'create' | 'join' | 'success' = 'create';

  groupName: string = '';
  groupCode: string = '';
  joinCode: string = '';
  isLoading: boolean = false;

  constructor(
    private groupService: GroupService,
    private authService: AuthService,
    private toastr: ToastrService,
    private router: Router
  ) {}

  onCreateGroup(): void {
    if (!this.groupName.trim()) {
      this.toastr.error('Group name is required');
      return;
    }

    this.isLoading = true;

    const token = this.authService.getToken();
    if (!token) {
      this.toastr.error('No authentication token found');
      this.isLoading = false;
      return;
    }

    navigator.geolocation.getCurrentPosition(
      (position) => {
        const payload = {
          name: this.groupName,
          location_lat: position.coords.latitude,
          location_long: position.coords.longitude
        };

        this.groupService.createGroup(token, payload).subscribe({
          next: (res) => {
            if (!res.data?.group_code) {
              this.toastr.error('Group code missing in response');
              this.isLoading = false;
              return;
            }
            this.groupCode = res.data.group_code;
            this.activeModal = 'success';
            this.isLoading = false;
          },
          error: (err) => {
            this.toastr.error(err?.error?.message || 'Failed to create group');
            this.isLoading = false;
          }
        });
      },
      () => {
        this.toastr.error('Location access denied');
        this.isLoading = false;
      }
    );
  }

  onJoinGroup(): void {
    if (!this.joinCode.trim()) {
      this.toastr.error('Group code is required');
      return;
    }

    const token = this.authService.getToken();
    if (!token) {
      this.toastr.error('No authentication token found');
      return;
    }

    this.groupService.joinGroup(token, this.joinCode).subscribe({
      next: () => {
        this.toastr.success('Successfully joined the group!');
        this.router.navigate(['/group', this.joinCode]);
      },
      error: (err) => {
        this.toastr.error(err?.error?.message || 'Failed to join group');
      }
    });
  }

  copyGroupCode(): void {
    navigator.clipboard.writeText(this.groupCode).then(() => {
      this.toastr.success('Group code copied!');
    });
  }
}