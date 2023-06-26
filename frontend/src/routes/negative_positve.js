import React, { useState, useEffect } from "react";

function NegativePositive() {
  const [articles, setArticles] = useState([]);
  const [articlesNeg, setArticlesNeg] = useState([]);
  const [showPositive, setShowPositive] = useState(true);
  const [showNegative, setShowNegative] = useState(false);

  async function fetchArticles() {
    try {
      const res = await fetch("http://localhost:4000/getPositive");
      const data = await res.json();
      setArticles(data);
      setShowPositive(true);
      setShowNegative(false);
    } catch (error) {
      console.log("Error fetching articles:", error);
    }
  }

  useEffect(() => {
    console.log("Updated articles:", articles);
  }, [articles]);

  async function fetchNegativeArticles() {
    try {
      const res = await fetch("http://localhost:4000/getNegative");
      const data = await res.json();
      setArticlesNeg(data);
      setShowPositive(false);
      setShowNegative(true);
    } catch (error) {
      console.log("Error fetching articles:", error);
    }
  }

  useEffect(() => {
    console.log("Updated articles:", articlesNeg);
  }, [articlesNeg]);

  const containerStyle = {
    backgroundColor: "#333",
    color: "#fff",
  };

  const tabStyle = {
    cursor: "pointer",
    padding: "10px 20px",
    backgroundColor: "#eee",
    marginRight: "5px",
  };

  const listItemStyle = {
    marginBottom: "10px",
  };

  const linkStyle = {
    color: "#fff",
    textDecoration: "none",
  };

  useEffect(() => {
    fetchArticles();
    fetchNegativeArticles();
  }, []);

  return (
    <div style={containerStyle}>
      <h1>Stock Chart</h1>
      <div style={{ display: "flex", marginBottom: "10px" }}>
        <div
          style={{
            ...tabStyle,
            backgroundColor: showPositive ? "#eee" : "#333",
          }}
          onClick={fetchArticles}
        >
          Get Positive News
        </div>
        <div
          style={{
            ...tabStyle,
            backgroundColor: showNegative ? "#eee" : "#333",
          }}
          onClick={fetchNegativeArticles}
        >
          Get Negative News
        </div>
      </div>
      {showPositive && (
        <ul>
          {articles.length > 0 ? (
            articles.map((article) => (
              <li key={article.title} style={listItemStyle}>
                <h2>{article.title}</h2>
                <p>{article.description}</p>
                <a href={article.url} style={linkStyle}>
                  link
                </a>
              </li>
            ))
          ) : (
            <p>No positive articles found.</p>
          )}
        </ul>
      )}
      {showNegative && (
        <ul>
          {articlesNeg.length > 0 ? (
            articlesNeg.map((article) => (
              <li key={article.title} style={listItemStyle}>
                <h2>{article.title}</h2>
                <p>{article.description}</p>
                <a href={article.url} style={linkStyle}>
                  link
                </a>
              </li>
            ))
          ) : (
            <p>No negative articles found.</p>
          )}
        </ul>
      )}
    </div>
  );
}

export default NegativePositive;
