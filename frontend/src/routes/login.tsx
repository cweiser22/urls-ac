import { createFileRoute } from '@tanstack/react-router'
import {LoginForm} from "@/components/LoginForm.tsx";

export const Route = createFileRoute('/login')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div className={"container"}>
    <div className={"w-full lg:w-120 mx-auto"}>
    <LoginForm/>
    </div>
  </div>
}
