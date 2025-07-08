import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private BASE_URL = 'https://swipeats-api.dcism.org/api/v1/auth';

  constructor(private http: HttpClient) {}

  signup(payload: { email: string; password: string; confirm_password: string }): Observable<any> {
    return this.http.post(`${this.BASE_URL}/signup`, payload);
  }

  login(payload: { email: string; password: string }): Observable<{ token: string }> {
    return this.http.post<{ token: string }>(`${this.BASE_URL}/login`, payload);
  }

  storeToken(token: string) {
    localStorage.setItem('auth_token', token);
  }

  getToken(): string | null {
    return localStorage.getItem('auth_token');
  }

  logout() {
    localStorage.removeItem('auth_token');
  }
}