"use client";

import { useState } from 'react';
import { Autocomplete, TextField, InputAdornment } from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import { SearchBarProps } from '../types/props';
import { COURSE_SERVICE } from '../constants/constants';

const SearchBar: React.FC<SearchBarProps> = ({ onSearch }) => {
  const [query, setQuery] = useState<string>('');
  const [suggestions, setSuggestions] = useState<string[]>([]);

  const fetchSuggestions = async (value: string) => {
    if (value) {
      try {
        const response = await fetch(`${COURSE_SERVICE}/suggest?q=${value}`);
        const data = await response.json();
        if (data.data)
          setSuggestions(data.data);
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

  const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement | HTMLDivElement>) => {
    if (e.key === 'Enter') {
      setSuggestions([]);
      onSearch(query, 1);
    }
  };

  return (
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
            onKeyDown={handleKeyPress}
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
            setSuggestions([]);
            onSearch(value, 1);
          }
        }}
        filterOptions={(options) => options} 
        open={suggestions.length > 0} 
      />
  );
};

export default SearchBar;