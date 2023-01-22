import Axios from "axios";

const axios = Axios.create({
  baseURL: `${import.meta.env.VITE_BASE_URL}`,
});
// axios.interceptors.request.use(
//   (config) => {
//     const tokenJson = localStorage.getItem("token");
//     const token = tokenJson ? JSON.parse(tokenJson) : "";
//     if (token) {
//       config.headers = { Authorization: `Bearer ${token}` };
//     }
//     return config;
//   },
//   (error) => Promise.reject(error)
// );

export default axios;
