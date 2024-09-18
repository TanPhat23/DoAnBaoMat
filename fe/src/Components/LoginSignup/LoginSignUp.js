import React, { useState } from "react";
import './Global.css'

import acc from "./photo/account.png"
import gmail from "./photo/email.png"
import pass from "./photo/key.png"


const LoginSignUp = () => {
    const [action, setAction] = useState("Sign Up");
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    // Hàm chuyển trang và xóa nội dung các trường nhập liệu
    const handleSwitch = (newAction) => {
        setAction(newAction);
        setEmail("");
        setPassword("");
        setName("");
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
                {action === "Login" ? <div></div> :

                <div className="input">
                    <img src={gmail} alt="" />
                    <input
                        type="email"
                        placeholder="Email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                </div>
                }
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
                    className={action === "Login" ? "submit gray" : "submit"}
                    onClick={() => handleSwitch("Sign Up")}
                >
                    Sign Up
                </button>
                <button
                    className={action === "Sign Up" ? "submit gray" : "submit"}
                    onClick={() => handleSwitch("Login")}
                >
                    Login
                </button>
            </div>
        </div>
    );
};

export default LoginSignUp;
