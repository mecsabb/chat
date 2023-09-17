import { useState } from 'react'
import './App.css'

function App() {

  const [data, setData] = useState(null);

  function onClickRequest() {
    fetch('http://localhost:3333',
    {
      method: "GET"
    })
    .then(response => response.json())
    .then(json => setData(json))
    .catch(error => console.error(error));
  }

  return (
    <>
      <h1>Go + React Chat App</h1>
      <div className="card">
        <button onClick={onClickRequest}>
          QUERY BACKEND
        </button>
      </div>
      <div>
        {data ? <pre>{JSON.stringify(data, null, 2)}</pre> : 'Loading...'}
      </div>
    </>
  )
}

export default App
