'use client';

import { useRouter } from "next/navigation";

const Logo = () => {
    const router = useRouter();

    return (
        <div className="flex flex-row items-center space-x-2 cursor-pointer">

            <button className="bg-rose-500 hover:bg-rose-500 w-10 h-10 md:block hidden"></button>

            <h1 className="lg:block hidden text-2xl font-semibold">Frame</h1>
        </div>
    );
}

export default Logo;
