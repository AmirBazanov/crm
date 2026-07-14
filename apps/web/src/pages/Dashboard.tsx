import React from 'react';
import { Card } from '../components/ui/Card';
import { Users, DollarSign, TrendingUp, Activity } from 'lucide-react';

export const Dashboard: React.FC = () => {
    const stats = [
        { label: 'Total Revenue', value: '$45,231.89', change: '+20.1%', icon: DollarSign, color: 'text-emerald-500' },
        { label: 'Active Users', value: '+2350', change: '+180.1%', icon: Users, color: 'text-blue-500' },
        { label: 'Sales', value: '+12,234', change: '+19%', icon: TrendingUp, color: 'text-orange-500' },
        { label: 'Active Now', value: '+573', change: '+201', icon: Activity, color: 'text-pink-500' },
    ];

    return (
        <div className="flex flex-col gap-6">
            <div className="flex justify-between items-center">
                <h1 className="text-2xl font-bold">Dashboard</h1>
                <div className="flex gap-2">
                    <select className="bg-secondary border border-border rounded-md px-3 py-1 text-sm outline-none focus:border-accent">
                        <option>Last 7 days</option>
                        <option>Last 30 days</option>
                        <option>Last 90 days</option>
                    </select>
                </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
                {stats.map((stat, index) => {
                    const Icon = stat.icon;
                    return (
                        <Card key={index} className="p-6">
                            <div className="flex justify-between items-start">
                                <div>
                                    <p className="text-sm font-medium text-muted">{stat.label}</p>
                                    <h3 className="text-2xl font-bold mt-2">{stat.value}</h3>
                                </div>
                                <div className={`p-2 rounded-lg bg-opacity-10 ${stat.color.replace('text-', 'bg-')}`}>
                                    <Icon className={stat.color} size={20} />
                                </div>
                            </div>
                            <div className="mt-4 flex items-center text-sm">
                                <span className="text-emerald-500 font-medium">{stat.change}</span>
                                <span className="text-muted ml-2">from last month</span>
                            </div>
                        </Card>
                    );
                })}
            </div>

            <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                <Card className="lg:col-span-2" title="Recent Activity">
                    <div className="overflow-x-auto">
                        <table className="w-full text-sm text-left">
                            <thead className="text-xs text-muted uppercase bg-secondary/50">
                                <tr>
                                    <th className="px-4 py-3 rounded-tl-lg">User</th>
                                    <th className="px-4 py-3">Action</th>
                                    <th className="px-4 py-3">Date</th>
                                    <th className="px-4 py-3 rounded-tr-lg">Status</th>
                                </tr>
                            </thead>
                            <tbody>
                                {[1, 2, 3, 4, 5].map((i) => (
                                    <tr key={i} className="border-b border-border hover:bg-secondary/30 transition-colors">
                                        <td className="px-4 py-3 font-medium">User {i}</td>
                                        <td className="px-4 py-3">Created new project</td>
                                        <td className="px-4 py-3 text-muted">Oct 2{i}, 2023</td>
                                        <td className="px-4 py-3">
                                            <span className="px-2 py-1 rounded-full text-xs font-medium bg-emerald-500/10 text-emerald-500">
                                                Completed
                                            </span>
                                        </td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                </Card>

                <Card title="Overview">
                    <div className="flex items-center justify-center h-64 text-muted">
                        Chart Placeholder
                    </div>
                </Card>
            </div>
        </div>
    );
};
