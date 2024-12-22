import React from "react";
import { useEffect } from "react";
import { useSearchParams } from "react-router-dom";

function VerifyForm() {
  const [searchParams] = useSearchParams();
  const token = searchParams.get("token");

  useEffect(() => {
    const verify = async () => {
      try {
        const response = await fetch(
          `https://localhost/api/v1/auth/verify/${token}`,
          {
            method: "GET",
          }
        );
        const data = await response.json();
        console.log(data);
            //   TODO: EXTEND, ADD LOGIC
      } catch (error) {
        console.error("Email verification error:", error);
      }
    };

    if (token) verify();
  }, [token]);

  return <div>Email Verification in Progress...</div>;
}

export default VerifyForm;
