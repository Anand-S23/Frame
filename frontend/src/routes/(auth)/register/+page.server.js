import { z } from 'zod';

const registerSchema = z.object({
    name: z
        .string({ required_error: 'Name is required' })
        .min(5, { message: "Username must be between 5-15 characters long" })
        .max(15, { message: "Username must be between 5-15 characters long" })
        .trim(),
    email: z
        .string({ required_error: 'Email is required' })
        .min(1, { message: 'Email is required' })
        .max(64, { message: 'Email must be less than 64 characters' })
        .email({ message: 'Email must be a valid email address' }),
    password: z
        .string({ required_error: 'Password is required' })
        .min(8, { message: 'Password must be at least 8 characters' })
        .max(32, { message: 'Password must be less than 32 characters' })
        .trim(),
});

export const actions = {
	default: async ({ request }) => {
		const formData = Object.fromEntries(await request.formData());

		try {
			const result = registerSchema.parse(formData);
			console.log(result);
		} catch (err) {
			const { fieldErrors: errors } = err.flatten();
			const { password, ...rest } = formData;
			return {
				data: rest,
				errors
			};
		}
	}
};

