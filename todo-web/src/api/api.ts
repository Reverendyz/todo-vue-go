// src/api/api.ts
import axios from 'axios'

export const api = axios.create({
  baseURL: 'http://localhost:8080',
  headers: {
    'Content-Type': 'application/json',
  },
})

api.interceptors.request.use((config) => {
  return config
})

api.interceptors.response.use(
  (response) => response,
  (error) => {
    return Promise.reject(error)
  },
)
