'use client';

import Navbar from "@/components/navbar/Navbar";
import useAuth from "@/hooks/useAuth";

export default function Home() {
    let user = useAuth();
    console.log(user);

    return (
        <>
            <Navbar />
        </>
    );
}
