'use client';

import Navbar from "@/components/navbar/Navbar";
import withAuth from "@/hocs/withAuth";

function Home() {
    return (
        <>
            <Navbar />
        </>
    );
}

export default withAuth(Home);
