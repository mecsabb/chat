// to be redone with tailwind

import { useState } from 'react'
import './test.css'

export default function Test() {

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

  function postForm(){
    fetch('http://localhost:3333/login',
    {
      method: "POST"
    })
    .then(response => response.json())
    .then(json => setData(json))
    .catch(error => console.error(error));
  }

  return (
    <>
      <h1 className='title'>Testing Dashboard</h1>
      <div className='container'>
        <div className="test-container">
          <div className="button-text">
            <button className='get-button' onClick={onClickRequest}>
              GET ROOT
            </button>
            <div className='request-text'>
              {data ? <pre>{JSON.stringify(data, null, 2)}</pre> : '...'}
            </div>
          </div>
          <div className="button-text">
            <button className='get-button' onClick={onClickRequest}>
              GET LOGIN
            </button>
            <div>
              {/* {data ? <pre>{JSON.stringify(data, null, 2)}</pre> : '...'} */}
            </div>
          </div>
          <div className="button-text">
            <button className='get-button' onClick={onClickRequest}>
              GET TEMP
            </button>
            <div>
              {/* {data ? <pre>{JSON.stringify(data, null, 2)}</pre> : '...'} */}
            </div>
          </div>
        </div>
        <div className='form-wrapper'>
          <form className='login-form' onSubmit={postForm}>
            <label className='input'>
              Username:
              <input className='input' type="text" name="name" />
            </label>
            <label className='input'>
              Password:
              <input className='input' type="text" name="pass" />
            </label>
            <input className='submit-button' type="submit" value="Submit" />
          </form>
        </div>
      </div>
    </>
  )
}
