import { z } from 'zod';

export const registerSchema = z
    .object({
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
        confirm: z
            .string({ required_error: 'Confirm Password is required' })
            .trim(),
    })
    .refine((data) => data.password === data.confirm, {
        message: "Passwords must match",
        path: ["confirm"],
    });

export type TRegisterSchema = z.infer<typeof registerSchema>;