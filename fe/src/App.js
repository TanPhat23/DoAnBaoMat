import React from 'react';
import './App.css';

import phanTanPhatImg from './photo/phat.png';
import chuongTuLuanImg from './photo/luan.png';
import ngoDangKhoaImg from './photo/khoa.png';
import voChiThongImg from './photo/thong.png';

function App() {
  return (
    <div id="App" className="background-color">
      <nav>
        <ul id="introduce">
          <li className="hover-container">
            <a href="">
              <img src={phanTanPhatImg} alt="Phan Tấn Phát" className="profile-img" />
              <div className="overlay blue"></div>
              <div className="overlay green"></div>
              <span className="hover-text">Phan Tấn Phát</span>
            </a>
          </li>
          <li className="hover-container">
            <a href="">
              <img src={chuongTuLuanImg} alt="Chương Tử Luân" className="profile-img" />
              <div className="overlay blue"></div>
              <div className="overlay green"></div>
              <span className="hover-text">Chương Tử Luân</span>
            </a>
          </li>
          <li className="hover-container">
            <a href="">
              <img src={ngoDangKhoaImg} alt="Ngô Đăng Khoa" className="profile-img" />
              <div className="overlay blue"></div>
              <div className="overlay green"></div>
              <span className="hover-text">Ngô Đăng Khoa</span>
            </a>
          </li>
          <li className="hover-container">
            <a href="">
              <img src={voChiThongImg} alt="Võ Chí Thông" className="profile-img" />
              <div className="overlay blue"></div>
              <div className="overlay green"></div>
              <span className="hover-text">Võ Chí Thông</span>
            </a>
          </li>
        </ul>
      </nav>
    </div>
  );
}

export default App;
