import { CourseResponse, PaginationResponse, SearchCourseResponse } from "./responses";

export interface SearchBarProps {
  onSearch: (query: string, page: number) => void;
}

export interface TableProps {
  data: SearchCourseResponse[];
}

export interface CourseDialogProps {
  open: boolean;
  handleClose: () => void;
  selectedCourse: CourseResponse | null;
}

export interface HomePageState {
  searchResults: SearchCourseResponse[];
  pagination: PaginationResponse;
  query: string;
}