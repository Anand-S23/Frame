import './globals.css';
import type { Metadata } from 'next';
import { Inter } from 'next/font/google';

import Navbar from '@/components/navbar/Navbar';

const font = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
    title: 'Capsule',
    description: 'Social Media for movie enthusiats to share their creations',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <html lang="en">
            <body className={font.className + ' dark'}>
                <Navbar />
                {children}
            </body>
        </html>
    );
}
