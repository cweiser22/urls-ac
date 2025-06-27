import {createFileRoute} from "@tanstack/react-router";
import {Collapsible, CollapsibleContent, CollapsibleTrigger} from "@/components/ui/collapsible.tsx";
import {ChevronDown, ChevronRight} from "lucide-react";
import {useState} from "react";
import {APIEndpointDoc} from "@/components/APIEndpointDoc.tsx";

export const Route = createFileRoute('/apidocs')({
    component: APIDocs,
})

const shortenEndpointDocs = (
    <>
        <h1 className="text-3xl font-bold mb-4">
            API Endpoint:{' '}
            <code className="bg-gray-100 px-2 py-1 rounded text-blue-600">
                /api/v1/shorten
            </code>
        </h1>
    <p className="mb-6 text-gray-700">
        Shortens a valid URL and returns both the original and shortened version.
    </p>

<div className="mb-6">
    <h2 className="text-xl font-semibold mb-2">Method</h2>
    <code className="bg-gray-100 px-2 py-1 rounded text-green-600">GET</code>
</div>

<div className="mb-6">
    <h2 className="text-xl font-semibold mb-2">Content-Type</h2>
    <code className="bg-gray-100 px-2 py-1 rounded">application/json</code>
</div>

<div className="mb-6">
    <h2 className="text-xl font-semibold mb-2">Request Payload</h2>
    <pre className="bg-gray-100 p-4 rounded text-sm overflow-x-auto">
{`{
  "longUrl": "https://www.example.com/some/very/long/url"
}`}
        </pre>
    <table className="w-full mt-4 text-left border border-gray-200">
        <thead className="bg-gray-50">
        <tr>
            <th className="border px-4 py-2">Field</th>
            <th className="border px-4 py-2">Type</th>
            <th className="border px-4 py-2">Required</th>
            <th className="border px-4 py-2">Description</th>
        </tr>
        </thead>
        <tbody>
        <tr>
            <td className="border px-4 py-2">longUrl</td>
            <td className="border px-4 py-2">string</td>
            <td className="border px-4 py-2">Yes</td>
            <td className="border px-4 py-2">
                The full URL to be shortened. Must be a valid absolute URL
                (including <code>http</code> or <code>https</code>).
            </td>
        </tr>
        </tbody>
    </table>
</div>

<div className="mb-6">
    <h2 className="text-xl font-semibold mb-2">Responses</h2>

    <h3 className="font-medium mt-4">✅ 200 OK</h3>
    <pre className="bg-green-50 p-4 rounded text-sm overflow-x-auto">
{`{
  "longUrl": "https://www.example.com/some/very/long/url",
  "shortUrl": "https://short.ly/abc123"
}`}
        </pre>

    <h3 className="font-medium mt-4 text-red-600">❌ 400 Bad Request</h3>
    <pre className="bg-red-50 p-4 rounded text-sm overflow-x-auto">
{`{
  "error": "Invalid URL"
}`}
        </pre>
</div>

<div className="mb-6">
    <h2 className="text-xl font-semibold mb-2">Example Request (curl)</h2>
    <pre className="bg-gray-100 p-4 rounded text-sm overflow-x-auto">
{`curl -X GET https://yourdomain.com/api/v1/shorten \\
  -H "Content-Type: application/json" \\
  -d '{"longUrl": "https://www.example.com"}'`}
        </pre>
</div>

<div>
    <h2 className="text-xl font-semibold mb-2">Notes</h2>
    <p className="text-gray-700">
        This endpoint expects the payload in the <strong>body</strong> of a{" "}
        <code>GET</code> request, which is not standard and may not be supported
        by all clients. Consider using{" "}
        <code className="bg-gray-100 px-1 rounded">POST</code> for improved
        compatibility.
    </p>
</div> </>
);

