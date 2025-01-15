import React from 'react';
import { useParams, useNavigate } from 'react-router-dom';

const ProductDetails = () => {
  const { id } = useParams(); // Получаем ID товара из URL
  const navigate = useNavigate(); // Для возврата назад

  // Пример данных (в реальном приложении получите данные из API или хранилища)
  const mockData = [
    { id: 1, title: "Smartphone", category: "Electronics", price: 500 },
    { id: 2, title: "Laptop", category: "Electronics", price: 1000 },
    { id: 3, title: "Shirt", category: "Clothing", price: 50 },
    { id: 4, title: "Shoes", category: "Clothing", price: 80 },
    { id: 5, title: "Makeup Kit", category: "Beauty & Health", price: 30 },
    { id: 6, title: "Dumbbells", category: "Sports", price: 120 },
    { id: 7, title: "Sofa", category: "Furniture", price: 700 },
    { id: 8, title: "Watch", category: "Accessories", price: 200 },
    { id: 9, title: "Headphones", category: "Electronics", price: 150 },
    { id: 10, title: "Jacket", category: "Clothing", price: 100 },
    { id: 11, title: "Perfume", category: "Beauty & Health", price: 60 },
    { id: 12, title: "Basketball", category: "Sports", price: 25 },
    { id: 13, title: "Dining Table", category: "Furniture", price: 800 },
    { id: 14, title: "Necklace", category: "Accessories", price: 300 },
    { id: 15, title: "Tablet", category: "Electronics", price: 400 },
    { id: 16, title: "Trousers", category: "Clothing", price: 70 },
    { id: 17, title: "Lipstick", category: "Beauty & Health", price: 20 },
    { id: 18, title: "Tennis Racket", category: "Sports", price: 150 },
    { id: 19, title: "Bed", category: "Furniture", price: 1000 },
    { id: 20, title: "Bracelet", category: "Accessories", price: 250 },
    { id: 21, title: "Camera", category: "Electronics", price: 600 },
    { id: 22, title: "Sweater", category: "Clothing", price: 90 },
    { id: 23, title: "Face Cream", category: "Beauty & Health", price: 40 },
    { id: 24, title: "Football", category: "Sports", price: 30 },
    { id: 25, title: "Bookshelf", category: "Furniture", price: 150 },
    { id: 26, title: "Earrings", category: "Accessories", price: 100 },
    { id: 27, title: "Monitor", category: "Electronics", price: 300 },
    { id: 28, title: "Dress", category: "Clothing", price: 120 },
    { id: 29, title: "Shampoo", category: "Beauty & Health", price: 15 },
    { id: 30, title: "Yoga Mat", category: "Sports", price: 50 },
    { id: 31, title: "Chair", category: "Furniture", price: 200 },
    { id: 32, title: "Ring", category: "Accessories", price: 400 },
    { id: 33, title: "Television", category: "Electronics", price: 800 },
    { id: 34, title: "Skirt", category: "Clothing", price: 60 },
    { id: 35, title: "Body Lotion", category: "Beauty & Health", price: 25 },
    { id: 36, title: "Golf Clubs", category: "Sports", price: 500 },
    { id: 37, title: "Wardrobe", category: "Furniture", price: 1200 },
    { id: 38, title: "Sunglasses", category: "Accessories", price: 150 },
    { id: 39, title: "Printer", category: "Electronics", price: 200 },
    { id: 40, title: "Blouse", category: "Clothing", price: 70 },
  ];

  // Находим товар по ID
  const product = mockData.find((item) => item.id === parseInt(id, 10));

  // Если товар не найден
  if (!product) {
    return (
      <div>
        <h1>Product not found</h1>
        <button onClick={() => navigate(-1)}>Go Back</button>
      </div>
    );
  }

  return (
    <div style={{ padding: '20px' }}>
      <h1>{product.title}</h1>
      <p><strong>Category:</strong> {product.category}</p>
      <p><strong>Price:</strong> ${product.price}</p>
      <p><strong>Description:</strong> {product.description}</p>
      <button onClick={() => navigate(-1)} style={{ marginTop: '20px', padding: '10px 15px' }}>
        Go Back
      </button>
    </div>
  );
};

export default ProductDetails;
