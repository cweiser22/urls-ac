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
        <div className={"container mx-auto max-w-3xl px-4 py-8"}>
            <h1 className={"text-2xl mb-2 font-bold"}>True 50/50 URL Generator</h1>
            <div className={"flex flex-col "}>
                <div className={"w-full lg:max-width-xs"}>
                    <CreateFiftyFiftyLink updateResult={updateResult}/>
                    {result ? <FiftyFiftyResult urlA={result.urlA} urlB={result.urlB} shortURL={result.shortUrl} probability={result.probability}/> : <></>}
                </div>
                <div className={""}>

                    <section className="max-w-3xl mx-auto p-6 space-y-8 text-gray-800">

                        <div>
                            <h2 className="text-2xl font-bold mb-2">What Is a 50/50 Link?</h2>
                            <p className="text-base leading-relaxed">
                                A <strong>FiftyFifty link</strong> is a unique, shareable URL that randomly redirects to one of two links you choose.
                                When someone clicks your custom link, there's a <strong>50% chance they'll be sent to URL A</strong> and a <strong>50% chance they'll land on URL B</strong>.
                                It's a fun and surprising way to share content — perfect for games, pranks, A/B testing, or just adding a bit of mystery to your links.
                            </p>
                        </div>


                        <div>
                            <h2 className="text-2xl font-bold mb-2">How to Use the FiftyFifty Link Generator</h2>
                            <ol className="list-decimal pl-5 space-y-3 text-base leading-relaxed">
                                <li>
                                    <strong>Enter Two URLs</strong><br/>
                                    Start by entering any two valid URLs — they can be websites, videos, blog posts, memes, or any links you'd like to randomly redirect between.
                                </li>
                                <li>
                                    <strong>Click “Generate”</strong><br/>
                                    After entering both links, hit the <span className="font-medium">Generate</span> button. This will create a unique <span className="font-medium">shortcode link</span> that you can copy and share with anyone.
                                </li>
                                <li>
                                    <strong>Share Your Link</strong><br/>
                                    When someone clicks your new link, they'll be randomly sent to one of the two destinations.
                                </li>
                            </ol>
                        </div>


                        <div>
                            <h2 className="text-2xl font-bold mb-2">Customize the Odds (Optional)</h2>
                            <p className="text-base leading-relaxed">
                                If you don't want simple 50/50 odds, you can customize the chances of getting either link by using the <strong>"Customize Probability"</strong> feature.
                            </p>
                            <ul className="list-disc pl-5 space-y-2 text-base leading-relaxed">
                                <li>
                                    To access it, <strong>toggle on the "Customize Probability" switch</strong> just below the URL inputs.
                                </li>
                                <li>
                                    A <strong>slider will appear</strong>, allowing you to set the odds however you'd like — for example, <em>70% chance of URL A and 30% chance of URL B</em>, or any other combination between 1% and 99%.
                                </li>
                                <li>
                                    Once set, the link you generate will follow your custom odds instead of an even split.
                                </li>
                            </ul>

                        </div>
                    </section>

                </div>

            </div>
        </div>
    )
}