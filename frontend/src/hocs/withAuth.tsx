import { NextPage } from 'next';
import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { isUserAuthenticated } from '@/lib/auth';

const withAuth = (WrappedComponent: NextPage) => {
    const AuthenticatedComponent: NextPage = (props) => {
        const router = useRouter();

        useEffect(() => {
            const checkAuth = async () => {
                const isAuthenticated = await isUserAuthenticated();
                console.log(isAuthenticated);

                if (!isAuthenticated) {
                    router.push('/login');
                }
            };

            checkAuth();
        }, []);

        return <WrappedComponent {...props} />;
    };

    return AuthenticatedComponent;
};

export default withAuth;
