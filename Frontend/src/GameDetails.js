import React, { useState, useEffect } from 'react';
import './GameDetails.css';

const GameDetails = ({ match }) => {
    const [gameId, setGameId] = useState(localStorage.getItem('GameAppGameID'));
    const [game, setGame] = useState({});
    const [players, setPlayers] = useState([]);
    const [isEditable, setIsEditable] = useState(false); // Add state variable

    useEffect(() => {
        fetch(`http://localhost:8080/getGameWithPlayers?id=${gameId}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        })
            .then(response => response.json())
            .then(data => {
                setGame(data.game);
                setPlayers(data.players);
            })
            .catch(error => console.log(error))
    }, [gameId]);

    function goToHome() {
        window.location.href = '/dashboard';
    }

    function toggleEditMode() { // Add function to toggle edit mode
        setIsEditable(!isEditable);
    }

    function handleInputChange(event) { // Add function to handle input change
        const { name, value } = event.target;
        setGame({ ...game, [name]: value });
        fetch('http://localhost:8080/teams'+gameId, {
            method: 'UPDATE',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({name, }),
        })
        // console.log(game);
    }

    return (
        <div>
            <button onClick={goToHome}>Back</button>
            <button onClick={toggleEditMode}>{isEditable ? 'Save' : 'Edit'}</button> {/* Add edit button */}
            <div className="container">
                <h1>Game Details</h1>
                {isEditable ? ( // Render editable inputs if in edit mode
                    <div>
                        <label>Name</label>
                        <input type="text" name="name" value={game.name} onChange={handleInputChange} />
                        <label>Date</label>
                        <input type="date" name="gameDate" value={game.gameDate} onChange={handleInputChange} />
                        <label>Time</label>
                        <input type="time" name="gameTime" value={game.gameTime} onChange={handleInputChange} />
                        <label>Time Period</label>
                        {/*<input type="text" name="gamePeriod" value={game.gamePeriod} onChange={handleInputChange} />*/}
                        <select value={game.gamePeriod} onChange={handleInputChange}>
                            <option value="">Select game period</option>
                            <option value="30 mins">30 mins</option>
                            <option value="1 hour">1 hour</option>
                            <option value="1.5 hours">1.5 hours</option>
                            <option value="2 hours">2 hours</option>
                            <option value="2.5 hours">2.5 hours</option>
                            <option value="3 hours">3 hours</option>
                            <option value="3.5 hours">3.5 hours</option>
                            <option value="4 hours">4 hours</option>
                            <option value="4.5 hours">4.5 hours</option>
                            <option value="5 hours">5 hours</option>
                            <option value="5.5 hours">5.5 hours</option>
                            <option value="6 hours">6 hours</option>
                            <option value="6.5 hours">6.5 hours</option>
                            <option value="7 hours">7 hours</option>
                            <option value="7.5 hours">7.5 hours</option>
                            <option value="8 hours">8 hours</option>
                        </select>
                    </div>
                ) : ( // Render non-editable text if not in edit mode
                    <div>
                        <h2>Name: {game.name}</h2>
                        <p>Date: {game.gameDate}</p>
                        <p>Time: {game.gameTime}</p>
                        <p>Time Period: {game.gamePeriod}</p>
                    </div>
                )}
                <br/>
                <br/>
                <h3>Players</h3>
                <ul>
                    {players.map(player => (
                        <li key={player.id}>  {player.name}</li>
                    ))}
                </ul>
            </div>
        </div>
    );
};

export default GameDetails;
