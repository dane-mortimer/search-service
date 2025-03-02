import React from 'react';
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
} from '@mui/material';
import { CourseDialogProps } from '../types/props';

const CourseDialog: React.FC<CourseDialogProps> = ({
  open,
  handleClose,
  selectedCourse,
}) => {
  return (
    <Dialog
      open={open}
      onClose={handleClose}
      sx={{
        '& .MuiDialog-paper': {
          width: '25vw', // 25% of the viewport width
          height: '40vh', // 25% of the viewport height
          maxWidth: 'none', // Override default max-width
          maxHeight: 'none', // Override default max-height
          margin: 'auto', // Center the dialog
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
          padding: '20px', // Add padding for better spacing
          borderRadius: '12px', // Rounded corners
          boxShadow: '0px 4px 20px rgba(0, 0, 0, 0.2)', // Subtle shadow
        },
      }}
    >
      {selectedCourse && (
        <>
          <DialogTitle
            sx={{
              fontFamily: 'Georgia, serif', // Professional serif font
              fontSize: '1.5rem', // Larger font size
              fontWeight: 'bold', // Bold text
              color: '#2c3e50', // Dark blue-gray color
              textAlign: 'center', // Center align title
              lineBreak: 'loose'
            }}
          >
            {selectedCourse.Title}
          </DialogTitle>
          <DialogContent
            sx={{
              fontFamily: 'Arial, sans-serif',
              fontSize: '1rem',
              color: '#34495e',
              marginBottom: '10px',
            }}
          >
            <DialogContentText style={{ textAlign: 'center' }}>
              {selectedCourse.Content}
            </DialogContentText>
            <br/><br/>
            <DialogContentText>
              <strong>Owner:</strong> {selectedCourse.Owner}
            </DialogContentText>
            <DialogContentText>
              <strong>Created Date:</strong>{' '}
              {new Date(selectedCourse.CreatedAt).toLocaleTimeString()} {new Date(selectedCourse.CreatedAt).toLocaleDateString()}
            </DialogContentText>
          </DialogContent>
        </>
      )}
      <DialogActions>
        <Button
          onClick={handleClose}
          color="primary"
          sx={{
            fontFamily: 'Arial, sans-serif',
            fontWeight: 'bold',
            textTransform: 'none', // Disable uppercase transformation
          }}
        >
          Close
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default CourseDialog;
