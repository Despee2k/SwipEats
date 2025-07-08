import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { AuthService } from '../../services/auth';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [RouterLink, FormsModule, CommonModule],
  templateUrl: './login.html',
  styleUrl: './login.css'
})
export class Login {
  email = '';
  password = '';
  error = '';
  message = '';

  constructor(private auth: AuthService, private router: Router) {}

  onLoginSubmit() {
    this.auth.login({ email: this.email, password: this.password }).subscribe({
      next: (res) => {
        this.error = '';
        this.message = 'Login successful! Redirecting...';
        this.auth.storeToken(res.token);
        setTimeout(() => this.router.navigate(['/lobby']), 1500);
      },
      error: (err) => {
        this.message = '';
        this.error = err?.error || 'Login failed. Check credentials.';
      }
    });
  }
}
