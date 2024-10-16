import React from 'react'
import { useNavigate } from 'react-router-dom'
import Logout from '../LoginSignup/logout'

const Employee = () => {
  const navigate = useNavigate();
  return (
    <div className='employee-container'>
        <div>Hello, Employee</div>
        <div>Let's get to work!!!</div>
        <button className='' onClick={() => navigate("/logout")}>
          Logout
        </button>
    </div>
  )
}

export default Employee
