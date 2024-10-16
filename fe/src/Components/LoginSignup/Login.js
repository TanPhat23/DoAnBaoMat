import React, { useState } from "react";
import './Global.css'
import axios from "axios";

import acc from "./photo/account.png"
import pass from "./photo/key.png"
import { useNavigate } from "react-router-dom";


const Login = () => {
    const [action, setAction] = useState("Login");
    const [name, setName] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();
    const handleLogin = async (e) => { 
        try {
            const response = await axios.post("http://localhost:8080/login", {
                username: name,
                password: password,
            }, { withCredentials: true });
            if (response.status === 200) {
                console.log("Login successful",navigate("/home"));        
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
                        onClick={() => handleLogin()}
                    >
                        Login
                    </button>
                </div>
        </div>
    );
};

export default Login;