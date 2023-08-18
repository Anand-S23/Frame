'use client';

import Container from "../Container";
import Logo from "./Logo";
import ProfileMenu from "./ProfileMenu";
import Search from "./Search";

const Navbar = () => {
    return (
        <div className="fixed w-full z-10">
            <div className="py-4 border-b-[1px]">
                <Container>
                    <div className="flex flex-row items-center justify-between gap-3 md:gap-0">
                        <Logo />
                        <Search />
                        <ProfileMenu />
                    </div>
                </Container>
            </div>
        </div>
    );
}

export default Navbar;
