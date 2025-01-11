import React from "react";
import Cards from "../components/Cards"; // Adjust the path as necessary
import PaginatedItems from "../components/Pagination";

function Home() {
    return (
        <div>
            <Cards/>
            <PaginatedItems itemsPerPage={4} />
        </div>
    );
}

export default Home;