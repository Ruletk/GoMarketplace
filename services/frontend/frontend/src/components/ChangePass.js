import React from "react";
import { useState } from "react";

function ChangePass() {
  const [email, setEmail] = useState("");

  const handlePasswordChangeRequest = async () => {
    try {
      const response = await fetch(
        "https://localhost/api/v1/auth/change-password",
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ email }),
        }
      );
      const data = await response.json();
      console.log(data);
      //   TODO: EXTEND, ADD LOGIC
    } catch (error) {
      console.error("Password change request error:", error);
    }
  };

  return (
    <div>
      <input
        type="email"
        placeholder="Email"
        onChange={(e) => setEmail(e.target.value)}
      />
      <button onClick={handlePasswordChangeRequest}>Change Password</button>
    </div>
  );
}

export default ChangePass;
