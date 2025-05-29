import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  // State for current expression, result, history, and editing ID
  const [expression, setExpression] = useState('');
  const [result, setResult] = useState('');
  const [history, setHistory] = useState([]);
  const [editingId, setEditingId] = useState(null);

  // Fetch calculation history on mount
  useEffect(() => {
    fetchHistory();
  }, []);

  // Function to fetch history (GET request)
  const fetchHistory = async () => {
    try {
      const response = await axios.get('http://localhost:8080/calculations');
      setHistory(response.data);
    } catch (error) {
      console.error('Error fetching history:', error);
    }
  };

  // Function to handle button clicks for digits and operations
  const handleButtonClick = (value) => {
    if (value === 'C') {
      setExpression('');
      setResult('');
      setEditingId(null);
    } else {
      setExpression((prev) => prev + value);
    }
  };

  // Function to handle equals button (POST or PATCH request)
  const handleEquals = async () => {
    if (!expression) return;

    try {
      let response;
      if (editingId) {
        // Update existing calculation (PATCH)
        response = await axios.patch(`http://localhost:8080/calculations/${editingId}`, {
          expression,
        });
      } else {
        // Submit new calculation (POST)
        response = await axios.post('http://localhost:8080/calculations', {
          expression,
        });
      }
      setResult(response.data.result);
      setExpression('');
      setEditingId(null);
      fetchHistory(); // Refresh history
    } catch (error) {
      console.error('Error submitting calculation:', error);
      setResult('Error');
    }
  };

  // Function to handle edit (load expression for editing)
  const handleEdit = (calc) => {
    setExpression(calc.expression);
    setEditingId(calc.id);
    setResult('');
  };

  // Function to handle delete (DELETE request)
  const handleDelete = async (id) => {
    try {
      await axios.delete(`http://localhost:8080/calculations/${id}`);
      fetchHistory(); // Refresh history
    } catch (error) {
      console.error('Error deleting calculation:', error);
    }
  };

  // Calculator button layout
  const buttons = [
    '7', '8', '9', '/',
    '4', '5', '6', '*',
    '1', '2', '3', '-',
    '0', '.', '=', '+',
    'C',
  ];

  return (
      <div className="App">
        <h1>Calculator</h1>
        <div className="calculator">
          <div className="display">
            <input type="text" value={expression} readOnly placeholder="Expression" />
            <div className="result">{result}</div>
          </div>
          <div className="buttons">
            {buttons.map((btn) => (
                <button
                    key={btn}
                    onClick={() => (btn === '=' ? handleEquals() : handleButtonClick(btn))}
                >
                  {btn}
                </button>
            ))}
          </div>
        </div>
        <div className="history">
          <h2>Calculation History</h2>
          <button onClick={fetchHistory}>Refresh History</button>
          <ul>
            {history.map((calc) => (
                <li key={calc.id}>
                  {calc.expression} = {calc.result}
                  <button onClick={() => handleEdit(calc)}>Edit</button>
                  <button onClick={() => handleDelete(calc.id)}>Delete</button>
                </li>
            ))}
          </ul>
        </div>
      </div>
  );
}

export default App;