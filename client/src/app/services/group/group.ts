import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { GetGroupResponse, CreateGroupResponse, JoinGroupResponse, GroupMember } from '../../types/group';
import { APIResponse } from '../../types/api';
import { API_URL_V1 } from '../../utils/constant';

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

  createGroup(
    token: string,
    payload: { name: string; location_lat: number; location_long: number }
  ): Observable<APIResponse<CreateGroupResponse>> {
    return this.http.post<APIResponse<CreateGroupResponse>>(`${this.GROUP_URL}/create`, payload, {
      headers: this.authHeader(token)
    });
  }

  joinGroup(token: string, groupCode: string): Observable<APIResponse<JoinGroupResponse>> {
    return this.http.post<APIResponse<JoinGroupResponse>>(`${this.GROUP_URL}/${groupCode}/join`, {}, {
      headers: this.authHeader(token)
    });
  }

  fetchGroupMembers(token: string, groupCode: string): Observable<APIResponse<GroupMember[]>> {
    return this.http.get<APIResponse<GroupMember[]>>(`${this.GROUP_URL}/${groupCode}/members`, {
      headers: this.authHeader(token)
    });
  }
}
