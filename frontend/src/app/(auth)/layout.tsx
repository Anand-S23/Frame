import type { Metadata } from 'next';

export const metadata: Metadata = {
    title: 'Frame',
    description: 'Social Media for movie enthusiats to share their creations',
};

export default function AuthLayout({ children }: { children: React.ReactNode }) {
    return (
        <>
            {children}
        </>
    );
}
