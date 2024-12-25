import React, { useState } from "react";


const LogoutButton = () => {
    const [token, setToken] = useState('');
    const [message, setMessage] = useState('');

    const handleLogout = async (event) => {
        event.preventDefault();

        try {
            const response = await fetch('http://localhost/api/v1/auth/logout', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
            });

            const data = await response.json();

            if (response.ok) {
                setMessage(data.message);
            } else {
                setMessage(`Error ${data.code}: ${data.message}`);
            }
        } catch (error) {
            console.error('Logout error:', error);
            setMessage('An unexpected error occurred.');
        }
    }

    return (
        <div>
            <button onClick={handleLogout} style={{ padding: '10px', fontSize: '16px', cursor: 'pointer' }}>Logout</button>
            <p style={{ color: message.includes('successfully') ? 'green' : 'red' }}>{message}</p>
        </div>
    );
}

export default LogoutButton;