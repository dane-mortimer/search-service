'use client';

import { useState } from 'react';
import {
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
} from '@mui/material';

interface CreateCourseButtonProps {
  onCreate: (course: { title: string; content: string; owner: string }) => void;
}

const CreateCourseButton: React.FC<CreateCourseButtonProps> = ({
  onCreate,
}) => {
  const [open, setOpen] = useState(false);
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [owner, setOwner] = useState('');

  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  const handleSave = () => {
    onCreate({ title, content, owner });
    setOpen(false);
    setTitle('');
    setContent('');
    setOwner('');
  };

  return (
    <>
      <Button
        variant="contained"
        color="primary"
        onClick={handleOpen}
      >
        Create Course
      </Button>

      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Create New Course</DialogTitle>
        <DialogContent>
          <TextField
            label="Title"
            fullWidth
            margin="normal"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <TextField
            label="Content"
            fullWidth
            margin="normal"
            value={content}
            onChange={(e) => setContent(e.target.value)}
          />
          <TextField
            label="Owner"
            fullWidth
            margin="normal"
            value={owner}
            onChange={(e) => setOwner(e.target.value)}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} color="secondary">
            Cancel
          </Button>
          <Button onClick={handleSave} color="primary">
            Save
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default CreateCourseButton;
