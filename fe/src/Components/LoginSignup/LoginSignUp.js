import React, { useState } from "react";
import './Global.css'

import acc from "./photo/account.png"
import pass from "./photo/key.png"


const LoginSignUp = () => {
    const [action, setAction] = useState("Login");
    const [name, setName] = useState("");
    const [password, setPassword] = useState("");

    // Hàm chuyển trang và xóa nội dung các trường nhập liệu
    const handleSwitch = (newAction) => {
        setAction(newAction);
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
                {action === "Login" ? <div></div> :<div className="input"> </div>}
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
                <div className="forgot-password">
                    Lost Password? <span>Click Here!</span>
                </div>
            <div className="submit-container">
                <nav
                    className="Login"
                    onClick={() => handleSwitch("Login")}
                >
                    Login
                </nav>
            </div>
        </div>
    );
};

export default LoginSignUp;
