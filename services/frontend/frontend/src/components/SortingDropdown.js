import React from "react";

const SortingDropdown = ({ onSortChange }) => {
  return (
    <div style={{ marginBottom: "20px" }}>
      <label>
        Сортировка:
        <select
          onChange={(e) => onSortChange(e.target.value)}
          style={{ marginLeft: "10px" }}
        >
          <option value="priceAsc">Цена: по возрастанию</option>
          <option value="priceDesc">Цена: по убыванию</option>
          <option value="titleAsc">Название: от А до Я</option>
          <option value="titleDesc">Название: от Я до А</option>
        </select>
      </label>
    </div>
  );
};

export default SortingDropdown;
