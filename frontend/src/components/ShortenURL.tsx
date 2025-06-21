import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {Input} from "@/components/ui/input.tsx";
import {Button} from "@/components/ui/button.tsx";
import {toast} from "sonner";
import {useState} from "react";

const apiBase = import.meta.env.VITE_API_HOST || 'http://localhost:8080';

interface Props{
    updateResult: (longUrl: string, shortUrl: string) => void;
}

export function ShortenURL({updateResult}: Props) {
    const [inputValue, setInputValue] = useState("");

    const handleClick = async () => {
        // Logic to shorten the URL will go here
        try {
            toast.success('Successfully shortened URL!');
            const response = await fetch(`${apiBase}/api/v1/mappings`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({longUrl: inputValue}),
                }
            );

            if (!response.ok) {
                if (response.status === 400) {
                    toast.error("Invalid URL format", {
                        description: "Please enter a valid URL.",
                    });
                } else {
                    toast.error("Failed to shorten URL. Please try again later.", {
                        description: `Error: ${response.statusText}`,
                    });
                }
                return;
            }

            const {shortUrl} = await response.json();
            updateResult(inputValue, shortUrl);
            setInputValue("");

        } catch (error) {
            toast.error('Failed to shorten URL. Please try again later.');
            console.error('Error shortening URL:', error);
        }

    }
    return (<Card className={"w-full max-w-xl mx-auto"}>
        <CardHeader>
            <CardTitle>urls.ac</CardTitle>
            <CardDescription>Shorten a URL</CardDescription>
        </CardHeader>
        <CardContent>
            <div className="flex w-full max-w-sm items-center gap-2">
                <Input onChange={(e) => setInputValue(e.target.value)} value={inputValue} type="text" placeholder="Long URL" />
                <Button  onClick={handleClick} type="submit" variant="outline" className={"bg-sky-500 text-white "}>
                    Shorten
                </Button>
            </div>
        </CardContent>

    </Card>);
}