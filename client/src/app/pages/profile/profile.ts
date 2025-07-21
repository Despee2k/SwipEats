import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../services/auth/auth';
import { UserService } from '../../services/user/user';
import { API_URL_V1 } from '../../utils/constant';
import { NavigationBar } from '../../components/navigation-bar/navigation-bar';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-profile',
  imports: [NavigationBar, FormsModule],
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

  constructor(
    private authService: AuthService,
    private userService: UserService
  ) {}

  ngOnInit(): void {
    const token = this.authService.getToken();
    if (!token) return;

    this.userService.getUserDetails(token).subscribe({
      next: (res) => {
        if (!res.data) return;

        this.name = res.data.name;
        this.email = res.data.email;
        this.profilePictureUrl = `${API_URL_V1}/uploads/${res.data.email}`;
        this.originalProfilePictureUrl = this.profilePictureUrl;
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
      this.newProfilePicture = file;
      this.profilePictureUrl = URL.createObjectURL(file);
    }
  }

  toggleEdit(): void {
    this.isEditable = !this.isEditable;
    if (!this.isEditable) {
      this.password = '';
      this.newProfilePicture = null;
      this.profilePictureUrl = this.originalProfilePictureUrl;
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
    }).subscribe((res) => {
      if (!res.data) return;
      
      this.password = '';
      this.newProfilePicture = null;
      this.isEditable = false;
    });
  }

  logout(): void {
    this.authService.logout();
    window.location.href = '/login';
  }
}
