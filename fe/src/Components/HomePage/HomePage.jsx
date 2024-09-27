import React from 'react'
import Narbar from "../Navbar/Navbar";
import bannerBackGound from "./banner.jpg"
import {FiArrowRight} from "react-icons/fi"
import {useNavigate} from "react-router-dom"

const HomePage = () => {
  const navigate = useNavigate();
  return (
    <div className='Home-container'>
      <Narbar/>
      <div className="home-banner-container">
        <div className="home-bannerImage-container">
          {/* <img src={bannerBackGound} alt=""/> */}
        </div>
        <div className="home-text-section">
          <h1 className='primary-heading'>
            YOUR SAFE IS MY WORK
          </h1>
          <button className='secondary-button' onClick={() => navigate('/loginPage')}>
            LET TRY NOW<FiArrowRight/>
          </button>
        </div>
      </div>
    </div>
  )
}

export default HomePage
