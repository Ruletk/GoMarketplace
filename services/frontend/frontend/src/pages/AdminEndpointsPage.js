import React, { useState } from 'react';

const AdminEndpointsPage = () => {
    const [responseMessage, setResponseMessage] = useState('');

    const handleValidate = async () => {
        try {
            const response = await fetch('http://localhost/api/v1/auth/validate', {
                method: 'GET',
                credentials: 'include',
            });
            const data = await response.json();
            setResponseMessage(`Validate: ${JSON.stringify(data)}`);
        } catch (error) {
            setResponseMessage(`Error during validation: ${error.message}`);
        }
    };

    const handleHardDeleteSessions = async () => {
        try {
            const response = await fetch('http://localhost/api/v1/auth/admin/sessions/hard-delete', {
                method: 'DELETE',
                credentials: 'include',
            });
            const data = await response.json();
            setResponseMessage(`Hard Delete: ${JSON.stringify(data)}`);
        } catch (error) {
            setResponseMessage(`Error during hard delete: ${error.message}`);
        }
    };

    const handleDeleteInactiveSessions = async () => {
        try {
            const response = await fetch('http://localhost/api/v1/auth/admin/sessions/delete-inactive', {
                method: 'DELETE',
                credentials: 'include',
            });
            const data = await response.json();
            setResponseMessage(`Delete Inactive: ${JSON.stringify(data)}`);
        } catch (error) {
            setResponseMessage(`Error during delete inactive: ${error.message}`);
        }
    };

    return (
        <div style={{ padding: '20px' }}>
            <h2>Admin Endpoints</h2>
            <div style={{ marginBottom: '10px' }}>
                <button onClick={handleValidate} style={buttonStyle}>Validate Token</button>
                <button onClick={handleHardDeleteSessions} style={buttonStyle}>Hard Delete Sessions</button>
                <button onClick={handleDeleteInactiveSessions} style={buttonStyle}>Delete Inactive Sessions</button>
            </div>
            {responseMessage && <p style={messageStyle}>{responseMessage}</p>}
        </div>
    );
};

const buttonStyle = {
    margin: '5px',
    padding: '10px 20px',
    backgroundColor: '#007BFF',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
};

const messageStyle = {
    marginTop: '20px',
    padding: '10px',
    backgroundColor: '#F8F9FA',
    border: '1px solid #CED4DA',
    borderRadius: '5px',
    fontFamily: 'monospace',
};

export default AdminEndpointsPage;
