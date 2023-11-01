import type { Actions, RequestEvent } from '@sveltejs/kit';
import { fail, redirect, error } from "@sveltejs/kit";
import { z } from "zod";
import { superValidate, setError } from "sveltekit-superforms/server";
import axios from "axios";

const loginEndpoint = "http://localhost:8080/login";

const loginSchema = z.object({
    email: z.string({ required_error: 'Email is required' }),
    password: z.string({ required_error: 'Password is required' }).trim()
});

export const load = async (event: RequestEvent) => {
    // TODO: Check if they are already logged in here
	const form = await superValidate(event, loginSchema);
	return { form };
}

export const actions: Actions = {
	login: async (event: RequestEvent) => {
		let form = await superValidate(event, loginSchema);

		if (!form.valid) {
			return fail(400, { form });
		}

        let errorMsg: string = '';
        await axios.post(loginEndpoint, form.data)
            .catch((error) => {
                if (error.response) {
                    let data = error.response.data;
                    errorMsg = data.error ?? '';
                }
            });

        if (errorMsg !== '') {
            throw error(400, errorMsg);
        }

        throw redirect(302, "/");
	}
}

