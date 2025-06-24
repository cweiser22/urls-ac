import {createFileRoute, Link} from '@tanstack/react-router'
import {ShortenURL} from "@/components/ShortenURL.tsx";
import {ShortenURLResult} from "@/components/ShortenURLResult.tsx";
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
            <h1 className={"text-2xl mb-2 font-bold text-center"}>urls.ac - A Modern URL Shortener</h1>
            <div className={"flex flex-col gap-4 mx-auto w-full max-w-lg"}>
                <div className={""}>
                    <ShortenURL updateResult={updateResult}/>
                    {result && <ShortenURLResult longURL={result.longUrl} shortURL={result.shortUrl} />}
                </div>
                <div className={"col-span-2 space-y-2"}>
                    <h1 className={"text-2xl mb-2 font-bold"}>A Fun and Simple URL Shortening Service</h1>
                    <p className={"text-md"}>Create a good old-fashioned short URL with our generator. Feel free to shorten
                    video links, news articles, academic papers, social media posts, or anything else you want. Simply punch in
                    the URL you want to shorten, click Shorten, and get your very own urls.ac short link. URL shortening
                        is a great way to make long URLs more manageable and shareable, especially on platforms with character limits.</p>


                </div>
                <div className={"col-span-2 space-y-2"}>
                    <h1 className={"text-2xl mb-2 font-bold"}>Other utilities</h1>
                    <p className={"text-md"}>Feel free to have fun with some of our other URL services: </p>
                    <ul className={"list-disc pl-4"}>
                        <li><Link to={"/fiftyfifty"} className={"text-sky-600 underline"}>Fifty Fifty Link Generator</Link>
                            - Enter two URLs and toss a coin. A fifty-fifty URL is a special link that points to two different long URLs.
                        When a user clicks a fifty-fifty link, they will be redirected to one of these two URLs, which will be determined randomly each time they click.
                        However, you are not restricted to 50/50 odds. Our generator allows you to customize the odds of receiving either of the two links. For example, you can opt to have a 70% chance of the first URL being chosen, and a 30% chance of the second URL being chosen.</li>
                    </ul>
                </div>
                <div className={"col-span-2 space-y-2"}>
                    <h1 className={"text-2xl mb-2 font-bold"}>Open-Source</h1>
                    <p className={"text-md"}>Urls.ac is an open-source, MIT-licensed platform created by Cooper Weiser. The source code for urls.ac can be located on <a href={"https://github.com/cweiser22/urls-ac"} className={"underline text-sky-600"}>Github</a>.</p>

                </div>
            </div>
        </div>
    )
}