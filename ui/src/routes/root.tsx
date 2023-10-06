//TODO fix image serving 

export default function Home() {
    return (
        <>
        <header>
            <nav className="dark:bg-gray-800 border-gray-700 px-4 py-2.5">
                <div className="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl">
                    <span className="text-xl font-extrabold px-5 text-white">Chat</span>
                    <div className="flex items-center lg:order-2">
                        <a href="/login" className="text-white px-10">Log in</a>
                        <a href="/#" className="text-white px-10">Get started</a>
                    </div>
                </div>
            </nav>
        </header>
        <body className="h-screen bg-gray-400"> 
            <div className="relative flex flex-col items-center justify-center px-6 py-8 mx-auto top-1/3 w-1/3">
                <img src='../assets/speech-bubble-1078.png' className=""></img>
                <div className="absolute top-1/2 font-semibold text-white break-words">
                    Introducing ChatConnect, the revolutionary chat app that's redefining communication in the digital age. 
                    Say goodbye to cluttered interfaces and hello to a sleek, user-friendly platform that prioritizes your conversations. 
                    With ChatConnect, you can effortlessly connect with friends, family, and colleagues in real-time, no matter 
                    where they are in the world. Our app boasts state-of-the-art encryption for maximum privacy and security, ensuring 
                    your messages stay confidential.
                </div>
            </div>
        </body>
        </>
    )
}
