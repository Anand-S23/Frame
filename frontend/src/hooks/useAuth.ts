'use client';

import { fetchAuthenticatedUser } from "@/lib/auth";
import { UserLoginResult } from "@/lib/types";
import { useEffect, useState } from "react";

const useAuth = () => {
    const [user, setUser] = useState<UserLoginResult | null>(null);

    useEffect(() => {
        const checkAuthentication = async () => {
            try {
                const authenticatedUser = await fetchAuthenticatedUser();
                setUser(authenticatedUser);
            } catch (error) {
                console.error('Error checking authentication: ', error);
            }
        }

        checkAuthentication();
    }, []);

    return user;
}

export default useAuth;

