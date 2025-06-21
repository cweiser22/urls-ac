export function Footer(){
    return (
        <footer className="w-full border-t px-4 py-4">
            <div className="container mx-auto text-center text-sm">
                <p>© {new Date().getFullYear()} Cooper Weiser - urls.ac, a modern URL shortener</p>
                <p>Made with ❤️</p>
            </div>
        </footer>
    )
}