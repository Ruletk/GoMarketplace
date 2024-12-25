import React, {useEffect, useState} from "react";
import {useSearchParams} from "react-router-dom";

function VerifyForm() {
    const [searchParams] = useSearchParams();
    const token = searchParams.get("token");
    const [status, setStatus] = useState("Verifying...");
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const verify = async () => {
            if (!token) {
                setStatus("Invalid or missing token.");
                setLoading(false);
                return;
            }

            try {
                const response = await fetch(
                    `http://localhost/api/v1/auth/verify/${token}`,
                    {
                        method: "GET",
                    }
                );

                if (!response.ok) {
                    const errorData = await response.json();
                    throw new Error(errorData.message || "Verification failed");
                }

                const data = await response.json();
                setStatus(data.message || "Email successfully verified!");
            } catch (error) {
                setStatus(error.message);
                console.error("Email verification error:", error);
            } finally {
                setLoading(false);
            }
        };

        verify();
    }, [token]);

    return (
        <div style={{maxWidth: "400px", margin: "50px auto", textAlign: "center"}}>
            <h2>Email Verification</h2>
            {loading ? (
                <p>Verifying your email...</p>
            ) : (
                <p style={{color: status.includes("successfully") ? "green" : "red"}}>
                    {status}
                </p>
            )}
        </div>
    );
}

export default VerifyForm;
