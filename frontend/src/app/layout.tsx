import { cn } from '@/lib/utils';
import './globals.css';
import type { Metadata } from 'next';
import { Inter } from 'next/font/google';

const font = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
    title: 'Frame',
    description: 'Social Media for movie enthusiats to share their creations',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <html lang="en">
            <body className={cn('light', font.className)}>
                {children}
            </body>
        </html>
    );
}
