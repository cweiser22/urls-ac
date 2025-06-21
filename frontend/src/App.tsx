import {MainLayout} from "@/layouts/MainLayout.tsx";
import {ShortenURL} from "@/components/ShortenURL.tsx";
import {ShortenURLResult} from "@/components/ShortenURLResult.tsx";
import {useState} from "react";
import {toast} from "sonner";


type Result = {
    longUrl: string;
    shortUrl: string;
}

function App() {
    const [result, setResult] = useState<Result | null>(null);

    const updateResult = (longUrl: string, shortUrl: string) => {
        setResult({ longUrl, shortUrl });
        toast.success("Successfully shortened URL!");
    }

  return (
    <>
      <MainLayout>
          <div className={"container"}>
              <div className={"mx-auto w-full"}>
                <ShortenURL updateResult={updateResult}/>
                  {result && <ShortenURLResult longURL={result.longUrl} shortURL={result.shortUrl} />}
              </div>

          </div>

      </MainLayout>
    </>
  )
}

export default App
