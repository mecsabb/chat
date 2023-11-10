import React from 'react'
import ReactDOM from 'react-dom/client'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"

import Home from './routes/root.tsx'
import Test from './routes/test.tsx'
import LoginPage from './routes/login.tsx'
import RegisterPage from './routes/register.tsx'
import ErrorPage from './error-page.tsx'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
    errorElement: <ErrorPage />
  },
  {
    path: "test",
    element: <Test />,
    errorElement: <ErrorPage />
  },
  {
    path: "login",
    element: <LoginPage />,
    errorElement: <ErrorPage />
  },
  {
    path: "register",
    element: <RegisterPage />,
    errorElement: <ErrorPage />
  }
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
