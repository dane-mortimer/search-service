export interface PaginationResponse {
  page: number;
  size: number;
  total_pages: number;
  total_items: number;
}

export interface CourseApiResponse {
  data: SearchCourseResponse[] | CourseResponse[] | CourseResponse | null;
  status: string;
  pagination: PaginationResponse;
}

export interface SearchCourseResponse {
  id: string;
  title: string;
  content: string;
}

export interface CourseResponse {
  ID: string;
  CreatedAt: string;
  Owner: string;
  Title: string;
  Content: string;
}

