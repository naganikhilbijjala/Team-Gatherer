import React, { useEffect, useState } from 'react';
import './MyDashboard.css';
import { FaTrash } from 'react-icons/fa';
import Modal from 'react-modal';



function MyDashboard() {
    const [teams, setTeams] = useState([]);


    const createGame = () => {
        if(!isLoggedIn()) {
            window.location.href = '/login';
        }
        else{
            window.location.href = '/create-game';
        }
    }





    useEffect(() => {
        fetch('http://localhost:8080/teams')
            .then(response => response.json())
            .then(data => setTeams(data))
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

    function handleDelete(ID) {
        console.log(ID);
        fetch('http://localhost:8080/teams/'+ID, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            }
        })
        window.location.href = '/dashboard';

    }

    function goToGameDetails(id) {
        localStorage.setItem('GameAppGameID', id);
        window.location.href = '/gameDetails';

    }

    return (
        <div>

        {/*<ExampleModal  isOpen={true} onRequestClose={() => {}} title="Example Modal" text="This is an example modal." />*/}
            <button onClick={goToHome}>Home</button>

            <div className="teams-container">
                {teams
                    .filter(team => team.Owner === parseInt(localStorage.getItem('GameAppPersonID')))
                    .map(team => (
                        <div className="team-card" key={team.id}>
                            <h2>{team.name}</h2>
                            <br/>

                            {/*<p>{team.id}</p>*/}
                            <p>Time: {team.gameDate}, {team.gameTime}</p>
                            <p>Period: {team.gamePeriod}</p>
                            <button onClick={() => goToGameDetails(team.id)}>Details</button>
                            <br/>
                            <FaTrash className="delete-icon" onClick={() => handleDelete(team.id)} />
                        </div>
                    ))
                }
            </div>
            <div className="team-card">
                <h3>Create New Game</h3>
                <button onClick={createGame}>Create</button>
            </div>
        </div>
    );
}

const MyModal = ({ isOpen, onClose }) => {
    const [text, setText] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        // Do something with the text input
        setText('');
        onClose();
    };

    onClose = () => {
        isOpen = false;
    }



    return (
        <Modal isOpen={isOpen}>
            <h2>My Modal</h2>
            <form onSubmit={handleSubmit}>
                <input type="text" value={text} onChange={(e) => setText(e.target.value)} />
                <button type="submit">Submit</button>
            </form>
            <button onClick={onClose}>Close</button>
        </Modal>
    );
};







export default MyDashboard;