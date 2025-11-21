import React from 'react';
import { Outlet } from 'react-router-dom';

export const AuthLayout: React.FC = () => {
    return (
        <div className="flex-center" style={{ minHeight: '100vh', backgroundColor: 'var(--bg-primary)' }}>
            <div style={{ width: '100%', maxWidth: '400px', padding: '1rem' }}>
                <Outlet />
            </div>
        </div>
    );
};
