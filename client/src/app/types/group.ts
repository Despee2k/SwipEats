export enum GroupStatusEnum {
    ACTIVE = 'active',
    CLOSED = 'matched',
    WAITING = 'waiting'
}

export type GetGroupResponse = {
    group_code: string;
    name: string;
    location_lat: string;
    location_long: string;
    is_owner: boolean;
    group_status: GroupStatusEnum;
    member_count: number;
    created_at: string;
}

export interface CreateGroupResponse {
  group_code: string;
}
