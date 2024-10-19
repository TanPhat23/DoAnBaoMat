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
import LoginIcon from "@mui/icons-material/Login"



const Navbar = () => {
  const [openMenu,setOpenMenu] = useState(false)
  const navigate = useNavigate();
  const menuOptions = [
    {
      text:"Home",
      icon:<HomeIcon/>
    },
    {
      text:"Login",
      icon:<LoginIcon/>
    },
  ]
  return (
    <nav className='justify-end mr-[100px]'>
      <div className="hidden md:flex space-x-10 font-bold text-gray-900 text-[20px] ">
        <a href="" className="  hover:text-blue-600 ">Home</a>
        <a href="" onClick={() => navigate('/login')} className=" hover:text-blue-600">Login</a>
      </div>
      <div className="navbar-menu-container md:hidden">
        <HiOutlineBars3 className="text-2xl cursor-pointer" onClick={() => setOpenMenu(true)} />
      </div>
      <Drawer open={openMenu} onClose={() => setOpenMenu(false)} anchor='right'>
        <Box sx = {{width: 250}} 
          role = "presentation"
          onClick = {() => setOpenMenu(false)}
          onKeyDow = {() => setOpenMenu(false)} 
        >
          <List>
            {menuOptions.map((item) => (
              <ListItem key={item.text} disablePadding >
                <ListItemButton>
                  <ListItemIcon>{item.icon}</ListItemIcon>
                  <ListItemText primary = {item.text}></ListItemText>
                </ListItemButton>
              </ListItem>
            )
            
            )}
          </List>
        </Box>
      </Drawer>
    </nav>
  )
}

export default Navbar