import React from "react";
import Plot from "react-plotly.js";
import Search from "../components/searchbar";

class Stock extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      stockChartXValues: [],
      stockChartYValues: [],
    };
  }

  componentDidMount() {
    this.fetchStock();
  }

  fetchStock() {
    const pointerToThis = this;
    var company = "infosys";
    const API_CALL_2 = `https://finnhub.io/api/v1/search?q=${company}&token=chtn9npr01qvn1d2ou7gchtn9npr01qvn1d2ou80`;

    fetch(API_CALL_2)
      .then(function (response) {
        return response.json();
      })
      .then(function (data) {
        let stockSymbol = data["result"][0]["symbol"];
        pointerToThis.fetchStockPrice(stockSymbol);
      });
  }

  fetchStockPrice(stockSymbol) {
    const pointerToThis = this;
    const API_KEY = "Y0DT30DJI0B5PD7H";
    let API_CALL = `https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=${stockSymbol}&outputsize=compact&apikey=${API_KEY}`;

    let stockChartXValuesFunction = [];
    let stockChartYValuesFunction = [];

    fetch(API_CALL)
      .then(function (response) {
        return response.json();
      })
      .then(function (data) {
        console.log(data);
        for (var key in data["Time Series (Daily)"]) {
          stockChartXValuesFunction.push(key);
          stockChartYValuesFunction.push(
            data["Time Series (Daily)"][key]["1. open"]
          );
        }

        pointerToThis.setState({
          stockChartXValues: stockChartXValuesFunction,
          stockChartYValues: stockChartYValuesFunction,
        });
      });
  }

  render() {
    const stockChartStyle = {
      backgroundColor: "#333",
      padding: "20px",
      boxShadow: "0 0 10px rgba(0, 0, 0, 0.5)",
    };

    const headingStyle = {
      marginBottom: "10px",
      color: "#fff",
    };

    const plotLayout = {
      width: 720,
      height: 440,
      title: "Stock Chart",
      paper_bgcolor: "#333",
      plot_bgcolor: "#333",
      font: {
        color: "#fff",
      },
    };

    return (
      <div style={stockChartStyle}>
        <h2 style={headingStyle}>Stock Chart</h2>
        <Plot
          data={[
            {
              x: this.state.stockChartXValues,
              y: this.state.stockChartYValues,
              type: "scatter",
              mode: "lines+markers",
              marker: { color: "red" },
            },
          ]}
          layout={plotLayout}
        />
      </div>
    );
  }
}

export default Stock;
