import { Component } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AuthService } from '../../services/auth';
import { ToastrService } from 'ngx-toastr';

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

  constructor(private auth: AuthService, private router: Router, private toastr: ToastrService) {}

  onSignupSubmit() {
    if (this.password !== this.confirmPassword) {
      this.toastr.error('Passwords do not match.', 'Error');
      return;
    }

    this.auth.signup({
      email: this.email,
      password: this.password,
      confirm_password: this.confirmPassword,
    }).subscribe({
      next: () => {
        this.toastr.success('Account created successfully!', 'Success');
        setTimeout(() => this.router.navigate(['/login']), 1500);
      },
      error: (err) => {
        this.toastr.error(err?.error?.message || 'Signup failed. Please try again.', 'Error');
      }
    });
  }
}
