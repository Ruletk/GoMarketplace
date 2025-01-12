import React from "react";

const ProductCategories = ({ onCategoryChange }) => {
  const categories = [
    { id: 1, name: "Electronics" },
    { id: 2, name: "Clothing" },
    { id: 3, name: "Beauty & Health" },
    { id: 4, name: "Sports" },
    { id: 5, name: "Furniture" },
  ];

  return (
    <div style={{ marginRight: "20px" }}>
      <h3>Categories</h3>
      <ul style={{ listStyle: "none", padding: 0 }}>
        {categories.map((category) => (
          <li
            key={category.id}
            style={{
              cursor: "pointer",
              marginBottom: "10px",
              color: "blue",
              textDecoration: "underline",
            }}
            onClick={() => onCategoryChange(category.name)}
          >
            {category.name}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ProductCategories;
