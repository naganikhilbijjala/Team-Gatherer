import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import Account from './Account';
import reportWebVitals from './reportWebVitals';
import MyDashboard from "./MyDashboard";
import {BrowserRouter, Link, Route, Routes} from 'react-router-dom'
import Homepage from "./Landing";
import CreateGame from "./CreateGame";
import ViewAllGames from "./ViewAllGames";
import GameDetails from "./GameDetails";
import JoinedGames from "./JoinedGames";



const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
      <BrowserRouter>

          <Routes>
              <Route path="/" element={<Homepage />} />
              <Route path="/dashboard" element={<MyDashboard />} />
                <Route path="/login" element={<Account />} />
                <Route path="/create-game" element={<CreateGame />} />
                <Route path="/join-game" element={<ViewAllGames />} />
                <Route path="/gameDetails" element={<GameDetails />} />
                <Route path="/joined-games" element={<JoinedGames />} />

          </Routes>


        </BrowserRouter>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
