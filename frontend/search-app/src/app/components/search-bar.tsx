"use client";

import { useState } from 'react';
import { Autocomplete, TextField, InputAdornment } from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import { SearchBarProps } from '../types/types';

const SearchBar: React.FC<SearchBarProps> = ({ onSearch }) => {
  const [query, setQuery] = useState<string>('');
  const [suggestions, setSuggestions] = useState<string[]>([]);

  const fetchSuggestions = async (value: string) => {
    if (value) {
      try {
        const response = await fetch(`http://localhost:8080/suggest?q=${value}`);
        const data = await response.json();
        setSuggestions(data); // Assuming the API returns an array of strings
      } catch (error) {
        console.error('Error fetching suggestions:', error);
      }
    } else {
      setSuggestions([]);
    }
  };

  const handleInputChange = async (event: React.SyntheticEvent, value: string) => {
    setQuery(value);
    await fetchSuggestions(value);
  };

  const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      onSearch(query, 1);
    }
  };

  return (
    <div style={{ width: '100%', maxWidth: '600px', margin: '20px auto' }}>
      <Autocomplete
        freeSolo
        options={suggestions}
        inputValue={query}
        onInputChange={handleInputChange}
        renderInput={(params) => (
          <TextField
            {...params}
            fullWidth
            variant="outlined"
            placeholder="Search..."
            onKeyPress={handleKeyPress}
            InputProps={{
              ...params.InputProps,
              startAdornment: (
                <InputAdornment position="start">
                  <SearchIcon />
                </InputAdornment>
              ),
            }}
          />
        )}
        onChange={(event, value) => {
          if (value) {
            onSearch(value, 1); // Trigger search when a suggestion is selected
          }
        }}
      />
    </div>
  );
};

export default SearchBar;