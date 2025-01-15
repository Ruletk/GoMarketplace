import React, { useState, useEffect } from 'react';
import ReactPaginate from 'react-paginate';
import '../styles/Pagination.css'; // Подключаем стили для пагинации

const Pagination = () => {
  const [products, setProducts] = useState([]);
  const [filteredProducts, setFilteredProducts] = useState([]);
  const [currentPage, setCurrentPage] = useState(0);
  const [itemsPerPage] = useState(10);
  const [selectedCategory, setSelectedCategory] = useState('');
  const [categoryLocked, setCategoryLocked] = useState(false);
  const [minPrice, setMinPrice] = useState('');
  const [maxPrice, setMaxPrice] = useState('');
  const [sortType, setSortType] = useState(''); // Добавляем состояние для сортировки

  // Пример данных
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

  useEffect(() => {
    setProducts(mockData);
    setFilteredProducts(mockData);
  }, []);

  // Обработка смены категории
  const handleCategoryChange = (category) => {
    setSelectedCategory(category);
    setCategoryLocked(true);
    const filtered = category ? products.filter((product) => product.category === category) : products;
    setFilteredProducts(filtered);
    setCurrentPage(0); // Сбрасываем пагинацию на первую страницу
  };

  // Обработка смены страницы
  const handlePageClick = (event) => {
    const newPage = event.selected;
    setCurrentPage(newPage);
  };

  // Обработка фильтрации по цене
  const handlePriceFilter = () => {
    const filtered = products.filter((product) => {
      return (
        (!selectedCategory || product.category === selectedCategory) &&
        product.price >= (minPrice || 0) &&
        product.price <= (maxPrice || Infinity)
      );
    });
    setFilteredProducts(filtered);
    setCurrentPage(0); // Сбрасываем пагинацию на первую страницу
  };

  // Обработка сортировки
  const handleSortChange = (sortType) => {
    setSortType(sortType);
    let sortedProducts = [...filteredProducts];
    if (sortType === 'price-asc') {
      sortedProducts.sort((a, b) => a.price - b.price);
    } else if (sortType === 'price-desc') {
      sortedProducts.sort((a, b) => b.price - a.price);
    } else if (sortType === 'title-asc') {
      sortedProducts.sort((a, b) => a.title.localeCompare(b.title));
    } else if (sortType === 'title-desc') {
      sortedProducts.sort((a, b) => b.title.localeCompare(a.title));
    }
    setFilteredProducts(sortedProducts);
  };

  // Сброс фильтров
  const resetFilters = () => {
    setSelectedCategory('');
    setCategoryLocked(false);
    setMinPrice('');
    setMaxPrice('');
    setSortType('');
    setFilteredProducts(products);
    setCurrentPage(0); // Сбрасываем пагинацию на первую страницу
  };

  // Получаем текущие элементы для отображения
  const offset = currentPage * itemsPerPage;
  const currentItems = filteredProducts.slice(offset, offset + itemsPerPage);

  return (
    <div>
      <div className="category-filter">
        <button onClick={() => handleCategoryChange('')} disabled={categoryLocked}>All</button>
        <button onClick={() => handleCategoryChange('Electronics')} disabled={categoryLocked}>Electronics</button>
        <button onClick={() => handleCategoryChange('Clothing')} disabled={categoryLocked}>Clothing</button>
        <button onClick={() => handleCategoryChange('Beauty & Health')} disabled={categoryLocked}>Beauty & Health</button>
        <button onClick={() => handleCategoryChange('Sports')} disabled={categoryLocked}>Sports</button>
        <button onClick={() => handleCategoryChange('Furniture')} disabled={categoryLocked}>Furniture</button>
        <button onClick={() => handleCategoryChange('Accessories')} disabled={categoryLocked}>Accessories</button>
      </div>

      {categoryLocked && (
        <div className="filter-controls">
          <div className="price-filter">
            <input
              type="number"
              placeholder="Min price"
              value={minPrice}
              onChange={(e) => setMinPrice(e.target.value)}
            />
            <input
              type="number"
              placeholder="Max price"
              value={maxPrice}
              onChange={(e) => setMaxPrice(e.target.value)}
            />
            <button onClick={handlePriceFilter}>Apply Price Filter</button>
          </div>
          <div className="sort-filter">
            <select value={sortType} onChange={(e) => handleSortChange(e.target.value)}>
              <option value="">Sort By</option>
              <option value="price-asc">Price: Low to High</option>
              <option value="price-desc">Price: High to Low</option>
              <option value="title-asc">Title: A to Z</option>
              <option value="title-desc">Title: Z to A</option>
            </select>
          </div>
          <button onClick={resetFilters}>Reset Filters</button>
        </div>
      )}

      <div className="items-list">
        {currentItems.map((item) => (
          <div key={item.id} className="item-card">
            <h3>{item.title}</h3>
            <p>Category: {item.category}</p>
            <p>Price: ${item.price}</p>
          </div>
        ))}
      </div>

      <ReactPaginate
        previousLabel={'< previous'}
        nextLabel={'next >'}
        breakLabel={'...'}
        breakClassName={'break-me'}
        pageCount={Math.ceil(filteredProducts.length / itemsPerPage)}
        marginPagesDisplayed={2}
        pageRangeDisplayed={5}
        onPageChange={handlePageClick}
        containerClassName={'pagination'}
        activeClassName={'active'}
      />
    </div>
  );
};

export default Pagination;