import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { API_URL_V1 } from '../../utils/constant';
import { UpdateUserPayload, UpdateUserResponse, UserDetails } from '../../types/user';
import { APIResponse } from '../../types/api';

@Injectable({ providedIn: 'root' })
export class UserService {
  private USER_URL = `${API_URL_V1}/user`;

  constructor(private http: HttpClient) {}

  getUserDetails(token: string): Observable<APIResponse<UserDetails>> {
    return this.http.get<APIResponse<UserDetails>>(this.USER_URL, {
      headers: new HttpHeaders({
        Authorization: `Bearer ${token}`
      })
    });
  }

  updateUser(token: string, payload: UpdateUserPayload): Observable<APIResponse<UpdateUserResponse>> {
    const formData = new FormData();

    if (payload.name) formData.append('name', payload.name);
    if (payload.password) formData.append('password', payload.password);
    if (payload.clear_image !== undefined) formData.append('clear_image', payload.clear_image.toString());
    if (payload.profile_picture) formData.append('profile_picture', payload.profile_picture);

    return this.http.patch<APIResponse<UpdateUserResponse>>(
      `${this.USER_URL}/update`,
      formData,
      {
        headers: new HttpHeaders({
          Authorization: `Bearer ${token}`
        })
      }
    );
  }
}
