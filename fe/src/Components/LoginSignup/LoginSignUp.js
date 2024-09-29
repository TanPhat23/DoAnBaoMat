import React, { useState } from "react";
import './Global.css'
import axios from "axios";

import acc from "./photo/account.png"
import gmail from "./photo/email.png"
import pass from "./photo/key.png"
import { Navigate, useNavigate } from "react-router-dom";
import {useCookies} from 'react-cookie';


const LoginSignUp = () => {
    const [action, setAction] = useState("Login");
    const [name, setName] = useState("");
    const [password, setPassword] = useState("");
    const [role,setRole]=useState("");
    const navigate = useNavigate()
    
    
    const [cookies, setCookie, removeCookie] = useCookies(['cookie-name'], {
        doNotParse: true,
      });

    const handleSwitch = (newAction) => {
        setAction(newAction);
        setPassword("");
        setName("");
        setRole("");
    };

    const handleLogin = async (e) => { 
    
        try {
            const response = await axios.post("http://localhost:8080/login", {
                username: name,
                password: password,
            });
            
            // Set the token cookie
            if (response.data.Token) {
                // Set the token cookie
                setCookie("token", response.data.Token, { path: '/', maxAge: 600 }); 
                console.log("Login successful, token set");
                if(response.data.Role == "employee") navigate("/Employee")
                else if(response.data.Role == "quest") navigate("/Menu")
            } else {
                console.log("No token received in response");
            }
        } catch (error) {
            console.log("Error logging in:", error);
        }
    };
    
    

    return (
        <div className="container">
            <div className="header">
                <div className="text">{action}</div>
                <div className="underline"></div>
            </div>
            <div className="inputs">
                <div className="input">
                    <img src={acc} alt="" />
                       <input
                            type="text"
                            placeholder="Name"
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                        />
                    </div>
                <div className="input">
                    <img src={pass} alt="" />
                    <input
                        type="password"
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </div>

            </div>
            {action === "Sign Up" ? <div></div> :
                <div className="forgot-password">
                    Lost Password? <span>Click Here!</span>
                </div>
            }
            
            <div className="submit-container">
                <button
                    className={action === "Sign Up" ? "submit gray" : "submit"}
                    onClick={()=> handleLogin()}
                >
                    Login
                </button>
            </div>
        </div>
    );
};

export default LoginSignUp;