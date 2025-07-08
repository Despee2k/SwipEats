import { Component } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AuthService } from '../../services/auth';

@Component({
  selector: 'app-signup',
  standalone: true,
  imports: [RouterLink, FormsModule, CommonModule],
  templateUrl: './signup.html',
  styleUrl: './signup.css'
})
export class Signup {
  email = '';
  password = '';
  confirmPassword = '';
  message = '';
  error = '';

  constructor(private auth: AuthService, private router: Router) {}

  onSignupSubmit() {
    if (this.password !== this.confirmPassword) {
      this.error = 'Passwords do not match.';
      this.message = '';
      return;
    }

    this.auth.signup({
      email: this.email,
      password: this.password,
      confirm_password: this.confirmPassword,
    }).subscribe({
      next: () => {
        this.error = '';
        this.message = 'Account created! Redirecting...';
        setTimeout(() => this.router.navigate(['/login']), 1500);
      },
      error: (err) => {
        this.message = '';
        this.error = err?.error || 'Signup failed. Please try again.';
      }
    });
  }
}
