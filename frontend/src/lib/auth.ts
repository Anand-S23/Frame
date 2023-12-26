import { AUTHENTICATED_USER_ENDPOINT } from "./consts";
import { UserLoginResult } from "./types";

export const fetchAuthenticatedUser = async () => {
    try {
        const response = await fetch(AUTHENTICATED_USER_ENDPOINT, {
            credentials: 'include'
        });

        const data = await response.json();
        return data as UserLoginResult;

    } catch (error) {
        console.error('Error fetching authenticated user:', error);
        throw error;
    }
};
