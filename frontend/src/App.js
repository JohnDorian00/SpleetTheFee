import logo from './logo.svg';
import './App.css';
import {useEffect, useState} from "react";

function App() {
  const [text, setText] = useState('')
  const [loading, setLoading] = useState(false)
  useEffect(() => {
    setLoading(true)
    fetch("http://localhost/api/hello")
        .then(response => response.text())
        .then(text => {
          console.log(text)
            setText(text)
        })
        .finally(() => {
          setLoading(false)
        })
  }, [])


  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        {loading ? (
            <div>Loading...</div>
        ) : (
            <>
              <h1>{text}</h1>
            </>
        )}
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
