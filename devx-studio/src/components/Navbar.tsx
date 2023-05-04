import React from "react";
import { MdOutlineSpaceDashboard, MdLoop } from 'react-icons/md';
import {RiSurveyLine, RiTodoLine} from 'react-icons/ri';
import "./Navbar.css"

function Navbar() {
    return (
        <div className="navbar">
            <div className="navbar-logo">
                <h5 className="logo">DevX Studio</h5>
            </div>
            <div className="navbar-menu">
                <h3><a href="/"><MdOutlineSpaceDashboard /></a></h3>
                <h3><a href="/feed"><MdLoop /></a></h3>
                <h3><a href="/classic-survey"><RiSurveyLine /></a></h3>
                <h3><a href="/backlog"><RiTodoLine /></a></h3>
            </div>
        </div>
    )
}

export default Navbar