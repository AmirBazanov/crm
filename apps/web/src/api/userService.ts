import { authService } from './authService';

export interface UserProfile {
    id: string;
    email: string;
    firstName: string;
    lastName: string;
    nickname: string;
    country: number;
    createdAt: string;
    updatedAt: string;
    // Extended profile data (localStorage only)
    phone?: string;
    bio?: string;
    role?: string;
}

export const userService = {
    /**
     * Get user profile from localStorage
     * Combines auth user data with extended profile data
     */
    getUserProfile(): UserProfile | null {
        const user = authService.getUser();
        if (!user) return null;

        // Get extended profile data from localStorage
        const extendedProfile = this.getExtendedProfile();

        return {
            id: user.id,
            email: user.email,
            firstName: user.firstName || '',
            lastName: user.lastName || '',
            nickname: user.nickname || user.username || '', // Fallback to username if nickname not present
            country: user.country || 0,
            createdAt: user.createdAt || new Date().toISOString(),
            updatedAt: user.updatedAt || new Date().toISOString(),
            phone: extendedProfile?.phone || '',
            bio: extendedProfile?.bio || '',
            role: extendedProfile?.role || 'User',
        };
    },

    /**
     * Update user profile in localStorage
     */
    updateUserProfile(profile: Partial<UserProfile>): void {
        const currentUser = authService.getUser();
        if (!currentUser) return;

        // Update base user data if core fields changed
        if (profile.email || profile.nickname || profile.firstName || profile.lastName || profile.country) {
            authService.setUser({
                ...currentUser,
                email: profile.email || currentUser.email,
                nickname: profile.nickname || currentUser.nickname,
                firstName: profile.firstName || currentUser.firstName,
                lastName: profile.lastName || currentUser.lastName,
                country: profile.country !== undefined ? profile.country : currentUser.country,
            });
        }

        // Store extended profile data
        const extendedProfile = this.getExtendedProfile() || {};
        const updatedExtendedProfile = {
            ...extendedProfile,
            phone: profile.phone,
            bio: profile.bio,
            role: profile.role,
        };

        localStorage.setItem('userProfile', JSON.stringify(updatedExtendedProfile));
    },

    /**
     * Get extended profile data from localStorage
     */
    getExtendedProfile(): Partial<UserProfile> | null {
        const profile = localStorage.getItem('userProfile');
        return profile ? JSON.parse(profile) : null;
    },

    /**
     * Get user initials from name
     */
    getUserInitials(firstName?: string, lastName?: string, nickname?: string): string {
        if (firstName && lastName) {
            return (firstName[0] + lastName[0]).toUpperCase();
        }
        if (nickname) {
            return nickname.substring(0, 2).toUpperCase();
        }

        const user = authService.getUser();
        const name = user?.firstName || user?.nickname || user?.email || 'U';

        if (user?.firstName && user?.lastName) {
            return (user.firstName[0] + user.lastName[0]).toUpperCase();
        }

        return name.substring(0, 2).toUpperCase();
    },

    /**
     * Clear user profile data
     */
    clearProfile(): void {
        localStorage.removeItem('userProfile');
    },
};
