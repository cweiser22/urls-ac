import { useState } from "react";
import { Card, CardContent } from "@/components/ui/card.tsx";
import { Button } from "@/components/ui/button.tsx";
import {toast} from "sonner";

interface ShortenLinkProps {
    urlA: string;
    urlB: string;
    shortURL: string;
    probability: number;
}

export function FiftyFiftyResult({ urlA, urlB, shortURL, probability }: ShortenLinkProps) {
    const [copied, setCopied] = useState(false);

    const handleCopy = async () => {
        await navigator.clipboard.writeText(shortURL);
        setCopied(true);
        toast.success("Copied to clipboard");
        setTimeout(() => setCopied(false), 1500); // Reset after 1.5s
    };

    return (
        <Card className="bg-indigo-500 text-white w-full  mx-auto mt-6 shadow-lg">
            <CardContent className="flex justify-between items-center px-6 py-4">
                {/* URL Column */}
                <div className="flex flex-col gap-1 overflow-hidden">
                    <a
                        href={shortURL}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="font-medium break-all no-underline leading-tight"
                    >
                        {shortURL}
                    </a>
                    <a
                        href={urlA}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="block break-all text-white/80 no-underline text-xs"
                    >
                        {Math.round(probability * 100)}% chance of {urlA}
                    </a>
                    <a
                        href={urlB}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="block break-all text-white/80 no-underline text-xs"
                    >
                        {100 - Math.round(probability * 100)}% chance of {urlB}
                    </a>
                </div>

                {/* Copy Button */}
                <Button
                    type="button"
                    variant="secondary"
                    className="text-xs py-1.5 ml-4 shrink-0"
                    onClick={handleCopy}
                >
                    {copied ? "Copied!" : "Copy"}
                </Button>
            </CardContent>
        </Card>
    );
}
