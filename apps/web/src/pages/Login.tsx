import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { Card } from '../components/ui/Card';
import { Input } from '../components/ui/Input';
import { Button } from '../components/ui/Button';
import { LogIn } from 'lucide-react';
import { authService } from '../api/authService';

export const Login: React.FC = () => {
    const navigate = useNavigate();
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState('');
    const [formData, setFormData] = useState({
        email: '',
        password: '',
    });

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsLoading(true);
        setError('');

        try {
            const response = await authService.login({
                email: formData.email,
                password: formData.password,
            });

            // Store tokens and user data
            authService.setTokens(response.accessToken, response.refreshToken);
            authService.setUser(response.user);

            // Navigate to dashboard
            navigate('/');
        } catch (err: any) {
            setError(err.message || 'Login failed. Please try again.');
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="animate-fade-in">
            <div className="text-center mb-8">
                <h1 className="text-3xl font-bold mb-2">Welcome Back</h1>
                <p className="text-muted">Sign in to your account to continue</p>
            </div>

            <Card>
                {error && (
                    <div className="error-banner mb-4">
                        {error}
                    </div>
                )}
                <form onSubmit={handleSubmit} className="flex flex-col gap-4">
                    <Input
                        label="Email Address"
                        type="email"
                        placeholder="name@company.com"
                        value={formData.email}
                        onChange={(e) => setFormData({ ...formData, email: e.target.value })}
                        required
                    />

                    <Input
                        label="Password"
                        type="password"
                        placeholder="••••••••"
                        value={formData.password}
                        onChange={(e) => setFormData({ ...formData, password: e.target.value })}
                        required
                    />

                    <div className="flex justify-between items-center text-sm">
                        <label className="flex items-center gap-2 cursor-pointer">
                            <input type="checkbox" className="rounded border-gray-300" />
                            <span className="text-secondary">Remember me</span>
                        </label>
                        <Link to="/forgot-password" className="text-accent hover:underline">
                            Forgot password?
                        </Link>
                    </div>

                    <Button type="submit" className="w-full mt-2" isLoading={isLoading} icon={<LogIn size={18} />}>
                        Sign In
                    </Button>
                </form>
            </Card>

            <p className="text-center mt-6 text-sm text-muted">
                Don't have an account?{' '}
                <Link to="/register" className="text-accent hover:underline font-medium">
                    Sign up
                </Link>
            </p>
        </div>
    );
};
