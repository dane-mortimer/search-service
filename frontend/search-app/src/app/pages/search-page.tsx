"use client"

import { useState } from 'react';
import { Container, Pagination } from '@mui/material';
import SearchBar from '../components/search-bar';
import SearchTable from '../components/table';
import { ApiResponse, SearchResult, Pagination as PaginationType } from '../types/types';

const SearchPage: React.FC = () => {
  const defaultPagination = { page: 1, size: 10, total_pages: 1, total_items: 0 }
  const [searchResults, setSearchResults] = useState<SearchResult[]>([]);
  const [pagination, setPagination] = useState<PaginationType>(defaultPagination);
  const [query, setQuery] = useState<string>('');

  const handleSearch = async (query: string, page: number) => {
    try {
      const response = await fetch(`http://localhost:8080/search?q=${query}&page=${page}&size=${pagination.size}`);
      const result: ApiResponse = await response.json();

      if (result.data) {
        setSearchResults(result.data);
        setPagination(result.pagination);
      }
      else {
        setSearchResults([]);
        setPagination(defaultPagination);
      }

      setQuery(query);
    } catch (error) {
      console.error('Error fetching search results:', error);
    }
  };

  const handlePageChange = (event: React.ChangeEvent<unknown>, page: number) => {
    handleSearch(query, page);
  };

  return (
    <Container>
      <SearchBar onSearch={handleSearch} />
      <SearchTable data={searchResults} />
      {pagination.total_pages > 1 && (
        <div style={{ display: 'flex', justifyContent: 'center', marginTop: '20px' }}>
          <Pagination
            count={pagination.total_pages}
            page={pagination.page}
            onChange={handlePageChange}
            color="primary"
          />
        </div>
      )}
    </Container>
  );
};

export default SearchPage;