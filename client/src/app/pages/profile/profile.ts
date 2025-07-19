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

  constructor(
    private authService: AuthService,
    private userService: UserService
  ) {}

  ngOnInit(): void {
    const token = this.authService.getToken();
    if (!token) return;

    const email = this.decodeEmailFromToken(token);
    this.email = email;
    this.profilePictureUrl = `${API_URL_V1}/user/profile-picture/${email}`;
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
    }
  }

  toggleEdit(): void {
    this.isEditable = !this.isEditable;
    if (!this.isEditable) {
      this.password = '';
      this.newProfilePicture = null;
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
      
      this.profilePictureUrl = res.data.profile_picture;
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
