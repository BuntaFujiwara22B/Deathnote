import axios from 'axios';

// Interface para los datos de víctima
interface Victim {
  id: number;
  full_name: string;
  cause?: string;
  details?: string;
  created_at: string;
  image_url: string;
}

// Configuración de axios con tipos inferidos
const api = axios.create({
  baseURL: 'http://localhost:8080',
  headers: {
    'Content-Type': 'application/json'
  }
});

// Interceptor para respuestas
api.interceptors.response.use(
  response => response.data,
  error => {
    console.error('API Error:', error.response?.data || error.message);
    return Promise.reject(error);
  }
);

// Exportación de métodos de la API
export const registerVictim = (data: Omit<Victim, 'id' | 'created_at'>) => {
  return api.post<Victim>('/victimas', data);
};

export const getVictims = () => {
  return api.get<Victim[]>('/victimas');
};

export const updateCause = (id: number, cause: string) => {
  return api.put<Victim>(`/victimas/${id}/cause`, { cause });
};

export const updateDetails = (id: number, details: string) => {
  return api.put<Victim>(`/victimas/${id}/details`, { details });
};