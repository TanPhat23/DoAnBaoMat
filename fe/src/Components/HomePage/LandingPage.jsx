import React, { useEffect, useState } from 'react';
import Narbar from "../Navbar/Navbar";
import bannerBackGound from "./banner.jpg";
import { FiArrowRight } from "react-icons/fi";
import { useNavigate } from "react-router-dom";
import TypeWriter from "typewriter-effect";


const LandingPage = () => {
  const navigate = useNavigate();
  
  return (
    <div className="Home-container">
      <Narbar />
      <div className="home-banner-container">
        <div className="home-bannerImage-container">
          <img src={bannerBackGound} alt=""/>
        </div>
        <div className="home-text-section">
          <h1 className="primary-heading">
            <TypeWriter
              options={{
                strings: ['Welcome to our website','YOUR SAFE IS OUR WORK'],
                autoStart: true,
                loop: true,
                deleteSpeed: 'natural',
                delay:'natural',
              }}
            />
          </h1>
          <button
            className="secondary-button"
            onClick={() => navigate("/login")}
          >
            LET TRY NOW
            <FiArrowRight />
          </button>
        </div>
      </div>  
    </div>
  );
}



export default LandingPage;