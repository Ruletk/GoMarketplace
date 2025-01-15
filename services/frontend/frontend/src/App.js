import React from 'react';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import NavBar from './components/NavBar'; // Путь к компоненту навигации
import Login from './pages/Login';
import Registration from './pages/Registration';
import RequestPassword from './pages/RequestPassword';
import ChangePassword from './pages/ChangePassword';
import Verify from './pages/Verify';
import AdminEndpointsPage from './pages/AdminEndpointsPage'; // Импортируем страницу с эндпоинтами
import Home from './pages/Home'; // Импортируем компонент Home
import ProductPage from './pages/ProductPage'; // Импортируем компонент ProductPage
import GoodPage from './pages/GoodPage'; // Импортируем компонент ProductDetails


function App() {
    return (
        <Router>
            <NavBar/>
            <Routes>
                <Route path="/home" element={<Home/>}/>
                <Route path="/login" element={<Login/>}/>
                <Route path="/registration" element={<Registration/>}/>
                <Route path="/requestpassword" element={<RequestPassword/>}/>
                <Route path="/changepassword" element={<ChangePassword/>}/>
                <Route path="/verify" element={<Verify/>}/>
                <Route path="/admin-endpoints" element={<AdminEndpointsPage/>}/> {/* Новый маршрут */}
                <Route path="/product-page" element={<ProductPage/>}/> {/* Новый маршрут */}
                <Route path="/product/:id" element={<GoodPage />} />
            </Routes>
        </Router>
    );
}

export default App;
