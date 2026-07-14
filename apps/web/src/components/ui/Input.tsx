import React from 'react';
import { clsx } from 'clsx';

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
    label?: string;
    error?: string;
    fullWidth?: boolean;
}

export const Input: React.FC<InputProps> = ({
    className,
    label,
    error,
    fullWidth = true,
    id,
    ...props
}) => {
    const inputId = id || React.useId();

    return (
        <div className={clsx('input-group', fullWidth && 'w-full', className)}>
            {label && (
                <label htmlFor={inputId} className="input-label">
                    {label}
                </label>
            )}
            <input
                id={inputId}
                className={clsx('input-field', error && 'input-error')}
                {...props}
            />
            {error && <span className="input-error-msg">{error}</span>}
        </div>
    );
};
