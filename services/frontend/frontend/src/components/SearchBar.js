import React, { useState } from "react";
import axios from "axios";

const ProductSearchBar = () => {
  const [filters, setFilters] = useState({
    category: "",
    company: "",
    pagesize: 10,
    offset: 0,
    minprice: 0,
    maxprice: 1000000000000,
    sort: "asc",
    sortby: "price",
    search: "",
  });

  const [products, setProducts] = useState([]);
  const [totalCount, setTotalCount] = useState(0);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFilters({ ...filters, [name]: value });
  };

  const fetchProducts = async () => {
    try {
      const response = await axios.get("http://localhost/api/v1/product/products", {
        params: filters,
      });
      setProducts(response.data.products);
      setTotalCount(response.data.total_count);
    } catch (error) {
      console.error("Error fetching products:", error);
    }
  };

  return (
    <div className="product-search-bar">
      <div className="filters">
        <input
          type="text"
          name="search"
          placeholder="Search products..."
          value={filters.search}
          onChange={handleInputChange}
        />
        <input
          type="text"
          name="category"
          placeholder="Category IDs (comma-separated)"
          value={filters.category}
          onChange={handleInputChange}
        />
        <input
          type="text"
          name="company"
          placeholder="Company IDs (comma-separated)"
          value={filters.company}
          onChange={handleInputChange}
        />
        <input
          type="number"
          name="minprice"
          placeholder="Min Price"
          value={filters.minprice}
          onChange={handleInputChange}
        />
        <input
          type="number"
          name="maxprice"
          placeholder="Max Price"
          value={filters.maxprice}
          onChange={handleInputChange}
        />
        <select name="sort" value={filters.sort} onChange={handleInputChange}>
          <option value="asc">Ascending</option>
          <option value="desc">Descending</option>
        </select>
        <select name="sortby" value={filters.sortby} onChange={handleInputChange}>
          <option value="price">Price</option>
          <option value="name">Name</option>
          <option value="popularity">Popularity</option>
          <option value="date">Date</option>
        </select>
        <button onClick={fetchProducts}>Search</button>
      </div>

      <div className="results">
        <h3>Total Products: {totalCount}</h3>
        <ul>
          {products.map((product) => (
            <li key={product.id}>
              <h4>{product.name}</h4>
              <p>{product.description}</p>
              <p>Price: ${product.price}</p>
              <p>Category ID: {product.category_id}</p>
              <p>Company ID: {product.company_id}</p>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default ProductSearchBar;