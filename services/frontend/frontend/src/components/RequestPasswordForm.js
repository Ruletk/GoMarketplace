import React, {useState} from "react";

function RequestPasswordForm() {
    const [newPassword, setNewPassword] = useState("");
    const [message, setMessage] = useState(null);
    const [loading, setLoading] = useState(false);

    const token = "your_token_here"; // Получите токен из URL или другого источника

    const handleResetPassword = async () => {
        if (!newPassword) {
            setMessage("Password cannot be empty!");
            return;
        }

        setLoading(true);
        setMessage(null);

        try {
            const response = await fetch(
                `http://localhost/api/v1/auth/change-password/${token}`,
                {
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
                    body: JSON.stringify({newPassword}),
                }
            );

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message || "Something went wrong");
            }

            const data = await response.json();
            setMessage(data.message || "Password reset successful!");
        } catch (error) {
            setMessage(error.message);
            console.error("Reset password error:", error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div style={{maxWidth: "400px", margin: "0 auto", textAlign: "center"}}>
            <h2>Reset Password</h2>
            <input
                type="password"
                placeholder="Enter new password"
                value={newPassword}
                onChange={(e) => setNewPassword(e.target.value)}
                style={{
                    padding: "10px",
                    width: "100%",
                    marginBottom: "10px",
                    borderRadius: "5px",
                    border: "1px solid #ccc",
                }}
            />
            <button
                onClick={handleResetPassword}
                disabled={loading}
                style={{
                    padding: "10px 20px",
                    borderRadius: "5px",
                    border: "none",
                    background: "#007bff",
                    color: "white",
                    cursor: "pointer",
                    opacity: loading ? 0.6 : 1,
                }}
            >
                {loading ? "Resetting..." : "Reset Password"}
            </button>
            {message && <p style={{marginTop: "10px"}}>{message}</p>}
        </div>
    );
}

export default RequestPasswordForm;
