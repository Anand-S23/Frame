'use client';

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

// const loginEndpoint = "http://localhost:8080/login";

const LoginForm = () => {
    return (
        <div className="flex flex-col w-full max-w-md gap-1.5 p-5">
            <form>
                <Input 
                    type="email" id="email" name="email" placeholder="Email"
                />

                <Input 
                    type="password" id="password" name="password" placeholder="Password"
                />

                <Button type="submit" className="mt-2">Submit</Button>
            </form>
        </div>
    );
}

export default LoginForm;
