'use client';

import { Home } from 'lucide-react';
import { MessageCircle } from 'lucide-react';
import ProfileAvatar from './ProfileAvatar';

const Menu = () => {
    return (
        <div className="flex flex-row items-center cursor-pointer">
            <div className="flex flex-row items-center px-2">
                <Home className='px-2 h-10 w-10 hover:text-gray-400'/>
                <MessageCircle className='px-2 h-10 w-10 hover:text-gray-400'/>
            </div>

            <ProfileAvatar />
        </div>
    );
}

export default Menu;
