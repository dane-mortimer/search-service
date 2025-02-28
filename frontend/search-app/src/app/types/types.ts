// types.ts
export interface SearchResult {
  id: string;
  title: string;
  content: string;
}

export interface Pagination {
  page: number;
  size: number;
  total_pages: number;
  total_items: number;
}

export interface ApiResponse {
  data: SearchResult[];
  pagination: Pagination;
}

export interface SearchBarProps {
  onSearch: (query: string, page: number) => void;
}

export interface TableProps {
  data: SearchResult[];
}

export interface HomePageState {
  searchResults: SearchResult[];
  pagination: Pagination;
  query: string;
}