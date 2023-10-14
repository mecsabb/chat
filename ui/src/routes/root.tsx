//TODO fix image serving 

import Navbar from "../components/Navbar"

export default function Home() {
    return (
        <>
        <Navbar/>
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
