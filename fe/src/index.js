import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import LoginPage from "./Components/LoginSignup/Login";
import Employee from './Components/Auth/Employee';
import Menu from './Components/HomePage/HomePage';
import Admin from './Components/Auth/Admin';

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App/>,
  },
  {
    path: "loginPage",
    element: <LoginPage/>,
  },
  {
    path: "Employee",
    element: <Employee/>,
  },
  {
    path: "Quest",
    element: <Menu/>,
  },
  {
    path: "Admin",
    element: <Admin/>,
  },
]);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
reportWebVitals();
