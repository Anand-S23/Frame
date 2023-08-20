'use client';

import { useRouter } from "next/navigation";
import Image from "next/image";

const Logo = () => {
    const router = useRouter();

    return (
        <div className="flex flex-row items-center space-x-4 cursor-pointer">
            <Image
                src="/images/logo.png"
                className="md:block hidden"
                width="32"
                height="32"
                alt=""
            />

            <h1 className="lg:block hidden text-2xl font-semibold">Frame</h1>
        </div>
    );
}

export default Logo;
