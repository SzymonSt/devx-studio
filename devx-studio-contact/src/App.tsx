import React, { useEffect, useState } from 'react';
import { Routes, Route } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import DemoRequest from './Pages/DemoRequest';
import Waitlist from './Pages/Waitlist';

function App() {

  return (
    <div className='main-screen'>
    <Routes>
      <Route path="/" element={<DemoRequest />} />
      <Route path='/waitlist' element={<Waitlist />}/>
    </Routes>
  </div>
  );
}

export default App;
