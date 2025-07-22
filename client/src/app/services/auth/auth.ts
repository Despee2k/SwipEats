import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { API_URL_V1 } from '../../utils/constant';
import { APIResponse } from '../../types/api';
import { LoginResponse } from '../../types/auth';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private AUTH_URL = API_URL_V1 + '/auth';

  constructor(private http: HttpClient) {}

  signup(payload: { email: string; password: string; confirm_password: string }): Observable<any> {
    return this.http.post(`${this.AUTH_URL}/signup`, payload);
  }

  login(payload: { email: string; password: string }): Observable<APIResponse<LoginResponse>> {
    return this.http.post<APIResponse<LoginResponse>>(`${this.AUTH_URL}/login`, payload);
  }

  storeUserData(userData: LoginResponse) {
    if (typeof window !== 'undefined' && localStorage) {
      localStorage.setItem('user_id', userData.user_id.toString());
      localStorage.setItem('auth_token', userData.token);
    }
  }

  getUserId(): number | null {
    if (typeof window !== 'undefined' && localStorage) {
      const userId = localStorage.getItem('user_id');
      const parsed = parseInt(userId || '', 10);
      return isNaN(parsed) ? null : parsed;
    }
    return null;
}

  getToken(): string | null {
  if (typeof window !== 'undefined' && localStorage) {
      return localStorage.getItem('auth_token');
    }
    return null;
  }

  logout() {
    if (typeof window !== 'undefined' && localStorage) {
      localStorage.removeItem('auth_token');
      localStorage.removeItem('user_id');
    }
  }
}