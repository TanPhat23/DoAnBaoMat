import React from 'react';
import "./App.css"

// import phanTanPhatImg from './photo/phat.png';
// import chuongTuLuanImg from './photo/luan.png';
// import ngoDangKhoaImg from './photo/khoa.png';
// import voChiThongImg from './photo/thong.png';


import { BrowserRouter ,Route, Routes} from 'react-router-dom';

import LandingPage from './Components/HomePage/LandingPage';
import Login from './Components/LoginSignup/Login';
import HomePage from './Components/HomePage/HomePage';
import Logout from './Components/LoginSignup/logout';
function App() {

  return (
    <div className="App">
    <BrowserRouter>
      <Routes>
        <Route path = "/" element={<LandingPage/>}/>
        <Route path = "/login" element ={<Login/>}/>
        <Route path = "/home" element ={<HomePage/>}/>
        <Route path = "/logout" element ={<Logout/>}/>
      </Routes>
    </BrowserRouter>
    </div>
  );
}
export default App;