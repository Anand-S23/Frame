import { AUTHENTICATED_USER_ENDPOINT } from "./consts";

export const fetchAuthenticatedUser = async () => {
    try {
        const response = await fetch(AUTHENTICATED_USER_ENDPOINT, {
            credentials: 'include'
        });

        console.log(response);

        const data = await response.json();
        console.log(data);

    } catch (error) {
        console.error('Error fetching authenticated user:', error);
        throw error;
    }
};
