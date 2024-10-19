import React, { useState } from 'react'
import {BsCart2} from "react-icons/bs";
import {HiOutlineBars3} from "react-icons/hi2";
import {useNavigate} from "react-router-dom"
import {
  Box,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import HomeIcon from "@mui/icons-material/Home";
import LogoutIcon from '@mui/icons-material/Logout';
import axios from 'axios';



const NavbarLogout = () => {
  const [openMenu,setOpenMenu] = useState(false)
  const navigate = useNavigate();
  const handleLogout = async (e) => { 
    try {
      const response = await axios.delete("http://localhost:8080/logout",{ withCredentials: true });
      navigate("/");
    } catch (error) {
      console.log("Error logging in:", error);
    }
  }
  const menuOptions = [
    {
      text:"Home",
      icon:<HomeIcon/>
    },
    {
      text:"Logout",
      icon:<LogoutIcon/>
    },
  ]
  return (
    <nav className='justify-end mr-[100px]'>
      <div className="hidden md:flex space-x-10 font-bold text-gray-900 text-[20px] ">
        <a href="/" className="  hover:text-blue-600 ">Home</a>
        <a href="/" onClick={() =>  handleLogout()} className=" hover:text-blue-600">Logout</a>
      </div>
      <div className="navbar-menu-container md:hidden">
        <HiOutlineBars3 className="text-2xl cursor-pointer" onClick={() => setOpenMenu(true)} />
      </div>
     
    </nav>
  )
}

export default NavbarLogout