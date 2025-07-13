import { Component, ElementRef, ViewChild } from '@angular/core';
import { AuthService } from '../../services/auth/auth';
import { Router, RouterLink } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { InputField } from '../../components/input-field/input-field';
import {
  trigger,
  transition,
  style,
  animate
} from '@angular/animations';

@Component({
  selector: 'app-login',
  imports: [CommonModule, RouterLink, FormsModule, InputField],
  templateUrl: './login.html',
  styleUrl: './login.css',
  animations: [
    trigger('fadeIn', [
      transition(':enter', [ // when element enters the DOM
        style({ opacity: 0, transform: 'translateY(10px)' }),
        animate('300ms ease-out', style({ opacity: 1, transform: 'translateY(0)' }))
      ])
    ])
  ]
})
export class Login {
  showElement = false;
  email = '';
  password = '';
  @ViewChild('submitButton') submitButtonRef!: ElementRef<HTMLButtonElement>;
  
  constructor(private auth: AuthService, private router: Router, private toastr: ToastrService) {}

  ngAfterViewInit() {
    setTimeout(() => {
      this.showElement = true;
    }, 100);
  }

  disableButton() {
    if (this.submitButtonRef.nativeElement) {
      this.submitButtonRef.nativeElement.disabled = true; // Disable button to prevent multiple clicks
    }
  }

  enableButton() {
    if (this.submitButtonRef.nativeElement) {
      this.submitButtonRef.nativeElement.disabled = false; // Enable button
    }
  }

  onLoginSubmit() {
    this.disableButton();

    this.auth.login({ email: this.email, password: this.password }).subscribe({
      next: (res) => {
        if (!res || !res.data || !res.data.token) {
          this.toastr.error('Login failed. Invalid response from server.', 'Error');
          this.enableButton(); // Re-enable button on error
          return;
        }
        this.toastr.success('Login successful! Redirecting...', 'Success');
        this.auth.storeToken(res.data.token);
        setTimeout(() => this.router.navigate(['/lobby']), 1500);
      },
      error: (err) => {
        this.toastr.error(err?.error?.message || 'Login failed. Check credentials.', 'Error');
        this.enableButton(); // Re-enable button on error
      }
    });
  }
}
