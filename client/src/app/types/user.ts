export interface UserDetails {
  id: number;
  name: string;
  email: string;
  profile_picture: string;
}

export interface UpdateUserPayload {
  name?: string;
  password?: string;
  clear_image?: boolean;
  profile_picture?: File | null;
}

export interface UpdateUserResponse {
  name: string;
  profile_picture: string;
}
