import { fail } from "@sveltejs/kit"
import { z } from "zod"
import { superValidate } from "sveltekit-superforms/server"

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

export const load = async (event) => {
	const form = await superValidate(event, registerSchema);
	return { form };
}

export const actions = {
	default: async (event) => {
		const form = await superValidate(event, registerSchema);
		console.log(form);

		if (!form.valid) {
			return fail(400, { form });
		}

		return { form };
	}
}

