import React, { useEffect, useState } from "react";
import useStyles from "./style";
import {
  Button,
  Checkbox,
  Grid,
  TextField,
  Paper,
  Typography,
  IconButton,
} from "@material-ui/core";
import { Delete } from '@material-ui/icons'
import { CheckItem } from 'context'

interface ChecklistInputProps {
  value: CheckItem
  onChange: (value: CheckItem) => void
  deleteCheckItem: (index: number) => void
}


const ChecklistInput:React.FC<ChecklistInputProps> = ({
  value,
  onChange,
  deleteCheckItem
}) => {
  const classes = useStyles();
  const[status, setStatus] = useState(value.state)
  const [title, setTitle] = useState(value.title)

  const updateValue = () => {
    onChange({
      index: value.index,
      state: status,
      title
    })
  }

  return (
    <div onBlur={updateValue}>
      <Checkbox 
        value={status} 
        onChange={e => setStatus(e.target.checked)}
      />
      <TextField
        label="Title"
        placeholder="Title"
        variant="outlined"
        className={classes.form_text}
        required
        size="small"
        onChange={e => setTitle(e.target.value)}
      />
      <IconButton onClick={() => deleteCheckItem(value.index)}>
        <Delete />
      </IconButton>
    </div>
  );
};

export default ChecklistInput;
