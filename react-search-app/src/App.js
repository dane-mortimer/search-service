import React, { useState } from "react";
import { TextField, Autocomplete, Table, TableHead, TableRow, TableCell, TableBody } from "@mui/material";
import "./App.css";

function App() {
  const [query, setQuery] = useState("");
  const [value, setValue] = useState("");
  const [suggestions, setSuggestions] = useState([]);
  const [results, setResults] = useState([]);
  const [error, setError] = useState("No Results");

  const handleChange = async (e) => {

    if (e.key === 'Enter') {
      // Prevent's default 'Enter' behavior.
      e.defaultMuiPrevented = true;      
      handleSearch(suggestions.length > 0 ? suggestions[0] : value)
      return
    }
  };

  const handleInputChange = async (e) => {

    if (!e)
      return
    const value = e.target.value;

    setValue(value)

    if (value) {
      try {
        const response = await fetch(`http://localhost:8080/suggest?q=${value}`);
        const data = await response.json();
        setSuggestions(data);
        
        setQuery(data[0]);
      } catch (error) {
        console.error("Error fetching suggestions:", error);
      }
    } else {
      setSuggestions([]);
    }
  };

  const handleSearch = async (query) => {
    try {
      const response = await fetch(
        `http://localhost:8080/search?q=${query}&page=1&size=10`
      );
      const data = await response.json();
      if (data.data) { 
        setError(null)
        setResults(data.data);
      }
      else 
        setError("No Results")
    } catch (error) {
      console.error("Error searching:", error);
    }
  };

  return (
    <div className="App">
      <h1>Search App</h1>
      <div className="search-container">
        <Autocomplete
          disablePortal
          filterOptions={(x) => x}
          getOptionLabel={(option) => option || ""}
          options={suggestions}
          sx={{ width: 300 }}
          value={value}
          onKeyUp={handleChange}
          onInputChange={handleInputChange}
          renderInput={(params) => <TextField {...params} label="Search" />}
        />
      </div>
      {error && (
        <h1>No Results</h1>
      )}
      {results.length > 0 && (
        <Table style={{width: "50dvw"}}>
          <TableHead>
            <TableRow>
              <TableCell>Title</TableCell>
              <TableCell>Content</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {results.map((result) => (
              <TableRow key={result.id}>
                <TableCell>{result.title}</TableCell>
                <TableCell>{result.content}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )}
    </div>
  );
}

export default App;
