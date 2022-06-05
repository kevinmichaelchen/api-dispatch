export const BASE_URL = `http://localhost:8081`;

export const BASE_INIT: RequestInit = {
  method: "POST",
  mode: "cors",
  headers: {
    "Content-Type": "application/json",
  },
};
