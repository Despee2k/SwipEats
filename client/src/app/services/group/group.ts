import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { GetGroupResponse, CreateGroupResponse, GroupMember } from '../../types/group';
import { APIResponse } from '../../types/api';
import { API_URL_V1 } from '../../utils/constant';

@Injectable({
  providedIn: 'root'
})
export class GroupService {
  private GROUP_URL = `${API_URL_V1}/group`;
  private WS_URL = `wss://swipeats-api.dcism.org/ws`;
  public socket: WebSocket | null = null;

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

  fetchGroupMembers(token: string, groupCode: string): Observable<APIResponse<GroupMember[]>> {
    return this.http.get<APIResponse<GroupMember[]>>(`${this.GROUP_URL}/${groupCode}/members`, {
      headers: this.authHeader(token)
    });
  }

  connectWebSocket(
    token: string,
    groupCode: string,
    onSuccess: (data: any) => void,
    onError: (err: any) => void
  ): void {
    if (this.socket) {
      this.socket.close();
    }

    this.socket = new WebSocket(this.WS_URL);

    this.socket.onopen = () => {
      const joinPayload = {
        type: 'join_group',
        token,
        group_code: groupCode
      };
      this.socket?.send(JSON.stringify(joinPayload));
    };

    this.socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.type === 'join_group_success') {
        onSuccess(data);
      } else if (data.type === 'join_group_error') {
        onError(data);
      }
    };

    this.socket.onerror = (err) => {
      onError(err);
    };
  }
}
