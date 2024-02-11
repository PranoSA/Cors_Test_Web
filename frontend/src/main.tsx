import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './Home/index'
import Application from './Application'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import ViewResults from './View_Results'

const router = createBrowserRouter([
  {
    path : "/",
    element : <App />
  },
  {
    path : "/application/:id",
    element : <Application />
  },
  {
    path : "/results/:id",
    element : <ViewResults />
  }
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
