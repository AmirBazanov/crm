import { api } from './client';

export interface LoginRequest {
    email: string;
    password: string;
}

export interface RegisterRequest {
    email: string;
    username: string;
    password: string;
}

export interface AuthResponse {
    accessToken: string;
    refreshToken: string;
    user: {
        id: string;
        email: string;
        firstName: string;
        lastName: string;
        nickname: string;
        country: number;
        createdAt: string;
        updatedAt: string;
    };
}

export interface RefreshRequest {
    refreshToken: string;
}

export interface RefreshResponse {
    accessToken: string;
    refreshToken: string;
}

export const authService = {
    async login(data: LoginRequest): Promise<AuthResponse> {
        return api.post<AuthResponse>('/auth/login', data);
    },

    async register(data: RegisterRequest): Promise<AuthResponse> {
        return api.post<AuthResponse>('/auth/register', data);
    },

    async refresh(data: RefreshRequest): Promise<RefreshResponse> {
        return api.post<RefreshResponse>('/auth/refresh', data);
    },

    async logout(refreshToken: string): Promise<{ message: string }> {
        return api.post('/auth/logout', { refreshToken });
    },

    // Token management
    getAccessToken(): string | null {
        return localStorage.getItem('accessToken');
    },

    getRefreshToken(): string | null {
        return localStorage.getItem('refreshToken');
    },

    setTokens(accessToken: string, refreshToken: string): void {
        localStorage.setItem('accessToken', accessToken);
        localStorage.setItem('refreshToken', refreshToken);
    },

    clearTokens(): void {
        localStorage.removeItem('accessToken');
        localStorage.removeItem('refreshToken');
        localStorage.removeItem('user');
    },

    setUser(user: any): void {
        localStorage.setItem('user', JSON.stringify(user));
    },

    getUser(): any | null {
        const user = localStorage.getItem('user');
        if (!user || user === 'undefined' || user === 'null') {
            return null;
        }
        try {
            return JSON.parse(user);
        } catch (e) {
            console.error('Failed to parse user data:', e);
            return null;
        }
    },

    isAuthenticated(): boolean {
        return !!this.getAccessToken();
    },
};
