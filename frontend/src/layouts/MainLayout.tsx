import type {ReactNode} from "react"
import {Toaster} from "sonner";
import {Link} from "@tanstack/react-router";

export function MainLayout({ children }: { children: ReactNode }) {
    return (
        <div className="min-h-screen flex flex-col w-full">
            <header className="w-full border-b px-4">
                <div className="container mx-auto py-4 flex items-center justify-between">
                    <Link to={"/"}><h1 className="text-lg font-semibold">Urls.ac</h1></Link>
                    <nav className="space-x-4 text-sm md:block hidden">
                        <Link to={"/login"} className="hover:underline">Log In</Link>
                        <Link to={"/signup"} className="hover:underline">Sign up</Link>
                    </nav>
                </div>
            </header>

            <main className="flex-1 items-center justify-center flex flex-col p-4">

                    {children}

            </main>
            <Toaster />
        </div>
    )
}
