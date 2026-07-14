import React from 'react';
import { clsx } from 'clsx';

interface CardProps extends React.HTMLAttributes<HTMLDivElement> {
    title?: string;
    action?: React.ReactNode;
}

export const Card: React.FC<CardProps> = ({
    className,
    title,
    action,
    children,
    ...props
}) => {
    return (
        <div className={clsx('card glass-panel', className)} {...props}>
            {(title || action) && (
                <div className="card-header">
                    {title && <h3 className="card-title">{title}</h3>}
                    {action && <div className="card-action">{action}</div>}
                </div>
            )}
            <div className="card-content">{children}</div>
        </div>
    );
};
