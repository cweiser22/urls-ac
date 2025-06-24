import { Card, CardContent } from "@/components/ui/card"

export function NewFooter() {
    return (
        <Card className="mt-16 rounded-2xl shadow-md">
            <CardContent className="flex justify-between items-center px-6 py-2 text-xs text-muted-foreground container mx-auto">
                <span>© {new Date().getFullYear()} Cooper Weiser</span>
                <div className="flex gap-4 items-center">
                        <a href="https://github.com/cweiser22/urls-ac" className="hover:underline flex flex-row items-center gap-1" >
                             Github
                        </a>

                    <a href="https://www.termsfeed.com/live/f9d3360a-2efd-4c57-8641-1869e00209c8" className="hover:underline" >Terms of Service</a>

                    <a href="https://www.termsfeed.com/live/e2c6fc22-1e69-424e-8623-b89ac0b47bdc" className="hover:underline">Privacy Policy</a>
                </div>

            </CardContent>
        </Card>
    )
}
