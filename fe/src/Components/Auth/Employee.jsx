import React, { useEffect, useState } from 'react';
import { useNavigate } from "react-router-dom";
import axios from 'axios';

const Employee = () => {
  const [user, setUser] = useState();
  const [image, setImage] = useState();
  const navigate = useNavigate();

  const getUser = async () => {
    try {
      const response = await axios.get("http://localhost:8080/currentuser",
      {withCredentials : true});
      console.log(response.data);
      checkUserImage(response.data.Username);
      setUser(response.data);
      if(response.status == 404)
        navigate("/login")
    } catch (error) {
      console.log("Logging error: ",error);
    } 
  }

  const checkUserImage = (user) => {
    try {
      const importedImage = require(`./photo/${user}.png`);
      setImage(importedImage);
    } catch (error) {
      setImage(null); 
    }
  }

  useEffect(() => {
    getUser();
  }, []);
  return (
    <div>
      {user ? (
        <div className="border border-blue-600 rounded-xl shadow-md p-5 m-2 ml-[600px] text-center max-w-xs inline-block">
          {image ? (
            <img className="w-4/5 h-auto rounded-full mb-3 mx-auto"  src={image} alt={`${user.Username}'s picture`} />
          ) : (
            <p>No image found for {user.Username}.</p>
          )}
          <h2 className="font-sans text-gray-800 m-0">Hello, {user.Username}</h2>
          <div className="font-sans text-gray-700">Let's get to work!!!</div>
        </div>
      ) : (
        <div>No user data available.</div>
      )}
    </div>
  );
};

export default Employee;
