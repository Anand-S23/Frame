'use client';

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { TRegisterSchema, RegisterSchema } from "@/lib/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

const RegisterForm = () => {
    const {
        register,
        handleSubmit,
        formState: { errors },
        reset,
    } = useForm<TRegisterSchema>({
        resolver: zodResolver(RegisterSchema),
    });

    const onSubmit = async (data: TRegisterSchema) => {
        // TODO: submit to server
        reset();
    };

    return (
        <div className="flex flex-col w-full max-w-md gap-1.5 p-5">
            <form onSubmit={handleSubmit(onSubmit)}>
                <Input
                    {...register("username")}
                    type="text" id="username" name="username" placeholder="Username"
                    className="mt-2 {$errors.username ? 'border-red-500' : ''}"
                />
                { errors.username && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.username.message}`}
                    </p>
                )}

                <Input 
                    {...register("email")}
                    type="email" id="email" name="email" placeholder="Email"
                    className="mt-2 {$errors.email ? 'border-red-500' : ''}"
                />
                { errors.email && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.email.message}`}
                    </p>
                )}

                <Input 
                    {...register("password")}
                    type="password" id="password" name="password" placeholder="Password"
                    className="mt-2 {$errors.password ? 'border-red-500' : ''}"
                />
                { errors.password && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.password.message}`}
                    </p>
                )}

                <Input 
                    {...register("confirm")}
                    type="password" id="confirm" name="confirm" placeholder="Confirm Password"
                    className="mt-2 {$errors.password ? 'border-red-500' : ''}"
                />
                { errors.confirm && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.confirm.message}`}
                    </p>
                )}

                <Button type="submit" className="mt-2">Submit</Button>
            </form>
        </div>
   );
} 

export default RegisterForm;
