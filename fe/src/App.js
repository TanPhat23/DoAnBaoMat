import React from 'react';
import "./App.css"

// import phanTanPhatImg from './photo/phat.png';
// import chuongTuLuanImg from './photo/luan.png';
// import ngoDangKhoaImg from './photo/khoa.png';
// import voChiThongImg from './photo/thong.png';


import { BrowserRouter ,Route, Routes} from 'react-router-dom';
<<<<<<< HEAD
import LoginSignUp from './Components/LoginSignup/LoginSignUp';
=======

import LandingPage from './Components/HomePage/LandingPage';
import LoginSignUp from './Components/LoginSignup/Login';
>>>>>>> acaa9e204e1c8694837405112d07192b3df6c443
import HomePage from './Components/HomePage/HomePage';
function App() {

  return (
<<<<<<< HEAD
    <div id="App">
    <HomePage/>
  </div>
  );
}
export default App;
=======
    <div className="App">
    <BrowserRouter>
      <Routes>
        <Route path = "/" element={<LandingPage/>}/>
        <Route path ="/login" element ={<LoginSignUp/>}/>
        <Route path= "/home" element ={<HomePage/>}/>
      </Routes>
    </BrowserRouter>
    </div>
  );
}
export default App;
>>>>>>> acaa9e204e1c8694837405112d07192b3df6c443
