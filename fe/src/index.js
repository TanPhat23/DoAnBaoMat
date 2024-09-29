import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import LoginPage from "./Components/LoginSignup/LoginSignUp";
import Employee from './Components/Auth/Employee';
import Menu from './Components/Auth/Quest';

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
]);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
reportWebVitals();
