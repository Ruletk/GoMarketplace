import React, { useEffect } from 'react';

const Validate = () => {
    useEffect(() => {
        const validateToken = async () => {
            try {
                const response = await fetchWithAuth('http://localhost/api/v1/auth/validate', {
                    method: 'GET',
                });
                const data = await response.json();
                console.log(data);
            } catch (error) {
                console.error('Validation error:', error);
            }
        };

        validateToken();
    }, []);

    return <div>Validating...</div>;
};

export default Validate;
