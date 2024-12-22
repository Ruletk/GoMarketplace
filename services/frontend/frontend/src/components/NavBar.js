import React from "react";

function NavBar() {
  return (
    <nav>
      <ul>
        <li>
          <a href="/login">Login</a>
        </li>
        <li>
          <a href="/registration">Registration</a>
        </li>
        <li>
          <a href="/changepassword">Change Password</a>
        </li>
        <li>
          <a href="/requestpassword">Request Password</a>
        </li>
        <li>
          <a href="/verify">Verify</a>
        </li>
      </ul>
    </nav>
  );
}

export default NavBar;