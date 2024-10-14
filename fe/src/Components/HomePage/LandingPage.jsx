import React, { useEffect, useState } from 'react';
import Narbar from "../Navbar/Navbar";
import bannerBackGound from "./banner.jpg";
import { FiArrowRight } from "react-icons/fi";
import { useNavigate } from "react-router-dom";

const LandingPage = () => {
  const navigate = useNavigate();
  
  return (
    <div className="div">
      <Narbar />
      <div className="relative flex pt-12">
        <div className="fixed top-0 left-0 w-screen h-screen z-[-2] overflow-hidden">
          <img src={bannerBackGound} alt="" />
        </div>
        <div className="flex-1 flex flex-col justify-center items-start">
          <h1 className="text-[4rem] text-gray-700 max-w-full pl-[10%]">YOUR SAFE IS MY WORK</h1>
          <button
            className="secondary-button"
            onClick={() => navigate("/login")}
          >
            LET TRY NOW
            <FiArrowRight/>
          </button>
        </div>
      </div>  
    </div>
  );
}

export default LandingPage;