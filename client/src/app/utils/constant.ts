export const PRODUCTION = true;

export const BASE_URL = 'https://swipeats-api.dcism.org';
export const BASE_URL_LOCAL = 'http://localhost:42562';
export const API_URL_V1 = `${PRODUCTION ? BASE_URL : BASE_URL_LOCAL}/api/v1`;