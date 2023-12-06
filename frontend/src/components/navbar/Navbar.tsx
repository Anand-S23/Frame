'use client';

import { Nunito } from 'next/font/google';

import Container from "../Container";
import Logo from "./Logo";
import Menu from "./Menu";
import SearchBar from "./SearchBar";

const font = Nunito({ subsets: ['latin'] });

const Navbar = () => {
    return (
        <div className="fixed w-full z-10">
            <div className="py-4 border-b-[1px]">
                <Container>
                    <div className={font.className + " flex flex-row items-center justify-between gap-3 md:gap-0"}>
                        <Logo />
                        <SearchBar />
                        <Menu />
                    </div>
                </Container>
            </div>
        </div>
    );
}

export default Navbar;