const fiftyFiftyDocs: React.ReactNode = (
    <div className="max-w-4xl mx-auto p-6 bg-white shadow rounded-xl text-gray-900">
        <h1 className="text-3xl font-bold mb-4">
            API Endpoint:{' '}
            <code className="bg-gray-100 px-2 py-1 rounded text-blue-600">
                /api/v1/fiftyfifty/
            </code>
        </h1>

        <p className="mb-6 text-gray-700">
            Creates a "50/50" short link that redirects to one of two URLs with a given
            probability.
        </p>

        <div className="mb-6">
            <h2 className="text-xl font-semibold mb-2">Method</h2>
            <code className="bg-gray-100 px-2 py-1 rounded text-purple-600">POST</code>
        </div>

        <div className="mb-6">
            <h2 className="text-xl font-semibold mb-2">Content-Type</h2>
            <code className="bg-gray-100 px-2 py-1 rounded">application/json</code>
        </div>

        <div className="mb-6">
            <h2 className="text-xl font-semibold mb-2">Request Body</h2>
            <pre className="bg-gray-100 p-4 rounded text-sm overflow-x-auto">
{`{
  "probability": 0.7,
  "urlA": "https://example.com/winner",
  "urlB": "https://example.com/loser"
}`}
      </pre>

            <table className="w-full mt-4 text-left border border-gray-200">
                <thead className="bg-gray-50">
                <tr>
                    <th className="border px-4 py-2">Field</th>
                    <th className="border px-4 py-2">Type</th>
                    <th className="border px-4 py-2">Required</th>
                    <th className="border px-4 py-2">Description</th>
                </tr>
                </thead>
                <tbody>
                <tr>
                    <td className="border px-4 py-2">probability</td>
                    <td className="border px-4 py-2">float</td>
                    <td className="border px-4 py-2">No</td>
                    <td className="border px-4 py-2">
                        A number between 0 and 1 indicating the chance of redirecting to
                        <code>urlA</code>. Defaults to <code>0.5</code> if omitted.
                    </td>
                </tr>
                <tr>
                    <td className="border px-4 py-2">urlA</td>
                    <td className="border px-4 py-2">string</td>
                    <td className="border px-4 py-2">Yes</td>
                    <td className="border px-4 py-2">The first target URL.</td>
                </tr>
                <tr>
                    <td className="border px-4 py-2">urlB</td>
                    <td className="border px-4 py-2">string</td>
                    <td className="border px-4 py-2">Yes</td>
                    <td className="border px-4 py-2">The second target URL.</td>
                </tr>
                </tbody>
            </table>
        </div>

        <div className="mb-6">
            <h2 className="text-xl font-semibold mb-2">Responses</h2>

            <h3 className="font-medium mt-4">✅ 200 OK</h3>
            <pre className="bg-green-50 p-4 rounded text-sm overflow-x-auto">
{`{
  "shortUrl": "https://yourdomain.com/ff/abc123",
  "urlA": "https://example.com/winner",
  "urlB": "https://example.com/loser"
}`}
      </pre>

            <h3 className="font-medium mt-4 text-red-600">❌ 400 Bad Request</h3>
            <pre className="bg-red-50 p-4 rounded text-sm overflow-x-auto">
{`{
  "error": "Invalid JSON"
}`}
      </pre>
            <pre className="bg-red-50 p-4 rounded text-sm overflow-x-auto mt-2">
{`{
  "error": "Invalid URL A: must be a valid absolute URL"
}`}
      </pre>
            <pre className="bg-red-50 p-4 rounded text-sm overflow-x-auto mt-2">
{`{
  "error": "Probability must be between 0 and 1"
}`}
      </pre>

            <h3 className="font-medium mt-4 text-red-600">❌ 500 Internal Server Error</h3>
            <pre className="bg-red-50 p-4 rounded text-sm overflow-x-auto">
{`{
  "error": "Failed to create link"
}`}
      </pre>
        </div>

        <div className="mb-6">
            <h2 className="text-xl font-semibold mb-2">Example Request (curl)</h2>
            <pre className="bg-gray-100 p-4 rounded text-sm overflow-x-auto">
{`curl -X POST https://yourdomain.com/api/v1/fiftyfifty \\
  -H "Content-Type: application/json" \\
  -d '{
    "probability": 0.6,
    "urlA": "https://example.com/winner",
    "urlB": "https://example.com/loser"
  }'`}
      </pre>
        </div>

        <div>
            <h2 className="text-xl font-semibold mb-2">Notes</h2>
            <p className="text-gray-700">
                This endpoint generates a short code based on the given inputs and returns
                a redirectable short URL. If the <code>probability</code> is omitted, it
                defaults to <code>0.5</code>.
            </p>
        </div>
    </div>
);



function APIDocs(){
    return (
        <div className={"container mx-auto max-w-3xl px-4 py-8"}>
            <h1 className="text-3xl font-bold mb-4">API Documentation</h1>

            <APIEndpointDoc title={"Shorten URL"} defaultOpen={true} children={shortenEndpointDocs}/>
            <APIEndpointDoc title={"Create FiftyFifty Link"} defaultOpen={false} children={fiftyFiftyDocs}/>
        </div>
    )
}
