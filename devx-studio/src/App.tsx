import React from 'react';
import './App.css';
import { Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Dashboard from './pages/Dashboard';
import ContinuousFeedback from './pages/ContinuousFeedback';
import ClassicSurveys from './pages/ClassicSurveys';
import Backlog from './pages/Backlog';
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  return (
    <div className='canv'>
      <Navbar />
      <div className='main-screen'>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path='/feed' element={<ContinuousFeedback />}/>
          <Route path='/classic-survey' element={<ClassicSurveys />}/>
          <Route path='/backlog' element={<Backlog />}/>
        </Routes>
      </div>
    </div>
  );
}

export default App;
