import React, { useState } from "react";
import "../styles/NavBar.css";

function NavBar() {
  const [overlayOpen, setOverlayOpen] = useState(false);

  const toggleOverlay = () => {
    setOverlayOpen(!overlayOpen);
  };

  const closeOverlay = () => {
    setOverlayOpen(false);
  };

  return (
    <nav className="navbar">
      <div className="logo">
        <a href="/home">GoMarketPlace</a>
      </div>
      <div className="search-bar">
        <input type="text" placeholder="What are you looking for?!" />
        <button>Search</button>
        <button onClick={toggleOverlay} className="dropdown-button">
          Filters
        </button>
      </div>
      <div className="right-section">
        <a href="/product-page">Catalog</a>
        <a href="#">Cart</a>
        <a href="/registration">Sign up</a>
      </div>

      {/* Оверлей для фильтров */}
      {overlayOpen && (
        <div className="overlay" onClick={closeOverlay}>
          <div className="modal" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header">
              <h2>Choose Filters</h2>
              <button onClick={closeOverlay} className="close-button">X</button>
            </div>
            <div className="modal-body">
              <label>
                <input type="checkbox" name="filter1" />
                Option 1
              </label>
              <label>
                <input type="checkbox" name="filter2" />
                Option 2
              </label>
              <label>
                <input type="checkbox" name="filter3" />
                Option 3
              </label>
            </div>
            <div className="modal-footer">
              <button>Apply Filters</button>
            </div>
          </div>
        </div>
      )}
    </nav>
  );
}

export default NavBar;
