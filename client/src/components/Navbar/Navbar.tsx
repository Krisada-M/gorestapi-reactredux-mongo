import React from "react";
import "./Navbar.scss";
import { NavLink } from "react-router-dom";

const Navbar = () => {
  return (
    <div className="area">
      <div className="navbar">
        <div className="head">
          <NavLink to="/">
            <h2>Navbar</h2>
          </NavLink>
        </div>
        <div className="content">
          <NavLink to="/income">
            <h4>Income</h4>
          </NavLink>
          <NavLink to="/expense">
            <h4>Expense</h4>
          </NavLink>
        </div>
      </div>
    </div>
  );
};

export default Navbar;
