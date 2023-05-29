import React, { useState, useEffect } from "react";

const r = fetch("http://localhost:4000/company")

function NegativePositive() {
  const [articles, setArticles] = useState([]);
  const [articlesNeg, setArticlesNeg] = useState([]);

  async function fetchArticles() {
    try {
      const res = await fetch("http://localhost:4000/getPositive");
      const data = await res.json();
      setArticles(data);
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
    } catch (error) {
      console.log("Error fetching articles:", error);
    }
  }

  useEffect(() => {
    console.log("Updated articles:", articles);
  }, [articlesNeg]);

  return (
    <div>
      <h1>Positive News</h1>
      <button onClick={fetchArticles}>Get Positive News</button>
      <ul>
        {console.log(articles)}
        {articles.map((article) => (
            <li key={article.title}>
              <h2>{article.title}</h2>
              <p>{article.description}</p>
              <p>{article.url}</p>
            </li>
          ))
         
        }
      </ul>
      <h1>Negative News</h1>
      <button onClick={fetchNegativeArticles}>Get Negative News</button>
      <ul>
        {console.log(articles)}
        {articlesNeg.map((article) => (
            <li key={article.title}>
              <h2>{article.title}</h2>
              <p>{article.description}</p>
              <p>{article.url}</p>
            </li>
          ))
         
        }
      </ul>
    </div>
  );
}

export default NegativePositive;
