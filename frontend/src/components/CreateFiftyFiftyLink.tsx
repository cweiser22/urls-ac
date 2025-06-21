import {useState} from "react";
import {Card, CardContent, CardDescription, CardHeader, CardTitle} from "@/components/ui/card.tsx";
import {Input} from "@/components/ui/input.tsx";
import {Button} from "@/components/ui/button.tsx";
import {Switch} from "@/components/ui/switch"
import {MySlider} from "@/components/ui/my-slider.tsx";
import {Label} from "@/components/ui/label.tsx";
import {toast} from "sonner";

interface Props {
    updateResult: (urlA: string, urlB: string, shortUrl: string, probability: number) => void;
}

const apiBase = import.meta.env.VITE_API_HOST || 'http://localhost:8080';

export function CreateFiftyFiftyLink({updateResult}: Props) {
    const [url1Value, setUrl1Value] = useState("");
    const [url2Value, setUrl2Value] = useState("");
    const [probability, setProbability] = useState([.5]);
    const [customize, setCustomize] = useState(false);

    const handleUrl1Change = (e: React.ChangeEvent<HTMLInputElement>) => {
        setUrl1Value(e.target.value);
    }

    const handleUrl2Change = (e: React.ChangeEvent<HTMLInputElement>) => {
        setUrl2Value(e.target.value);
    }

    const handleToggleCustomize = () => {
        setCustomize(!customize);
        if (!customize) {
            setProbability([.5]); // Reset to default probability when toggling off
        } else {
            setProbability([.5]); // Set to 0 when toggling on, allowing customization
        }
    }

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        // Logic to create the fifty-fifty link will go here
        try {
            const response = await fetch(`${apiBase}/api/v1/ff/`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    urlA: url1Value,
                    urlB: url2Value,
                    probability: probability[0],
                }),
            });

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
                return
            }

            const data = await response.json();
            // Handle success, e.g., show a success message or redirect
            console.log('Fifty-Fifty Link Created:', data);
            updateResult(data.urlA, data.urlB, data.shortUrl, probability[0]);
            setUrl1Value("");
            setUrl2Value("");
            setProbability([.5]); // Reset probability to default
        } catch (error) {
            console.error('Error creating fifty-fifty link:', error);
            // Handle error, e.g., show an error message
        }
    }

    return (
        <Card>
            <CardHeader>
                <CardTitle>
                    Create a Fifty-Fifty Link
                </CardTitle>
                <CardDescription>
                    Enter two URLs
                </CardDescription>
            </CardHeader>
            <CardContent>
                <div className="flex flex-col w-full items-center gap-2">
                    <div className={"flex lg:flex-row flex-col space-y-2 justify-between w-full"}>
                        <div>
                            <Label className={"mb-1 text-indigo-500"} htmlFor="url1">Link 1</Label>
                            <Input className={"md:min-w-xs"} name={"url1"} value={url1Value} onChange={handleUrl1Change}
                                   type="text" placeholder="URL 1"/>


                        </div>
                        <div className={"flex-1 "}></div>
                        <div>
                            <Label className={"mb-1 text-teal-500"} htmlFor="url2">Link 2</Label>
                            <Input className={"md:min-w-xs"} name="url2" value={url2Value} onChange={handleUrl2Change}
                                   type="text" placeholder="URL 2"/>
                        </div>
                    </div>

                        <div className={"w-full flex flex-row items-center justify-between"}>
                            <Label className="flex items-center gap-2 mt-4">
                                Customize Probability
                            </Label>

                            <Switch checked={customize} onCheckedChange={handleToggleCustomize}/>
                        </div>
                    <div className={"w-full flex flex-row  items-center justify-center"}>

                            <span className="w-[5ch] text-center">{Math.round((probability[0] * 100))}%</span>
                            <MySlider disabled={!customize} title={"Link 1 Probability"}
                                      onValueChange={(e) => setProbability(e)} value={probability} max={1} min={0}
                                      step={.01} className="flex-1 mt-4 bg-sky-600"/>
                            <span className="w-[5ch] text-center">{Math.round(100 - (probability[0] * 100))}%</span>

                    </div>

                    <Button onClick={handleSubmit} type="submit" variant="outline" className={"bg-indigo-500 text-white w-full max-w-sm mt-2"}>
                        Create {Math.round(probability[0] * 100)}/{100 - Math.round(probability[0] * 100)} Link
                    </Button>


                </div>
            </CardContent>
        </Card>
    )
}