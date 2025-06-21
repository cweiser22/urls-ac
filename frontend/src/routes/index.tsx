import {createFileRoute, Link} from '@tanstack/react-router'
import {ShortenURL} from "@/components/ShortenURL.tsx";
import {ShortenLink} from "@/components/Result.tsx";
import {useState} from "react";
import {toast} from "sonner";

export const Route = createFileRoute('/')({
    component: Index,
})

type Result = {
    longUrl: string;
    shortUrl: string;
}

function Index() {
    const [result, setResult] = useState<Result | null>(null);

    const updateResult = (longUrl: string, shortUrl: string) => {
        setResult({ longUrl, shortUrl });
        toast.success("Successfully shortened URL!");
    }

    return (
        <div className={"container "}>
            <div className={"grid grid-cols-1 lg:grid-cols-3 gap-4 mx-auto w-full"}>
                <div className={""}>
                    <ShortenURL updateResult={updateResult}/>
                    {result && <ShortenLink longURL={result.longUrl} shortURL={result.shortUrl} />}
                </div>
                <div className={"col-span-2 space-y-2"}>
                    <h1 className={"text-2xl mb-2 font-bold"}>urls.ac - A Modern URL Shortener</h1>
                    <p className={"text-md"}>Create a good old-fashioned short URL with our generator, or have fun with some of our other utilities: </p>
                    <ul className={"list-disc pl-4"}>
                        <li><Link to={"/fiftyfifty"} className={"text-sky-600 underline"}>Fifty Fifty Link Generator</Link> - Enter two URLs and toss a coin</li>
                    </ul>
                </div>

            </div>
        </div>
    )
}