import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

export default function Register() {
    return (
        <div className="flex flex-col w-full max-w-md gap-1.5 p-5">
            <form>
                <Input
                    type="text" id="username" name="username" placeholder="Username"
                    className="mt-2 {$errors.username ? 'border-red-500' : ''}"
                />

                <Input 
                    type="email" id="email" name="email" placeholder="Email"
                    className="mt-2 {$errors.email ? 'border-red-500' : ''}"
                />

                <Input 
                    type="password" id="password" name="password" placeholder="Password"
                    className="mt-2 {$errors.password ? 'border-red-500' : ''}"
                />

                <Input 
                    type="password" id="confirm" name="confirm" placeholder="Confirm Password"
                    className="mt-2 {$errors.password ? 'border-red-500' : ''}"
                />

                <Button type="submit" className="mt-2">Submit</Button>
            </form>
        </div>
   );
}
