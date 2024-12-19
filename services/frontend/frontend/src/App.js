import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./pages/Login";
import ChangePassword from "./pages/ChangePassword";
import Registration from "./pages/Registration";
import RequestPassword from "./pages/RequestPassword";
import Verify from "./pages/Verify";
import "./styles/App.css";
import NavBar from "./components/NavBar";

function App() {
  return (
    <Router>
      <NavBar />
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/registration" element={<Registration />} />
        <Route path="/requestpassword" element={<RequestPassword />} />
        <Route path="/changepassword" element={<ChangePassword />} />
        <Route path="/verify" element={<Verify />} />
      </Routes>
    </Router>
  );
}

export default App;
