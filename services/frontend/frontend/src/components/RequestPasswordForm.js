import React from "react";
import { useState } from "react";

function RequestPasswordForm() {
  const [newPassword, setNewPassword] = useState("");
  const token = "your_token_here"; // Получите токен из URL или другого источника

  const handleResetPassword = async () => {
    try {
      const response = await fetch(
        `https://localhost/api/v1/auth/change-password/${token}`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ newPassword }),
        }
      );
      const data = await response.json();
      console.log(data);
    //   TODO: EXTEND, ADD LOGIC
    } catch (error) {
      console.error("Reset password error:", error);
    }
  };

  return (
    <div>
      <input
        type="password"
        placeholder="New Password"
        onChange={(e) => setNewPassword(e.target.value)}
      />
      <button onClick={handleResetPassword}>Reset Password</button>
    </div>
  );
}

export default RequestPasswordForm;
