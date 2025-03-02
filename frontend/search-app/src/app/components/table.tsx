'use client';

import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
} from '@mui/material';
import { TableProps } from '../types/props';
import { useState } from 'react';
import { COURSE_SERVICE } from '../constants/constants';
import { CourseResponse } from '../types/responses';
import CourseDialog from './course-dialog';

const SearchTable: React.FC<TableProps> = ({ data }) => {
  const [selectedCourse, setSelectedCourse] = useState<CourseResponse | null>(
    null
  );
  const [open, setOpen] = useState(false);

  const handleRowClick = async (id: string) => {
    try {
      const response = await fetch(`${COURSE_SERVICE}/${id}`);
      const apiResponse = await response.json();
      if (apiResponse.status != 'success') return;
      setSelectedCourse(apiResponse.data);
      setOpen(true);
    } catch (error) {
      console.error('Error fetching course details:', error);
    }
  };

  const handleClose = () => {
    setOpen(false);
    setSelectedCourse(null);
  };

  return (
    <>
      <TableContainer component={Paper} style={{ marginTop: '20px' }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Title</TableCell>
              <TableCell>Content</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data.length === 0 ? (
              <TableRow>
                <TableCell
                  colSpan={3}
                  align="center"
                  style={{ textAlign: 'center' }}
                >
                  No Results
                </TableCell>
              </TableRow>
            ) : (
              data.map((item) => (
                <TableRow
                  key={item.id}
                  hover
                  onClick={() => handleRowClick(item.id)}
                  style={{ cursor: 'pointer' }}
                >
                  <TableCell>{item.id}</TableCell>
                  <TableCell>{item.title}</TableCell>
                  <TableCell>{item.content}</TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </TableContainer>
      <CourseDialog open={open} handleClose={handleClose} selectedCourse={selectedCourse} />
    </>
  );
};

export default SearchTable;
