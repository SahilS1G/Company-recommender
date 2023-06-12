import React from "react";
import Navbar from "../components/navbar"
import Stock from "./Stock"
import NegativePositive from "./negative_positve"

function Company() {
 return (
    <>
    <Navbar/>
    <NegativePositive/>
    <Stock/>
    
    </>
 )
}

export default Company