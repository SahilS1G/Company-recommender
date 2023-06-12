import React from "react";

function Footer() {
    return (
      <footer style={footerStyle}>
        <p style={footerText}>© 2023 Your Website. All rights reserved.</p>
      </footer>
    );
  }

  const footerStyle = {
    backgroundColor: '#333',
    color: '#fff',
    padding: '10px',
    textAlign: 'center',
    height: '50px',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    
  };
  
  const footerText = {
    margin: 0,
  };

export default Footer;