import React, {useState} from 'react';
import { Link } from 'react-router-dom';
import App from "./Account";

function Homepage() {

    function isLoggedIn() {
        if(localStorage.getItem('GameAppPersonID') != null) {
            return true;
        }
        else {
            return false;
        }
    }


    function goToCreateGame() {
        if(!isLoggedIn()) {
            window.location.href = '/login';
        }
        else{
            window.location.href = '/create-game';

        }
    }
    function goToJoinGame() {
        window.location.href = '/join-game';

    }
    function goToMyDashboard() {

        if(!isLoggedIn()) {
            window.location.href = '/login';
        }
        else{
            window.location.href = '/dashboard';
        }
    }

    function goToJoinedGames() {

            if(!isLoggedIn()) {
                window.location.href = '/login';
            }
            else{
                window.location.href = '/joined-games';
            }
    }


    function logout() {
        localStorage.removeItem('GameAppPersonEMail');
        localStorage.removeItem('GameAppPersonID');
        localStorage.removeItem('GameAppPersonName');
        window.location.href = '/';
    }

    function goToLogin() {
        window.location.href = '/login';
    }

    return (
        <div>
        <div>

            <h1>Welcome to the Group Sports Game App!</h1>
            <p>Plan and join group sports games with ease.</p>
                <button onClick={goToCreateGame}>Create a Game</button>
                <button onClick={goToJoinGame}>Join a Game</button>
                <button onClick={goToMyDashboard}>View My Dashboard</button>
                <button onClick={goToJoinedGames}>Joined Games</button>
            {isLoggedIn() &&
                <div>
                <p>Logged in as {localStorage.getItem('GameAppPersonName')}</p>
                <button onClick={logout}>Logout</button>
                </div>
            }
            {!isLoggedIn() &&
                <div>
                <p>Not logged in</p>
                <button onClick={goToLogin}>Login</button>
                </div>
            }


        </div>

        </div>
    );
}

export default Homepage;