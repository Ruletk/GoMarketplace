import React from "react";
import '../styles/NavBar.css';

function NavBar() {
    return (
        <nav>
            <div class="logo"><a href="/home"> GoMarketPlace</a></div>
    <div class="search-bar">
      <input type="text" placeholder="What are you looking for?!" />
      <button>Search</button>
    </div>
    <div class="right-section">
    
      <a href="#">Catalog</a>
      <a href="#">Cart</a>
      <a href="/registration">Sign up</a>
    </div>
  </nav>
    );
}

export default NavBar;
