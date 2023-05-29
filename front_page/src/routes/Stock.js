import React from "react";
import Plot from 'react-plotly.js';

class Stock extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            stockChartXValues:[],
            stockChartYValues:[]
        }
    }

    componentDidMount() {
        this.fetchStock();
    }

    fetchStock() {

        const pointerToThis = this;
        
        console.log(pointerToThis);
        const stockSymbol = "INFY"
        const API_KEY = 'Y0DT30DJI0B5PD7H'
        let API_CALL = `https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=${stockSymbol}&outputsize=compact&apikey=${API_KEY}`;

        let stockChartXValuesFunction = [];
        let stockChartYValuesFunction = [];

        fetch(API_CALL)
            .then(
                function(response) {
                    return response.json()
                }
                
            )
            .then(
                function(data) {
                    console.log(data);
                    for (var key in data['Time Series (Daily)']) {
                        stockChartXValuesFunction.push(key);
                        stockChartYValuesFunction.push(data['Time Series (Daily)'][key]['1. open']);
                    }

                    // console.log(stockChartYValuesFunction)
                    // console.log(stockChartYValuesFunction)

                    pointerToThis.setState({
                        stockChartXValues: stockChartXValuesFunction,
                        stockChartYValues:  stockChartYValuesFunction
                    })
                }
            )
    }

    render() {
        return (
            <Plot
            data={[
              {
                x: this.state.stockChartXValues,
                y: this.state.stockChartYValues,
                type: 'scatter',
                mode: 'lines+markers',
                marker: {color: 'red'},
              }
            ]}
            layout={{width: 720, height: 440, title: 'A Fancy Plot'}}
          />
        )
    }
}

export default Stock
