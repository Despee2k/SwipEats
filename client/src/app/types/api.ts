export type APIResponse<T> = {
    message: string;
    data?: T;
    details?: { [key: string]: any };
}