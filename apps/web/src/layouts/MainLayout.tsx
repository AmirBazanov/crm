import React from 'react';
import { Outlet, Link, useLocation } from 'react-router-dom';
import { LayoutDashboard, Users, Settings, LogOut, Bell, Search } from 'lucide-react';
import { clsx } from 'clsx';
import { userService } from '../api/userService';

export const MainLayout: React.FC = () => {
    const location = useLocation();

    const navItems = [
        { icon: LayoutDashboard, label: 'Dashboard', path: '/' },
        { icon: Users, label: 'Contacts', path: '/contacts' },
        { icon: Settings, label: 'Settings', path: '/settings' },
    ];

    const userInitials = userService.getUserInitials();

    return (
        <div className="app-layout">
            {/* Sidebar */}
            <aside className="sidebar">
                <div className="sidebar-header">
                    <div className="logo">CRM System</div>
                </div>

                <nav className="sidebar-nav">
                    {navItems.map((item) => {
                        const Icon = item.icon;
                        const isActive = location.pathname === item.path;

                        return (
                            <Link
                                key={item.path}
                                to={item.path}
                                className={clsx('nav-item', isActive && 'active')}
                            >
                                <Icon size={20} />
                                <span>{item.label}</span>
                            </Link>
                        );
                    })}
                </nav>

                <div className="sidebar-footer">
                    <button className="nav-item w-full">
                        <LogOut size={20} />
                        <span>Logout</span>
                    </button>
                </div>
            </aside>

            {/* Main Content */}
            <main className="main-content">
                {/* Header */}
                <header className="header glass-panel">
                    <div className="search-bar">
                        <Search size={18} className="text-muted" />
                        <input type="text" placeholder="Search..." className="bg-transparent border-none outline-none text-sm ml-2 w-full text-white" />
                    </div>

                    <div className="header-actions">
                        <button className="icon-btn">
                            <Bell size={20} />
                            <span className="badge">2</span>
                        </button>
                        <Link to="/profile" className="user-avatar">
                            <div className="avatar-circle">{userInitials}</div>
                        </Link>
                    </div>
                </header>

                {/* Page Content */}
                <div className="page-container">
                    <Outlet />
                </div>
            </main>
        </div>
    );
};
