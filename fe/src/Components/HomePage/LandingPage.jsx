import React, { useEffect, useState } from 'react';
import Narbar from "../Navbar/Navbar";
import bannerBackGound from "./banner.jpg";
import { FiArrowRight } from "react-icons/fi";
import { useNavigate } from "react-router-dom";

const LandingPage = () => {
  const navigate = useNavigate();
  
  return (
    <div className="div">
        <div className='h-screen w-screens'>
          <Narbar />
        </div>
        <div className="home-bannerImage-container ">
          <img src={bannerBackGound} alt=""/>
        </div>
        <div className="flex-col ml-[100px] mt-[100px] ">
          <h1 className="primary-heading ">YOUR SAFE IS MY WORK</h1>
          <button 
            className="secondary-button"
            onClick={() => navigate("/login")}
          >
            LET TRY NOW
            <FiArrowRight />
          </button> 
        </div>
      </div> 
   
  );
}

export default LandingPage;