import { Injectable } from '@angular/core';
import { API_URL_V1 } from '../../utils/constant';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { GetGroupResponse, CreateGroupResponse } from '../../types/group';
import { APIResponse } from '../../types/api';

@Injectable({
  providedIn: 'root'
})
export class GroupService {
  private GROUP_URL = API_URL_V1 + '/group';

  constructor(private http: HttpClient) {}

  private authHeader(token: string): HttpHeaders {
    return new HttpHeaders({
      'Authorization': `Bearer ${token}`
    });
  }

  fetchUserGroups(token: string): Observable<APIResponse<GetGroupResponse[]>> {
    return this.http.get<APIResponse<GetGroupResponse[]>>(`${this.GROUP_URL}/`, {
      headers: this.authHeader(token)
    });
  }

  createGroup(token: string, payload: { name: string; location_lat: number; location_long: number }): Observable<APIResponse<CreateGroupResponse>> {
    return this.http.post<APIResponse<CreateGroupResponse>>(`${this.GROUP_URL}/create`, payload, {
      headers: this.authHeader(token)
    });
  }

  joinGroup(token: string, groupCode: string): Observable<APIResponse<{ message: string }>> {
    return this.http.post<APIResponse<{ message: string }>>(`${this.GROUP_URL}/${groupCode}/join`, {}, {
      headers: this.authHeader(token)
    });
  }
}
