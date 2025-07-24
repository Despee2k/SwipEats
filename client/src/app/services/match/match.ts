import { Injectable } from '@angular/core';
import { API_URL_V1 } from '../../utils/constant';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { APIResponse } from '../../types/api';
import { GroupRestaurant } from '../../types/restaurants';

@Injectable({
  providedIn: 'root'
})
export class MatchService {
  private MATCH_URL = `${API_URL_V1}/match`;
  public socket: WebSocket | null = null;

  constructor(private http: HttpClient) {}

  private authHeader(token: string): HttpHeaders {
    return new HttpHeaders({
      'Authorization': `Bearer ${token}`
    });
  }

  fetchUserRecentMatches(token: string): Observable<APIResponse<GroupRestaurant[]>> {
    return this.http.get<APIResponse<GroupRestaurant[]>>(`${this.MATCH_URL}/`, {
      headers: this.authHeader(token)
    });
  }
}