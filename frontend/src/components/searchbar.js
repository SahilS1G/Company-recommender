import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

function SearchBar() {
  const [companyName, setCompanyName] = useState("");
  const navigate = useNavigate();

  const handleSubmit = (event) => {
    event.preventDefault(); // Prevent the default form submission behavior

    // Send an HTTP request to the Golang backend with the user's input
    fetch("http://localhost:4000/search", {
      method: "POST",
      body: JSON.stringify({ companyname: companyName }),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => response.json())
      .then((data) => {
        // Handle the response from the Golang backend
        // (e.g., update UI, display results, etc.)
        console.log(data);
        // Navigate to "/company" after successful form submission
        navigate("/company");
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  };

  const handleChange = (event) => {
    setCompanyName(event.target.value); // Update the state with the user's input
  };

  return (
    <div style={searchBarContainerStyle}>
      <form style={searchBarStyle} className="search-bar" onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Company Name"
          style={searchInputStyle}
          value={companyName}
          onChange={handleChange}
        />
        <button style={searchButtonStyle} type="submit">
          Search
        </button>
      </form>
    </div>
  );
}

// Styles
const searchBarContainerStyle = {
  backgroundColor: "#333",
  display: "flex",
  justifyContent: "center",
  alignItems: "center",
  height: "92.5vh",
};

const searchBarStyle = {
  display: "flex",
  alignItems: "center",
  width: "70%",
  maxWidth: "500px",
  height: "50px",
  backgroundColor: "#222",
  borderRadius: "25px",
  boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
};

const searchInputStyle = {
  flex: 1,
  border: "none",
  outline: "none",
  padding: "10px",
  fontSize: "16px",
  borderRadius: "25px",
  color: "#fff",
  backgroundColor: "transparent",
};

const searchButtonStyle = {
  backgroundColor: "#fff",
  color: "#333",
  border: "none",
  outline: "none",
  padding: "10px 20px",
  fontSize: "16px",
  borderRadius: "25px",
  cursor: "pointer",
};

export default SearchBar;
