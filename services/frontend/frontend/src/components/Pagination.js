import React, { useState, useEffect } from "react";
import ReactPaginate from "react-paginate";
import "../styles/Pagination.css";

const PaginatedList = () => {
  const [items, setItems] = useState([]);
  const [filteredItems, setFilteredItems] = useState([]);
  const [currentPage, setCurrentPage] = useState(0);
  const [filters, setFilters] = useState({
    sort: "asc", // asc, desc, title
    minPrice: 0,
    maxPrice: Infinity,
  });

  const itemsPerPage = 10;

  // Временные данные
  const mockData = [
    { id: 1, title: "Item A", price: 10 },
    { id: 2, title: "Item B", price: 20 },
    { id: 3, title: "Item C", price: 15 },
    { id: 4, title: "Item D", price: 30 },
    { id: 5, title: "Item E", price: 25 },
    { id: 6, title: "Item F", price: 40 },
    { id: 7, title: "Item G", price: 5 },
    { id: 8, title: "Item H", price: 50 },
    { id: 9, title: "Item I", price: 35 },
    { id: 10, title: "Item J", price: 45 },
    { id: 11, title: "Item K", price: 60 },
    { id: 12, title: "Item L", price: 70 },
    { id: 13, title: "Item M", price: 80 },
    { id: 14, title: "Item N", price: 90 },
    { id: 15, title: "Item O", price: 100 },
    { id: 16, title: "Item P", price: 110 },
    { id: 17, title: "Item Q", price: 120 },
    { id: 18, title: "Item R", price: 130 },
    { id: 19, title: "Item S", price: 140 },
    { id: 20, title: "Item T", price: 150 },
    { id: 21, title: "Item U", price: 160 },
    { id: 22, title: "Item V", price: 170 },
    { id: 23, title: "Item W", price: 180 },
    { id: 24, title: "Item X", price: 190 },
    { id: 25, title: "Item Y", price: 200 },
  ];

  useEffect(() => {
    // Устанавливаем временные данные вместо запроса
    setItems(mockData);
    setFilteredItems(mockData);
  }, []);

  // Применение фильтров
  useEffect(() => {
    let updatedItems = [...items];

    // Фильтр по цене
    updatedItems = updatedItems.filter(
      (item) =>
        item.price >= filters.minPrice && item.price <= filters.maxPrice
    );

    // Сортировка
    if (filters.sort === "asc") {
      updatedItems.sort((a, b) => a.price - b.price);
    } else if (filters.sort === "desc") {
      updatedItems.sort((a, b) => b.price - a.price);
    } else if (filters.sort === "title") {
      updatedItems.sort((a, b) => a.title.localeCompare(b.title));
    }

    setFilteredItems(updatedItems);
  }, [filters, items]);

  // Обработка смены страницы
  const handlePageClick = (event) => {
    setCurrentPage(event.selected);
  };

  // Отображаемые элементы
  const offset = currentPage * itemsPerPage;
  const currentItems = filteredItems.slice(offset, offset + itemsPerPage);

  // Обновление фильтров
  const updateFilters = (newFilters) => {
    setFilters((prev) => ({ ...prev, ...newFilters }));
    setCurrentPage(0); // Сброс на первую страницу
  };

  return (
    <div>
      <h1>Filter</h1>

      {/* Фильтры */}
      <div>
        <label>
          Sort:
          <select
            value={filters.sort}
            onChange={(e) => updateFilters({ sort: e.target.value })}
          >
            <option value="asc">Ascending</option>
            <option value="desc">Discending</option>
            <option value="title">by name</option>
          </select>
        </label>
        <label>
          Min price:
          <input
            type="number"
            value={filters.minPrice}
            onChange={(e) =>
              updateFilters({ minPrice: Number(e.target.value) })
            }
          />
        </label>
        <label>
          Max price:
          <input
            type="number"
            value={filters.maxPrice}
            onChange={(e) =>
              updateFilters({ maxPrice: Number(e.target.value) })
            }
          />
        </label>
      </div>

      {/* Список элементов */}
      <ul>
        {currentItems.map((item) => (
          <li key={item.id}>
            {item.title} - ${item.price}
          </li>
        ))}
      </ul>

      {/* Пагинация */}
      <ReactPaginate
        previousLabel={"Предыдущая"}
        nextLabel={"Следующая"}
        breakLabel={"..."}
        pageCount={Math.ceil(filteredItems.length / itemsPerPage)}
        marginPagesDisplayed={2}
        pageRangeDisplayed={5}
        onPageChange={handlePageClick}
        containerClassName={"pagination"}
        activeClassName={"active"}
      />
    </div>
  );
};

export default PaginatedList;
