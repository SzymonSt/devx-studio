import React from 'react';
import './App.css';
import { Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Dashboard from './pages/Dashboard';
import ClassicSurveys from './pages/ClassicSurveys';
import Backlog from './pages/Backlog';
import 'bootstrap/dist/css/bootstrap.min.css';
import Tags from './pages/Tags';
import Teams from './pages/Teams';
import ContinuousFeedbackPage from './pages/ContinuousFeedbackPage';
import SurveyPage from './pages/SurveyPage';

function App() {
  return (
    <div className='canv'>
      <Navbar />
      <div className='main-screen'>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path='/feed' element={<ContinuousFeedbackPage />}/>
          <Route path='/classic-survey' element={<ClassicSurveys />}/>
          <Route path='/backlog' element={<Backlog />}/>
          <Route path='/tags' element={<Tags />}/>
          <Route path='/teams' element={<Teams />}/>
          <Route path='/survey/:cfId/:surveyId' element={<SurveyPage />} />
        </Routes>
      </div>
    </div>
  );
}

export default App;
