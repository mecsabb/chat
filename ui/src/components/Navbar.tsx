// TODO implement sticky navbar

export default function Navbar() {
    return (
        <>
        <header>
            <nav className="dark:bg-gray-800 border-gray-700 px-4 py-2.5">
                <div className="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl">
                    <a href="/" className="text-xl font-extrabold px-5 text-white">Chat</a>
                    <div className="flex items-center lg:order-2">
                        <a href="/login" className="text-white px-10">Log in</a>
                        <a href="/#" className="text-white px-10">Get started</a>
                    </div>
                </div>
            </nav>
        </header>
        </>
    )
}