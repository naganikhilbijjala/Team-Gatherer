import React, { useEffect, useState } from 'react';
import './ViewAllGames.css';

function ViewAllGames() {
    const [teams, setTeams] = useState([]);
    const [name, setName] = useState(localStorage.getItem('GameAppPersonName'));
    const [ ID, setID] = useState(localStorage.getItem('GameAppPersonID'));
    const [UserJoinedteams, setUserJoinedteams] = useState([]);


    const createGame = () => {
        if(!isLoggedIn()) {
            window.location.href = '/login';
        }
        else{
            window.location.href = '/create-game';
        }
    }

    useEffect(() => {
        if(!isLoggedIn()) {
            window.location.href = '/login';
        }
    });

    useEffect(() => {
        fetch(`http://localhost:8080/getTeamsByUser?user_id=${localStorage.getItem('GameAppPersonID')}`)
            .then(response => response.json())
            .then(data => setUserJoinedteams(data.teams))
            .then(data => console.log(data))
            .catch(error => console.error(error));
    }, []);

    useEffect(() => {
        fetch('http://localhost:8080/teams')
            .then(response => response.json())
            .then(data => setTeams(data))
            .then(data => console.log(data))
            .catch(error => console.error(error));
    }, []);

    function goToHome() {
        window.location.href = '/';
    }

    function isLoggedIn() {
        if(localStorage.getItem('GameAppPersonID') != null) {
            return true;
        }
        else {
            return false;
        }
    }

    function joinGame(GameID) {
        console.log(GameID);
        if(!isLoggedIn()) {
            window.location.href = '/login';
        }
        else{
            fetch('http://localhost:8080/players', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({name, "team_id": GameID, "user_id": parseInt(ID)}),
            }).then(window.location.href = '/join-game');
         }
        return undefined;
    }

    function leaveGame(GameID) {
        console.log(GameID);
        if (!isLoggedIn()) {
            window.location.href = '/login';
        }
        else {
            fetch(`http://localhost:8080/leave-game`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({"team_id": GameID, "user_id": parseInt(ID)}),
            })
                .then(response => {
                    if (response.ok) {
                        window.location.href = '/join-game';
                    } else {
                        throw new Error('Failed to leave game');
                    }
                })
                .catch(error => console.error(error));
        }
    }

    return (
        <div>
            {/*<ExampleModal  isOpen={true} onRequestClose={() => {}} title="Example Modal" text="This is an example modal." />*/}
            <button onClick={goToHome}>Home</button>


            <div className="teams-container">
                {teams.map(team => (
                        <div className="team-card" key={team.id}>
                            <h2>{team.name}</h2>
                            {/*<p>{team.id}</p>*/}
                            <p>{team.gameTime}</p>
                            <p>{team.gameDate}</p>
                            <p>{team.gamePeriod}</p>
                            <p>{team.current}/{team.max} Filled</p>



                            {UserJoinedteams && UserJoinedteams.includes(team.id) ?
                                <button style={{ color: 'black', backgroundColor: 'darkgrey' }} onClick={() => leaveGame(team.id)}>Leave</button>
                                : (team.current >= team.max ?
                                        <button style={{ color: 'white', backgroundColor: 'red' }} disabled>Full</button>
                                        : <button onClick={() => joinGame(team.id)}>Join</button>
                                )
                            }

                        </div>
                    ))
                }
                <div className="team-card">
                    <h3>Create New Game</h3>
                    <button onClick={createGame}>Create</button>
                </div>
            </div>
        </div>
    );
}

export default ViewAllGames;