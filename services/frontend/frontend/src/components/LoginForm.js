import React, { useState } from "react";

const LoginForm = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();

    const requestBody = {
      email,
      password,
    };

    try {
      const response = await fetch("http://localhost/api/v1/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(requestBody),
      });

      const data = await response.json();

      if (response.ok) {
        setMessage(`Login successful! Token: ${data.token}`);
      } else {
        setMessage(`Error ${data.code}: ${data.message}`);
      }
    } catch (error) {
      console.error("Error:", error);
      setMessage("An unexpected error occurred.");
    }
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        maxWidth: "400px",
        margin: "50px auto",
        padding: "20px",
        border: "1px solid #ddd",
        borderRadius: "8px",
        backgroundColor: "#f9f9f9",
        boxShadow: "0 4px 6px rgba(0, 0, 0, 0.1)",
      }}
    >
      <h2
        style={{
          fontSize: "1.5rem",
          fontWeight: "bold",
          marginBottom: "20px",
          color: "#333",
        }}
      >
        Login
      </h2>
      <form
        onSubmit={handleSubmit}
        style={{
          width: "100%",
        }}
      >
        <div style={{ marginBottom: "15px" }}>
          <label
            style={{
              display: "block",
              fontSize: "1rem",
              color: "#555",
              marginBottom: "5px",
            }}
          >
            Email:
          </label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            style={{
              width: "100%",
              padding: "10px",
              fontSize: "1rem",
              border: "1px solid #ccc",
              borderRadius: "4px",
              outline: "none",
              boxSizing: "border-box",
            }}
          />
        </div>
        <div style={{ marginBottom: "15px" }}>
          <label
            style={{
              display: "block",
              fontSize: "1rem",
              color: "#555",
              marginBottom: "5px",
            }}
          >
            Password:
          </label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            style={{
              width: "100%",
              padding: "10px",
              fontSize: "1rem",
              border: "1px solid #ccc",
              borderRadius: "4px",
              outline: "none",
              boxSizing: "border-box",
            }}
          />
        </div>
        <button
          type="submit"
          style={{
            width: "100%",
            padding: "10px",
            fontSize: "1rem",
            color: "#fff",
            backgroundColor: "#007bff",
            border: "none",
            borderRadius: "4px",
            cursor: "pointer",
            marginTop: "10px",
          }}
        >
          Login
        </button>
      </form>
      {message && (
        <p
          style={{
            marginTop: "20px",
            fontSize: "1rem",
            textAlign: "center",
            color: message.includes("successful") ? "#4caf50" : "#f44336",
          }}
        >
          {message}
        </p>
      )}
    </div>
  );
};

export default LoginForm;
