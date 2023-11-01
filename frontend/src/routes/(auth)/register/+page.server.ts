import type { Actions, RequestEvent } from './$types';
import { fail, redirect } from "@sveltejs/kit"
import { z } from "zod"
import { superValidate } from "sveltekit-superforms/server"
import axios from "axios"

const registerEndpoint = "http://localhost:8080/register"

const registerSchema = z.object({
    username: z
        .string({ required_error: 'Username is required' })
        .min(5, { message: "Username must be at least 5 characters" })
        .max(25, { message: "Username must be less than 25" })
        .trim(),
    email: z
        .string({ required_error: 'Email is required' })
        .max(64, { message: 'Email must be less than 64 characters' })
        .email({ message: 'Email must be a valid email address' }),
    password: z
        .string({ required_error: 'Password is required' })
        .min(8, { message: 'Password must be at least 8 characters' })
        .max(32, { message: 'Password must be less than 32 characters' })
        .trim(),
});

export const load = async (event: RequestEvent) => {
    // TODO: Check if logged in here
	const form = await superValidate(event, registerSchema);
	return { form };
}

export const actions: Actions = {
	register: async (event: RequestEvent) => {
		let form = await superValidate(event, registerSchema);

		if (!form.valid) {
			return fail(400, { form });
		}

        await axios.post(registerEndpoint, form.data)
            .catch((error) => {
                if (error.response) {
                    let data = error.response.data;
                    
                    if (data.username) {
                        form.errors.username = new Array();
                        form.errors.username.push(data.username);
                    }

                    if (data.email) {
                        form.errors.email = new Array();
                        form.errors.email.push(data.email);
                    }

                    if (data.password) {
                        form.errors.password = new Array();
                        form.errors.password.push(data.password);
                    }
                }
            });

        if (form.errors.username || form.errors.email || form.errors.password) {
            return fail(400, { form });
        } 

        throw redirect(302, "/");
	}
}

