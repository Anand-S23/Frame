import { NextRequest, NextResponse } from "next/server";
import { isUserAuthenticated } from "./lib/auth";

const openPages = ['/login', '/register'];

export default async function middleware(req: NextRequest) {
    if (req.nextUrl.pathname.startsWith("/_next/")) {
        return NextResponse.next();
    }

    let isAuthenticated = await isUserAuthenticated();

    console.log(isAuthenticated, req.nextUrl.pathname);

    if (isAuthenticated && openPages.includes(req.nextUrl.pathname)) {
        return NextResponse.redirect(new URL('/', req.url));
    }

    if (!isAuthenticated && !openPages.includes(req.nextUrl.pathname)) {
        return NextResponse.redirect(new URL('/login', req.url));
    }

    return NextResponse.next();
}
