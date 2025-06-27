// components/APIEndpointDoc.tsx
import {useState} from "react";
import {Collapsible, CollapsibleContent, CollapsibleTrigger} from "@/components/ui/collapsible";
import {ChevronRight} from "lucide-react";

interface APIEndpointDocProps {
    title: string;
    defaultOpen?: boolean;
    children: React.ReactNode;
}

export function APIEndpointDoc({ title, defaultOpen = false, children }: APIEndpointDocProps) {
    const [isOpen, setIsOpen] = useState(defaultOpen);

    return (
        
        <Collapsible open={isOpen} onOpenChange={setIsOpen}>
            <CollapsibleTrigger className="flex flex-row items-center w-full mb-4 cursor-pointer space-x-2">
                <ChevronRight
                    className={`h-5 w-5 transition-transform duration-300 ${
                        isOpen ? "rotate-90" : ""
                    }`}
                />
                <h2 className="text-2xl font-bold">
                    {title}
                </h2>
            </CollapsibleTrigger>
            <CollapsibleContent>
                {children}
            </CollapsibleContent>
        </Collapsible>
    );
}
