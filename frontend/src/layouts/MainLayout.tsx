import type {ReactNode} from "react"
import {Toaster} from "sonner";

export function MainLayout({ children }: { children: ReactNode }) {
    return (
        <div className="min-h-screen flex flex-col w-full">
            <header className="w-full border-b px-4">
                <div className="container mx-auto py-4 flex items-center justify-between">
                    <h1 className="text-lg font-semibold">Urls.ac</h1>
                    <nav className="space-x-4 text-sm md:block hidden">
                        <a href="#" className="hover:underline">Log In (coming soon)</a>
                        <a href="#" className="hover:underline">Sign up (coming soon)</a>
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
