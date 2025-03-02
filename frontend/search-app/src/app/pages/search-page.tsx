'use client';

import { useState } from 'react';
import { Container, Pagination } from '@mui/material';
import SearchBar from '../components/search-bar';
import SearchTable from '../components/table';
import { COURSE_SERVICE } from '../constants/constants';
import {
  PaginationResponse,
  SearchCourseResponse,
} from '../types/responses';
import CreateCourseButton from '../components/create-button';
import useEnhancedEffect from '@mui/material/utils/useEnhancedEffect';

const SearchPage: React.FC = () => {
  const defaultPagination = {
    page: 1,
    size: 10,
    total_pages: 1,
    total_items: 0,
  };
  const [searchResults, setSearchResults] = useState<SearchCourseResponse[]>(
    []
  );
  const [pagination, setPagination] =
    useState<PaginationResponse>(defaultPagination);
  const [query, setQuery] = useState<string>('');

  const handleSearch = async (query: string, page: number) => {
    try {
      const response = await fetch(
        `${COURSE_SERVICE}/search?q=${query}&page=${page}&size=${pagination.size}`
      );
      const result = await response.json();

      if (result.data) {
        setSearchResults(result.data);
        setPagination(result.pagination);
      } else {
        setSearchResults([]);
        setPagination(defaultPagination);
      }

      setQuery(query);
    } catch (error) {
      console.error('Error fetching search results:', error);
    }
  };

  const handlePageChange = (
    event: React.ChangeEvent<unknown>,
    page: number
  ) => {
    handleSearch(query, page);
  };

  const handleCreateCourse = async (course: {
    title: string;
    content: string;
    owner: string;
  }) => {
    try {
      const response = await fetch(`${COURSE_SERVICE}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(course),
      });

      if (!response.ok) {
        throw new Error('Failed to create course');
      }

      // Optionally, refresh the search results after creating a course
      handleSearch(query, pagination.page);
    } catch (error) {
      console.error('Error creating course:', error);
    }
  };

  // useEnhancedEffect(() => {
  //   handleSearch("", 1) 
  // });

  return (
    <Container>
      <div
        style={{
          display: 'flex',
          alignItems: 'center',
          gap: '10px',
          margin: '20px auto',
          justifyContent: 'center',
          width: '30dvw',
        }}
      >
        <div style={{ flex: 1 }}>
          <SearchBar onSearch={handleSearch} />
        </div>
        <CreateCourseButton onCreate={handleCreateCourse} />
      </div>
      <SearchTable data={searchResults} />
      {pagination.total_pages > 1 && (
        <div
          style={{
            display: 'flex',
            justifyContent: 'center',
            marginTop: '20px',
          }}
        >
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
