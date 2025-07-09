import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { AuthService } from '../../services/auth';
import { ToastrService } from 'ngx-toastr';

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

  constructor(private auth: AuthService, private router: Router, private toastr: ToastrService) {}

  onLoginSubmit() {
    this.auth.login({ email: this.email, password: this.password }).subscribe({
      next: (res) => {
        this.toastr.success('Login successful! Redirecting...', 'Success');
        this.auth.storeToken(res.token);
        setTimeout(() => this.router.navigate(['/lobby']), 1500);
      },
      error: (err) => {
        this.toastr.error(err?.error?.message || 'Login failed. Check credentials.', 'Error');
      }
    });
  }
}
