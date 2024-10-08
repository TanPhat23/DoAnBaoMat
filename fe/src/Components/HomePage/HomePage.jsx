import axios from 'axios'
import React from 'react'
import { useState , useEffect} from 'react'
import { useNavigate } from 'react-router-dom'
import Admin from '../Auth/Admin'
import Employee from '../Auth/Employee'

const HomePage = () => {
  const[user,setUser] = useState();
  const navigate = useNavigate()

  const getUser = async () => {
      try {
        const response = await axios.get("http://localhost:8080/currentuser",
        {withCredentials : true});
        console.log(response.data);
        
        setUser(response.data);
        if(response.status == 404)
          navigate("/login")
      } catch (error) {
        console.log("Logging error: ",error);
      }
  }
    useEffect(()=>{
      getUser()
    },[])

  
  return (
    <div>
      {user ? <div>Welcome {user.Username} to our website!!!
        {user.Role === "admin" ? <Admin/> : <Employee/>}
         </div> : <div></div>}    
    </div>
  );
}

export default HomePage
