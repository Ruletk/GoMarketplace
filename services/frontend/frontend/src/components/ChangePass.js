import React, { useState } from "react";

function ChangePass() {
  const [email, setEmail] = useState("");
  const [message, setMessage] = useState(null);
  const [loading, setLoading] = useState(false);

  const handlePasswordChangeRequest = async () => {
    setLoading(true);
    setMessage(null);

    try {
      const response = await fetch(
        "http://localhost/api/v1/auth/change-password",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ email }),
        }
      );

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || "Something went wrong");
      }

      const data = await response.json();
      setMessage(data.message || "Password change request sent successfully!");
    } catch (error) {
      setMessage(error.message);
      console.error("Password change request error:", error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ maxWidth: "400px", margin: "0 auto", textAlign: "center" }}>
      <h2>Change Password</h2>
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
      <button
        onClick={handlePasswordChangeRequest}
        disabled={loading}
        style={{
          padding: "10px 20px",
          borderRadius: "5px",
          border: "none",
          background: "#007BFF",
          color: "white",
          cursor: "pointer",
          opacity: loading ? 0.6 : 1,
        }}
      >
        {loading ? "Processing..." : "Change Password"}
      </button>
      {message && <p style={{ marginTop: "10px" }}>{message}</p>}
    </div>
  );
}

export default ChangePass;
