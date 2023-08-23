import './globals.css';
import type { Metadata } from 'next';
import { Inter } from 'next/font/google';

import Navbar from '@/components/navbar/Navbar';
import Provider from './_trpc/Provider';

const font = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
    title: 'Frame',
    description: 'Social Media for movie enthusiats to share their creations',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <html lang="en">
            <body className={font.className + ' dark'}>
                <Navbar />
                <Provider> { children } </Provider>
            </body>
        </html>
    );
}
