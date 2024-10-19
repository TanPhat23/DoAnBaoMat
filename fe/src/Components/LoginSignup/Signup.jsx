import React, { useState } from "react";
import './Global.css'
import axios from "axios";

import acc from "./photo/account.png"
import pass from "./photo/key.png"
import { Router, useNavigate } from "react-router-dom";


const Signup = () => {
    const [action, setAction] = useState("SignUp");
    const [name, setName] = useState("");
    const [password, setPassword] = useState("");
    const [checkPassword, setCheckPassword] = useState(true);
    const [confirmPassword, setConfirmPassword] = useState("");
    const navigate = useNavigate()  

    const checkPasswordfunc = ()=>{
        if( password === confirmPassword){
            return true
        }else{
            setCheckPassword(false)
            return false
        }
    }


    const handleSignup = async (e) => { 
            try {
                if(checkPasswordfunc()){
                    const response = await axios.post("http://localhost:8080/signin", {
                        username: name,
                        password: password,
                    }, { withCredentials: true });
                    if (response.status === 200) {
                        console.log("Login successful")
                        navigate("/home")        
                    }
                }
            } catch (error) {
                console.log("Error logging in:", error);
            }
    };


    return (
        <div className="flex flex-col mt-[50px] justify-center ml-[350px] bg-white rounded-xl pb-[30px] w-[800px]">  
            <div className="flex flex-col items-center mt-2 gap-2 w-full">
                <div className="text-black text-4xl font-bold">{action}</div>
                <div className="w-[61px] h-[6px] bg-[#3c009d] rounded-md"></div>
            </div>
            <div className="inputs flex justify-center">
                <div className="input outline">
                    <img src={acc} alt="" />
                       <input
                            type="text"
                            placeholder="Username"
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                        />
                    </div>
                <div className="input outline">
                    <img src={pass} alt="" />
                    <input
                        type="password"
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </div>
                <div className="input outline " >
                    
                    <img src={pass} alt="" />
                    <input
                        type="password"
                        placeholder="ConfirmPassword"
                        value={confirmPassword}
                        onChange={(e) => setConfirmPassword(e.target.value)}
                    />
                    
                </div>
                

            </div>
            {checkPassword ?  "": <p className="checkpassword pl-[200px] mt-[10px] text-red-700">The password does not match</p>}
                {action === "Sign Up" ? <div></div> :
                    <div className="flex justify-center mt-8 mr-[250px]">
                        <p className="font-bold">Lost Password?</p> <span className="ml-2 text-purple-700 hover:text-blue-600 cursor-pointer">Click Here!</span>
                    </div>
                }
                
            
            <div className="flex justify-center mt-14 gap-5">
                
                <button
                    className={action === "Sign Up" ? "submit gray" : "submit"}
                    onClick={()=> handleSignup()}
                >
                    Sign Up
                </button>
                <button
                    className={action === "Sign Up" ? "submit gray" : "submit"}
                    onClick={()=> navigate("/login")}
                >
                    Login
                </button>
            </div>
        </div>
    );
};

export default Signup;