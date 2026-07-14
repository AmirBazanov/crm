import React from 'react';
import { clsx } from 'clsx';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    variant?: 'primary' | 'secondary' | 'ghost' | 'danger';
    size?: 'sm' | 'md' | 'lg';
    isLoading?: boolean;
    icon?: React.ReactNode;
}

export const Button: React.FC<ButtonProps> = ({
    className,
    variant = 'primary',
    size = 'md',
    isLoading,
    icon,
    children,
    disabled,
    ...props
}) => {
    const variantClasses = {
        primary: 'btn-primary',
        secondary: 'btn-secondary',
        ghost: 'btn-ghost',
        danger: 'btn-danger',
    };

    const sizeClasses = {
        sm: 'btn-sm',
        md: 'btn-md',
        lg: 'btn-lg',
    };

    return (
        <button
            className={clsx('btn', variantClasses[variant], sizeClasses[size], className)}
            disabled={disabled || isLoading}
            {...props}
        >
            {isLoading && <span className="spinner mr-2" />}
            {!isLoading && icon && <span className="icon mr-2">{icon}</span>}
            {children}
        </button>
    );
};
