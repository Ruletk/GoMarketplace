import React from "react";

const products = [
    { id: 1, name: "Product 1", price: "$10", image: "https://via.placeholder.com/150" },
    { id: 2, name: "Product 2", price: "$15", image: "https://via.placeholder.com/150" },
    { id: 3, name: "Product 3", price: "$20", image: "https://via.placeholder.com/150" },
    { id: 4, name: "Product 4", price: "$25", image: "https://via.placeholder.com/150" },
    { id: 5, name: "Product 5", price: "$35", image: "https://via.placeholder.com/150" },
    { id: 6, name: "Product 6", price: "$15", image: "https://via.placeholder.com/150" },
    { id: 7, name: "Product 7", price: "$65", image: "https://via.placeholder.com/150" },
    { id: 8, name: "Product 8", price: "$5", image: "https://via.placeholder.com/150" },

];

function HomePage() {
    return (
        <div style={{ padding: "20px", textAlign: "center" }}>
            <h1>Welcome to Our Marketplace</h1>
            <div
                style={{
                    display: "grid",
                    gridTemplateColumns: "repeat(auto-fit, minmax(200px, 1fr))",
                    gap: "20px",
                    marginTop: "20px",
                }}
            >
                {products.map((product) => (
                    <div
                        key={product.id}
                        style={{
                            border: "1px solid #ddd",
                            borderRadius: "5px",
                            padding: "10px",
                            textAlign: "center",
                            background: "#f9f9f9",
                        }}
                    >
                        <img
                            src={product.image}
                            alt={product.name}
                            style={{ width: "100%", height: "150px", objectFit: "cover", borderRadius: "5px" }}
                        />
                        <h2 style={{ fontSize: "1.2em", margin: "10px 0" }}>{product.name}</h2>
                        <p style={{ fontSize: "1em", fontWeight: "bold", color: "#28a745" }}>{product.price}</p>
                        <button
                            style={{
                                padding: "10px 15px",
                                border: "none",
                                background: "#007bff",
                                color: "#fff",
                                borderRadius: "5px",
                                cursor: "pointer",
                            }}
                        >
                            Add to Cart
                        </button>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default HomePage;
