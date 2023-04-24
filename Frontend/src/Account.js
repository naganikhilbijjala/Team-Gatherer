import React, { useState } from 'react';
import './Account.css';


function LoginPage() {
  const [email, setEmail] = useState('');
  const [passcode, setPasscode] = useState('');

  const handleLogin = () => {
    // Send login request to server
    fetch('http://localhost:8080/check', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, passcode }),
    })
        .then(response => response.json())
        .then(data => {
          console.log(data);
          // Handle login response here
            if(data.message == "User exists") {
                console.log("User exists");
                fetch(`http://localhost:8080/getUserInfo?email=${email}`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                })
                    .then(response => response.json())
                    .then(data => {
                        console.log(data);
                        localStorage.setItem('GameAppPersonID', data.id);
                        localStorage.setItem('GameAppPersonName', data.name);
                    });

                localStorage.setItem('GameAppPersonEMail', email);
                window.location.href = '/';
            }
            else{
                alert("User does not exist");
            }
        })
        .catch(error => {
          console.error(error);
          // Handle error here
        });
  };

  return (
      <div>
        <h2>Login</h2>
        <label>Email:</label>
        <input type="text" value={email} onChange={e => setEmail(e.target.value)} />
        <br />
        <label>Password:</label>
        <input type="password" value={passcode} onChange={e => setPasscode(e.target.value)} />
        <br />
        <button onClick={handleLogin}>Login</button>
      </div>
  );
}

function RegisterPage() {
  const [email, setEmail] = useState('');
  const [passcode, setPasscode] = useState('');
    const [name, setName] = useState('');

  const handleRegister = () => {
    // Send registration request to server
    fetch('http://localhost:8080/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({name, email, passcode}),
    })
        .then(response => response.json())
        .then(data => {
            localStorage.setItem('GameAppPersonID', data.id);
          // Handle registration response here
        }).then(data => {
            localStorage.setItem('GameAppPersonEMail', email);
            localStorage.setItem('GameAppPersonName', name);
            window.location.href = '/';
    })
        .catch(error => {
          console.error(error);
          // Handle error here
        });
  };

  return (
      <div>
        <h2>Register</h2>
          <label>Name:</label>
          <input type="text" value={name} onChange={e => setName(e.target.value)} />
          <br />
        <label>Email:</label>
        <input type="text" value={email} onChange={e => setEmail(e.target.value)} />
        <br />
        <label>Password:</label>
        <input type="password" value={passcode} onChange={e => setPasscode(e.target.value)} />
        <br />
        <button onClick={handleRegister}>Register</button>
      </div>
  );
}

function Account() {
  const [activeTab, setActiveTab] = useState('login');

  const handleTabChange = tab => {
    setActiveTab(tab);
  };

    function handleHomeButton() {
        window.location.href = '/';
    }

    return (
      <div>
        <div>
            <button onClick={handleHomeButton}>Home</button>
          <button onClick={() => handleTabChange('login')}>Login</button>
          <button onClick={() => handleTabChange('register')}>Register</button>
        </div>
        {activeTab === 'login' ? <LoginPage /> : <RegisterPage />}
      </div>
  );
}

export default Account;

