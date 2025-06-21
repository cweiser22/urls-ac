import {createFileRoute} from '@tanstack/react-router'
import {useState} from "react";
import {toast} from "sonner";
import {CreateFiftyFiftyLink} from "@/components/CreateFiftyFiftyLink.tsx";
import {FiftyFiftyResult} from "@/components/FiftyFiftyResult.tsx";

export const Route = createFileRoute('/fiftyfifty')({
    component: Index,
})

type Result = {
    urlA: string;
    urlB: string;
    probability: number;
    shortUrl: string;
}

function Index() {
    const [result, setResult] = useState<Result | null>(null);

    const updateResult = (urlA: string, urlB: string, shortUrl: string, probability: number) => {
        setResult({ probability, urlB, urlA, shortUrl });
        toast.success("Successfully shortened URL!");
    }

    return (
        <div className={"container "}>
            <div className={"flex flex-col"}>

                <div className={""}>
                    <h1 className={"text-2xl mb-2 font-bold"}>Fifty Fifty URL Generator</h1>
                    <p className={"text-md"}>Enter any two URLs, and a special shortcode link will be generated that has a 50% chance of being URL 1 and a 50% of being URL B.</p>
                    <p className={"text-md"}>If a 50/50 coin toss isn't for you though, feel free to use the slider to customize the odds.</p>
                </div>
                <div className={"w-full lg:max-width-xs"}>
                    <CreateFiftyFiftyLink updateResult={updateResult}/>
                    {result ? <FiftyFiftyResult urlA={result.urlA} urlB={result.urlB} shortURL={result.shortUrl} probability={result.probability}/> : <></>}
                </div>
            </div>
        </div>
    )
}