export const PRODUCTION = true;

export const PROD_URL = 'swipeats-api.dcism.org';
export const LOCAL_URL = 'localhost:42562';

export const BASE_API_URL = PRODUCTION ? `https://${PROD_URL}` : `http://${LOCAL_URL}`;
export const BASE_WS_URL = PRODUCTION ? `wss://${PROD_URL}` : `ws://${LOCAL_URL}`;

export const API_URL_V1 = `${BASE_API_URL}/api/v1`;