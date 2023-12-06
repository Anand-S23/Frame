'use client';

import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Search } from 'lucide-react';

const SearchBar = () => {
    return (
        <div className="w-3/4 lg:w-1/2 justify-center">
            <div className="flex w-full items-center space-x-2">
                <Input placeholder="Search" />
                <Button type="submit" className="bg-rose-500 p-2 h-full aspect-square">
                    <Search className="h-4 w-4 text-white"/>
                </Button>
            </div>
        </div>
    );
}

export default SearchBar;
