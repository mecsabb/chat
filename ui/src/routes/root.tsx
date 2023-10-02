import './root.css'

export default function Home() {
    return (
        <>
            <h1>"HOMEPAGE"</h1>
            <div className="dropdown">
            <div className="dropbtn">+</div>
            <div className="dropdown-content">
                <a href="#">Login</a>
                <a href="#">Sign Up</a>
            </div>
            </div>
        </>
    )
}