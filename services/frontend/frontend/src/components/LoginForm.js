import React, {useState} from 'react';

const LoginForm = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [message, setMessage] = useState('');

    const handleSubmit = async (event) => {
        event.preventDefault();

        const requestBody = {email, password};

        try {
            const response = await fetch('http://localhost/api/v1/auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(requestBody),
            });

            const data = await response.json();

            if (response.ok) {
                setMessage(`Login successful! Token saved in cookie.`);
            } else {
                setMessage(`Error ${data.code}: ${data.message}`);
            }
        } catch (error) {
            console.error('Error:', error);
            setMessage('An unexpected error occurred.');
        }
    };

    return (
        <div style={{
            padding: '20px',
            maxWidth: '400px',
            margin: '0 auto',
            backgroundColor: '#f9f9f9',
            borderRadius: '8px'
        }}>
            <h2 style={{textAlign: 'center', color: '#333'}}>Login</h2>
            <form onSubmit={handleSubmit} style={{display: 'flex', flexDirection: 'column', gap: '15px'}}>
                <div>
                    <label htmlFor="email" style={{fontSize: '14px', color: '#333'}}>Email:</label>
                    <input
                        id="email"
                        type="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                        style={{
                            padding: '10px',
                            fontSize: '14px',
                            border: '1px solid #ccc',
                            borderRadius: '4px',
                            outline: 'none',
                            width: '100%',
                        }}
                    />
                </div>
                <div>
                    <label htmlFor="password" style={{fontSize: '14px', color: '#333'}}>Password:</label>
                    <input
                        id="password"
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                        style={{
                            padding: '10px',
                            fontSize: '14px',
                            border: '1px solid #ccc',
                            borderRadius: '4px',
                            outline: 'none',
                            width: '100%',
                        }}
                    />
                </div>
                <button
                    type="submit"
                    style={{
                        padding: '10px',
                        backgroundColor: '#007BFF',
                        color: '#fff',
                        border: 'none',
                        borderRadius: '4px',
                        fontSize: '16px',
                        cursor: 'pointer',
                        transition: 'background-color 0.3s',
                    }}
                    onMouseOver={(e) => e.target.style.backgroundColor = '#0056b3'}
                    onMouseOut={(e) => e.target.style.backgroundColor = '#007BFF'}
                >
                    Login
                </button>
            </form>
            {message && <p style={{color: '#d9534f', textAlign: 'center', marginTop: '20px'}}>{message}</p>}
        </div>
    );
};

export default LoginForm;
