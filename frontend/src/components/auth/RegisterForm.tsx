'use client';

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { TRegisterSchema, RegisterSchema } from "@/lib/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import { cn } from "@/lib/utils";

const registerEndpoint = "http://localhost:8080/register";

const RegisterForm = () => {
    const router = useRouter();

    const {
        register,
        handleSubmit,
        formState: { errors },
        setError
    } = useForm<TRegisterSchema>({
        resolver: zodResolver(RegisterSchema),
    });

    const onSubmit = async (data: TRegisterSchema) => {
         const response = await fetch(registerEndpoint, {
            method: "POST",
            mode: "cors",
            body: JSON.stringify(data),
            headers: { "Content-Type": "application/json" },
        });

        if (!response.ok) {
            // TODO: Error handling
            return;
        }

        const responseData = await response.json();
        if (responseData.errors) {
            const errors = responseData.errors;

            if (errors.username) {
                setError("username", {
                    type: "server", 
                    message: errors.username
                })
            }

            if (errors.email) {
                setError("email", {
                    type: "server", 
                    message: errors.email
                })
            }

            if (errors.password) {
                setError("password", {
                    type: "server", 
                    message: errors.password
                })
            }
        }

        // TODO: potentially go striaght to login
        // Redirect to homepage
        router.push("/");
    };

    return (
        <div className="flex flex-col w-full max-w-md gap-1.5 p-5">
            <form onSubmit={handleSubmit(onSubmit)}>
                <Input
                    {...register("username")}
                    type="text" id="username" name="username" placeholder="Username"
                    className={cn("mt-2", errors.username ? 'border-red-500' : '')}
                />
                { errors.username && (
                    <p className="text-sm text-red-500 mx-2">
                        {`${errors.username.message}`}
                    </p>
                )}

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

                <Input 
                    {...register("confirm")}
                    type="password" id="confirm" name="confirm" placeholder="Confirm Password"
                    className={cn("mt-2", errors.confirm ? 'border-red-500' : '')}
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
