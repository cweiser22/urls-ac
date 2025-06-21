import { createFileRoute } from '@tanstack/react-router'
import {SignupForm} from "@/components/SignupForm.tsx";

export const Route = createFileRoute('/signup')({
    component: RouteComponent,
})

function RouteComponent() {
    return <div className={"container"}>
        <div className={"w-full lg:w-120 mx-auto"}>
            <SignupForm/>
        </div>
    </div>
}
