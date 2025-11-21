import React, { useState, useEffect } from 'react';
import { Card } from '../components/ui/Card';
import { User, Mail, Phone, MapPin, Calendar, Edit2, Save, X, Shield, Bell, Lock } from 'lucide-react';
import { userService } from '../api/userService';

export const Profile: React.FC = () => {
    const [isEditing, setIsEditing] = useState(false);
    const [isLoading, setIsLoading] = useState(true);
    const [userData, setUserData] = useState({
        firstName: '',
        lastName: '',
        nickname: '',
        email: '',
        phone: '',
        country: 0,
        role: 'User',
        createdAt: '',
        bio: '',
    });

    const [editedData, setEditedData] = useState(userData);

    // Load user data on component mount
    useEffect(() => {
        const profile = userService.getUserProfile();
        if (profile) {
            const profileData = {
                firstName: profile.firstName,
                lastName: profile.lastName,
                nickname: profile.nickname,
                email: profile.email,
                phone: profile.phone || '',
                country: profile.country || 0,
                role: profile.role || 'User',
                createdAt: new Date(profile.createdAt).toLocaleDateString(),
                bio: profile.bio || '',
            };
            setUserData(profileData);
            setEditedData(profileData);
        }
        setIsLoading(false);
    }, []);

    const handleEdit = () => {
        setIsEditing(true);
        setEditedData(userData);
    };

    const handleSave = () => {
        // Save to localStorage via userService
        userService.updateUserProfile({
            firstName: editedData.firstName,
            lastName: editedData.lastName,
            nickname: editedData.nickname,
            email: editedData.email,
            phone: editedData.phone,
            country: typeof editedData.country === 'string' ? parseInt(editedData.country) : editedData.country,
            bio: editedData.bio,
            role: editedData.role,
        });
        setUserData(editedData);
        setIsEditing(false);
    };

    const handleCancel = () => {
        setEditedData(userData);
        setIsEditing(false);
    };

    const handleChange = (field: string, value: string | number) => {
        setEditedData(prev => ({ ...prev, [field]: value }));
    };

    const activityData = [
        { action: 'Updated customer record', time: '2 hours ago', type: 'update' },
        { action: 'Created new deal', time: '5 hours ago', type: 'create' },
        { action: 'Sent email campaign', time: '1 day ago', type: 'email' },
        { action: 'Generated monthly report', time: '2 days ago', type: 'report' },
        { action: 'Updated profile settings', time: '1 week ago', type: 'settings' },
    ];

    const preferences = [
        { icon: Bell, label: 'Email Notifications', enabled: true },
        { icon: Shield, label: 'Two-Factor Authentication', enabled: true },
        { icon: Lock, label: 'Privacy Mode', enabled: false },
    ];

    if (isLoading) {
        return <div className="flex justify-center items-center h-64">Loading profile...</div>;
    }

    return (
        <div className="flex flex-col gap-6">
            <div className="flex justify-between items-center">
                <h1 className="text-2xl font-bold">Profile</h1>
                {!isEditing ? (
                    <button
                        onClick={handleEdit}
                        className="btn-primary flex items-center gap-2"
                    >
                        <Edit2 size={16} />
                        Edit Profile
                    </button>
                ) : (
                    <div className="flex gap-2">
                        <button
                            onClick={handleSave}
                            className="btn-primary flex items-center gap-2"
                        >
                            <Save size={16} />
                            Save
                        </button>
                        <button
                            onClick={handleCancel}
                            className="btn-secondary flex items-center gap-2"
                        >
                            <X size={16} />
                            Cancel
                        </button>
                    </div>
                )}
            </div>

            <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                {/* Profile Card */}
                <Card className="lg:col-span-2">
                    <div className="flex flex-col md:flex-row gap-6">
                        {/* Avatar Section */}
                        <div className="flex flex-col items-center gap-4">
                            <div className="profile-avatar">
                                {userData.firstName || userData.nickname ? (
                                    <span className="text-4xl font-bold">
                                        {userService.getUserInitials(userData.firstName, userData.lastName, userData.nickname)}
                                    </span>
                                ) : (
                                    <User size={48} />
                                )}
                            </div>
                            {isEditing && (
                                <button className="text-sm text-accent hover:underline">
                                    Change Photo
                                </button>
                            )}
                        </div>

                        {/* Info Section */}
                        <div className="flex-1">
                            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                                <div className="profile-field">
                                    <label className="profile-label">
                                        <User size={16} />
                                        First Name
                                    </label>
                                    {isEditing ? (
                                        <input
                                            type="text"
                                            value={editedData.firstName}
                                            onChange={(e) => handleChange('firstName', e.target.value)}
                                            className="input-field"
                                        />
                                    ) : (
                                        <p className="profile-value">{userData.firstName}</p>
                                    )}
                                </div>

                                <div className="profile-field">
                                    <label className="profile-label">
                                        <User size={16} />
                                        Last Name
                                    </label>
                                    {isEditing ? (
                                        <input
                                            type="text"
                                            value={editedData.lastName}
                                            onChange={(e) => handleChange('lastName', e.target.value)}
                                            className="input-field"
                                        />
                                    ) : (
                                        <p className="profile-value">{userData.lastName}</p>
                                    )}
                                </div>

                                <div className="profile-field">
                                    <label className="profile-label">
                                        <User size={16} />
                                        Nickname
                                    </label>
                                    {isEditing ? (
                                        <input
                                            type="text"
                                            value={editedData.nickname}
                                            onChange={(e) => handleChange('nickname', e.target.value)}
                                            className="input-field"
                                        />
                                    ) : (
                                        <p className="profile-value">{userData.nickname}</p>
                                    )}
                                </div>

                                <div className="profile-field">
                                    <label className="profile-label">
                                        <Mail size={16} />
                                        Email
                                    </label>
                                    {isEditing ? (
                                        <input
                                            type="email"
                                            value={editedData.email}
                                            onChange={(e) => handleChange('email', e.target.value)}
                                            className="input-field"
                                        />
                                    ) : (
                                        <p className="profile-value">{userData.email}</p>
                                    )}
                                </div>

                                <div className="profile-field">
                                    <label className="profile-label">
                                        <Phone size={16} />
                                        Phone
                                    </label>
                                    {isEditing ? (
                                        <input
                                            type="tel"
                                            value={editedData.phone}
                                            onChange={(e) => handleChange('phone', e.target.value)}
                                            className="input-field"
                                        />
                                    ) : (
                                        <p className="profile-value">{userData.phone}</p>
                                    )}
                                </div>

                                <div className="profile-field">
                                    <label className="profile-label">
                                        <MapPin size={16} />
                                        Country ID
                                    </label>
                                    {isEditing ? (
                                        <input
                                            type="number"
                                            value={editedData.country}
                                            onChange={(e) => handleChange('country', parseInt(e.target.value) || 0)}
                                            className="input-field"
                                        />
                                    ) : (
                                        <p className="profile-value">{userData.country}</p>
                                    )}
                                </div>

                                <div className="profile-field">
                                    <label className="profile-label">
                                        <Shield size={16} />
                                        Role
                                    </label>
                                    <p className="profile-value">
                                        <span className="role-badge">{userData.role}</span>
                                    </p>
                                </div>

                                <div className="profile-field">
                                    <label className="profile-label">
                                        <Calendar size={16} />
                                        Member Since
                                    </label>
                                    <p className="profile-value">{userData.createdAt}</p>
                                </div>
                            </div>

                            <div className="profile-field mt-4">
                                <label className="profile-label">Bio</label>
                                {isEditing ? (
                                    <textarea
                                        value={editedData.bio}
                                        onChange={(e) => handleChange('bio', e.target.value)}
                                        className="input-field"
                                        rows={3}
                                    />
                                ) : (
                                    <p className="profile-value text-secondary">{userData.bio}</p>
                                )}
                            </div>
                        </div>
                    </div>
                </Card>

                {/* Preferences Card */}
                <Card title="Preferences">
                    <div className="flex flex-col gap-3">
                        {preferences.map((pref, index) => {
                            const Icon = pref.icon;
                            return (
                                <div key={index} className="preference-item">
                                    <div className="flex items-center gap-3">
                                        <div className="preference-icon">
                                            <Icon size={18} />
                                        </div>
                                        <span className="text-sm">{pref.label}</span>
                                    </div>
                                    <label className="toggle-switch">
                                        <input
                                            type="checkbox"
                                            checked={pref.enabled}
                                            readOnly
                                        />
                                        <span className="toggle-slider"></span>
                                    </label>
                                </div>
                            );
                        })}
                    </div>
                </Card>
            </div>

            {/* Recent Activity */}
            <Card title="Recent Activity">
                <div className="flex flex-col gap-3">
                    {activityData.map((activity, index) => (
                        <div key={index} className="activity-item">
                            <div className="activity-dot"></div>
                            <div className="flex-1">
                                <p className="text-sm font-medium">{activity.action}</p>
                                <p className="text-xs text-muted mt-1">{activity.time}</p>
                            </div>
                            <span className={`activity-badge activity-${activity.type}`}>
                                {activity.type}
                            </span>
                        </div>
                    ))}
                </div>
            </Card>
        </div>
    );
};
