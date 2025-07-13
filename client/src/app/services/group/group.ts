import { Injectable } from '@angular/core';
import { API_URL_V1 } from '../../utils/constant';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { GetGroupResponse } from '../../types/group';
import { APIResponse } from '../../types/api';

@Injectable({
  providedIn: 'root'
})
export class GroupService {
  private GROUP_URL = API_URL_V1 + '/group';

  constructor(private http: HttpClient) { }

  fetchUserGroups(token: string): Observable<APIResponse<GetGroupResponse[]>> {
    return this.http.get<APIResponse<GetGroupResponse[]>>(`${this.GROUP_URL}/`, {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    });
  }
}
