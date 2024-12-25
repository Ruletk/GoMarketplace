import React, {useState} from "react";

function RegForm() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [message, setMessage] = useState(null);
    const [loading, setLoading] = useState(false);

    const handleRegister = async () => {
        setLoading(true);
        setMessage(null);

        try {
            const response = await fetch("http://localhost/api/v1/auth/register", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({email, password}),
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message || "Something went wrong");
            }

            const data = await response.json();
            setMessage(data.message || "Registration successful!");
        } catch (error) {
            setMessage(error.message);
            console.error("Registration error:", error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div style={{maxWidth: "400px", margin: "0 auto", textAlign: "center"}}>
            <h2>Register</h2>
            <input
                type="email"
                placeholder="Enter your email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                style={{
                    padding: "10px",
                    width: "100%",
                    marginBottom: "10px",
                    borderRadius: "5px",
                    border: "1px solid #ccc",
                }}
            />
            <input
                type="password"
                placeholder="Enter your password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                style={{
                    padding: "10px",
                    width: "100%",
                    marginBottom: "10px",
                    borderRadius: "5px",
                    border: "1px solid #ccc",
                }}
            />
            <button
                onClick={handleRegister}
                disabled={loading}
                style={{
                    padding: "10px 20px",
                    borderRadius: "5px",
                    border: "none",
                    background: "#28a745",
                    color: "white",
                    cursor: "pointer",
                    opacity: loading ? 0.6 : 1,
                }}
            >
                {loading ? "Registering..." : "Register"}
            </button>
            {message && <p style={{marginTop: "10px"}}>{message}</p>}
        </div>
    );
}

export default RegForm;
