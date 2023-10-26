import { fail, redirect } from "@sveltejs/kit"
import { z } from "zod"
import { superValidate } from "sveltekit-superforms/server"
import axios from "axios"

const loginEndpoint = "http://localhost:8080/login"

const loginSchema = z.object({
    email: z.string({ required_error: 'Email is required' }),
    password: z.string({ required_error: 'Password is required' }).trim()
});

export const load = async (event) => {
    // TODO: Maybe check if they are already logged in here?
	const form = await superValidate(event, loginSchema);
	return { form };
}

export const actions = {
	default: async (event) => {
		let form = await superValidate(event, loginSchema);

		if (!form.valid) {
			return fail(400, { form });
		}

        await axios.post(loginEndpoint, form.data)
            .catch((error) => {
                if (error.response) {
                    let data = error.response.data;
                    console.log(data)
                }
            });

        throw redirect(302, "/");
	}
}

