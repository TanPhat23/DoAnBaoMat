import React from 'react';

// import phanTanPhatImg from './photo/phat.png';
// import chuongTuLuanImg from './photo/luan.png';
// import ngoDangKhoaImg from './photo/khoa.png';
// import voChiThongImg from './photo/thong.png';


import { BrowserRouter ,Route, Routes} from 'react-router-dom';
import LoginSignUp from './Components/LoginSignup/LoginSignUp';


function App() {

  return (
    <BrowserRouter>
      <div id="App">
        <Routes>
          <Route path="/login" element={<LoginSignUp/>} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}
export default App;
