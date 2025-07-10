import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { API_URL_V1 } from '../utils/constant';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private AUTH_URL = API_URL_V1 + '/auth';

  constructor(private http: HttpClient) {}

  signup(payload: { email: string; password: string; confirm_password: string }): Observable<any> {
    return this.http.post(`${this.AUTH_URL}/signup`, payload);
  }

  login(payload: { email: string; password: string }): Observable<{ token: string }> {
    return this.http.post<{ token: string }>(`${this.AUTH_URL}/login`, payload);
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