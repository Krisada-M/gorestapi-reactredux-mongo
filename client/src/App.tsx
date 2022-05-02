import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar/Navbar";
import Home from "./pages/Home/Home";
import Income from "./pages/Income/Income";
import Expense from "./pages/Expense/Expense";
import PageNotFound from "./components/PageNotFound/PageNotFound";

import "./App.scss";

function App() {
  return (
    <div className="App">
      <Router>
        <div className="Navbar">
          <Navbar />
        </div>
        <div className="Main">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/income" element={<Income />} />
            <Route path="/expense" element={<Expense />} />
            <Route path="*" element={<PageNotFound />} />
          </Routes>
        </div>
      </Router>
    </div>
  );
}

export default App;
