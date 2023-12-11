'use client';

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { LoginSchema, TLoginSchema } from "@/lib/types";
import { cn } from "@/lib/utils";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

// const loginEndpoint = "http://localhost:8080/login";

const LoginForm = () => {
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<TLoginSchema>({
        resolver: zodResolver(LoginSchema),
    });

    const onSubmit = async (data: TLoginSchema) => {
        // TODO: Handle login submit
        console.log(data);
    }

    return (
        <div className="flex flex-col w-full max-w-md gap-1.5 p-5">
            <form onSubmit={handleSubmit(onSubmit)}>
                <Input 
                    {...register("email")}
                    type="email" id="email" name="email" placeholder="Email"
                    className={cn("mt-2", errors.email ? 'border-red-500' : '')}
                />
                { errors.email && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.email.message}`}
                    </p>
                )}

                <Input 
                    {...register("password")}
                    type="password" id="password" name="password" placeholder="Password"
                    className={cn("mt-2", errors.password ? 'border-red-500' : '')}
                />
                { errors.password && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.password.message}`}
                    </p>
                )}

                <Button type="submit" className="mt-2">Submit</Button>
            </form>
        </div>
    );
}

export default LoginForm;
