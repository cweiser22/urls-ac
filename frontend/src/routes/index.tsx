import { createFileRoute } from '@tanstack/react-router'
import {ShortenURL} from "@/components/ShortenURL.tsx";
import {ShortenLink} from "@/components/Result.tsx";
import {useState} from "react";
import {toast} from "sonner";

export const Route = createFileRoute('/')({
    component: Index,
})

function Index() {
    const [result, setResult] = useState<Result | null>(null);

    const updateResult = (longUrl: string, shortUrl: string) => {
        setResult({ longUrl, shortUrl });
        toast.success("Successfully shortened URL!");
    }
    
    return (
        <div className={"container"}>
            <div className={"mx-auto w-full"}>
                <ShortenURL updateResult={updateResult}/>
                {result && <ShortenLink longURL={result.longUrl} shortURL={result.shortUrl} />}
            </div>

        </div>
    )
}