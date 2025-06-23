import {type ReactNode} from "react"
import {Toaster} from "sonner";
import {Link} from "@tanstack/react-router";
import {Button} from "@/components/ui/button.tsx";
import { Menu } from "lucide-react"
import {Drawer, DrawerContent, DrawerTrigger} from "@/components/ui/drawer.tsx";
import {NewFooter} from "@/components/NewFooter.tsx";

export function MainLayout({ children }: { children: ReactNode }) {

    return (
        <div className="min-h-screen flex flex-col w-full">
            <header className="w-full border-b px-4">
                <div className="container mx-auto py-4 flex items-center justify-between">
                    <div className={"flex flex-row items-center space-x-4 justify-center"}>
                    <Link to={"/"}><h1 className="text-lg font-semibold">Urls.ac</h1></Link>
                        <nav className="space-x-4 text-sm md:block hidden">
                            <Link to={"/"} className="hover:underline">Shorten</Link>
                            <Link to={"/fiftyfifty"} className="hover:underline">FiftyFifty</Link>
                        </nav>
                    </div>
                    <nav className="space-x-4 text-sm md:block hidden">
                        {/*<Link to={"/login"} className="hover:underline">Log In</Link>
                        <Link to={"/signup"} className="hover:underline">Sign up</Link>*/}
                    </nav>
                    <div className="lg:hidden">
                        <Drawer direction={"right"}>
                            <DrawerTrigger asChild>
                                <Button variant="outline">
                                    <Menu className="h-5 w-5" />
                                </Button>
                            </DrawerTrigger>
                            <DrawerContent className="p-6 flex flex-col">
                                <ul className="space-y-4">
                                    <li>
                                        <Link to={"/"} className="font-semibold">Shorten</Link>
                                    </li>
                                    <li>
                                        <Link to={"/fiftyfifty"} className="font-semibold">FiftyFifty</Link>
                                    </li>
                                </ul>
                                <div className={"flex-1"}></div>

                                <ul className="space-y-4 text-muted-foreground text-sm">
                                    <li>
                                        <a href={"https://www.termsfeed.com/live/f9d3360a-2efd-4c57-8641-1869e00209c8"} className="font-semibold">Terms of Service</a>
                                    </li>
                                    <li>
                                        <a href={"https://www.termsfeed.com/live/e2c6fc22-1e69-424e-8623-b89ac0b47bdc"} className="font-semibold">Privacy Policy</a>
                                    </li>
                                </ul>
                            </DrawerContent>
                        </Drawer>
                    </div>
                </div>


            </header>

            <main className="flex-1 items-center justify-center lg:justify-start flex flex-col p-4">

                    {children}

            </main>
            <Toaster />
    <NewFooter/>
        </div>
    )
}
