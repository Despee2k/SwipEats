import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../services/auth/auth';
import { UserService } from '../../services/user/user';
import { API_URL_V1 } from '../../utils/constant';
import { NavigationBar } from '../../components/navigation-bar/navigation-bar';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-profile',
  imports: [NavigationBar, FormsModule, CommonModule],
  templateUrl: './profile.html',
  styleUrl: './profile.css'
})
export class Profile implements OnInit {
  isEditable = false;

  name: string = '';
  password: string = '';
  email: string = '';
  profilePictureUrl: string = '';
  newProfilePicture: File | null = null;
  originalProfilePictureUrl: string = '';
  imageLoadFailed: boolean = false;

  constructor(
    private authService: AuthService,
    private userService: UserService,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
    const token = this.authService.getToken();
    if (!token) return;

    this.userService.getUserDetails(token).subscribe({
      next: (res) => {
        if (!res.data) return;

        this.imageLoadFailed = false;
        this.name = res.data.name;
        this.email = res.data.email;
        this.originalProfilePictureUrl = `${API_URL_V1}/uploads/${encodeURIComponent(res.data.email)}`;
        this.profilePictureUrl = this.originalProfilePictureUrl + '?t=' + new Date().getTime(); // Prevent caching
      },
      error: (err) => {
        console.error('Failed to fetch user details', err);
      }
    });
  }

  decodeEmailFromToken(token: string): string {
    try {
      const payload = JSON.parse(atob(token.split('.')[1]));
      return payload.email;
    } catch {
      return '';
    }
  }

  onFileSelected(event: Event): void {
    const file = (event.target as HTMLInputElement)?.files?.[0];
    if (file) {
      this.imageLoadFailed = false;
      this.newProfilePicture = file;
      this.profilePictureUrl = URL.createObjectURL(file);
    }
  }

  toggleEdit(): void {
    this.isEditable = !this.isEditable;
    if (!this.isEditable) {
      this.password = '';
      this.newProfilePicture = null;
      this.profilePictureUrl = this.originalProfilePictureUrl + '?t=' + new Date().getTime(); // Reset to original picture
    }
  }

  saveProfile(): void {
    const token = this.authService.getToken();
    if (!token) return;

    this.userService.updateUser(token, {
      name: this.name,
      password: this.password,
      profile_picture: this.newProfilePicture,
      clear_image: false
    }).subscribe({
      next: (res) => {
        if (!res.data) return;

        this.password = '';
        this.newProfilePicture = null;
        this.isEditable = false;
        this.toastr.success('Profile updated successfully', 'Success');
      },
      error: (err) => {
        this.toastr.error('Failed to update profile', 'Error');
        console.error('Error updating profile', err);
      }
    });
  }

  logout(): void {
    this.authService.logout();
    window.location.href = '/login';
  }
}
