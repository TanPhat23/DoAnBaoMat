import React from 'react';
import "./App.css"

// import phanTanPhatImg from './photo/phat.png';
// import chuongTuLuanImg from './photo/luan.png';
// import ngoDangKhoaImg from './photo/khoa.png';
// import voChiThongImg from './photo/thong.png';


import { BrowserRouter ,Route, Routes} from 'react-router-dom';
import LoginSignUp from './Components/LoginSignup/LoginSignUp';
import HomePage from './Components/HomePage/HomePage';
function App() {

  return (
    <div id="App">
    <HomePage/>
  </div>
  );
}
export default App;