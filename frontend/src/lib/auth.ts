import { AUTHENTICATED_USER_ENDPOINT } from "./consts";
import { UserLoginResult } from "./types";

export const isUserAuthenticated = async () => {
    let response: Response;

    try {
        response = await fetch(AUTHENTICATED_USER_ENDPOINT, {
            credentials: 'include'
        });
    } catch (error) {
        console.error('Error fetching authenticated user:', error);
        return false;
    }

    const userData: UserLoginResult = await response.json();

    if (userData && userData.Username !== undefined && userData.User_ID !== undefined) {
        return true;
    }

    return false;
};

