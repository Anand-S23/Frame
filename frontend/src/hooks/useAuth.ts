'use client';

import { fetchAuthenticatedUser } from "@/lib/auth";
import { useEffect, useState } from "react";

const useAuth = () => {
    const [user, setUser] = useState<any>(null);

    useEffect(() => {
        const checkAuthentication = async () => {
            try {
                const authenticatedUser = await fetchAuthenticatedUser();
                setUser(authenticatedUser);
                console.log(authenticatedUser);
            } catch (error) {
                console.error('Error checking authentication: ', error);
            }
        }

        checkAuthentication();
    }, []);

    return user;
}

export default useAuth;

