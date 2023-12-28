'use client';

import { NextPage } from 'next';
import { useLayoutEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { isUserAuthenticated } from '@/lib/auth';
import { Spinner } from '@/components/ui/spinner';

const withoutAuth = (WrappedComponent: NextPage) => {
    const AuthenticatedComponent: NextPage = (props) => {
        const router = useRouter();
        const [isLoading, setIsLoading] = useState(true);

        useLayoutEffect(() => {
            const checkAuth = async () => {
                const isAuthenticated = await isUserAuthenticated();
                if (isAuthenticated) {
                    router.push('/');
                } else {
                    setIsLoading(false);
                }
            };

            checkAuth();
        }, []);

        if (isLoading) {
            return <Spinner />;
        }

        return <WrappedComponent {...props} />;
    };

    return AuthenticatedComponent;
};

export default withoutAuth;
