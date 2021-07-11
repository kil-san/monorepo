import React, { useState, useContext } from 'react' // eslint-disable-line
import { useHistory } from 'react-router-dom';
import { Button, Grid, TextField, Paper, Typography } from '@material-ui/core'
import { DatePicker } from 'components'
import useStyles from './style'
import { NoteContext } from 'context/NoteContext';

const Form = () => {
  const [selectedDate, setSelectedDate] = useState(new Date().toDateString())
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [errorText, setErrorText] = useState(false)
  const history = useHistory()
  const classes = useStyles()
  const { addNote } = useContext(NoteContext)

  const handleDateChange = (date: any) => {
    console.log(date)
    setSelectedDate(date.toDateString())
  }

  const handleTitleChange = (event: any) => {
    if (event.target.value.length > 0) {
      setErrorText(false)
    }
    setTitle(event.target.value)
  }

  const handleContentChange = (event: any) => {
    if (event.target.value.length > 0) {
      setErrorText(false)
    }
    setContent(event.target.value)
  }

  const handleSubmit = (event: any) => {
    event.preventDefault()
    if (title !== '' && content !== '') {
      addNote({
        title: title,
        content: content
      })
    }
    else setErrorText(true)
  }

  return (
    <div className={classes.root}>
      <Grid container justify="center" spacing={3}>
        <Grid item xs={12} sm={12} md={6}>
          <Paper elevation={3} className={classes.form}>
            <form>
            <Typography 
              variant="h3" 
              gutterBottom
              className={classes.title}
            >
              Note to self
            </Typography>
            <DatePicker
              selectedDate={selectedDate}
              setSelectedDate={handleDateChange}
            />
            <TextField
              id="outlined-textarea"
              label="Title"
              placeholder="Title"
              variant="outlined"
              className={classes.form_text}
              value={title}
              onChange={handleTitleChange}
              required
              error={errorText}
              helperText="Title is required"
            />
            <TextField
              id="outlined-textarea"
              label="Content"
              placeholder="Content"
              multiline
              rows={5}
              variant="outlined"
              className={classes.form_text}
              value={content}
              onChange={handleContentChange}
              required
              error={errorText}
              helperText="Content is required"
            />
            <Button 
              className={classes.submit} 
              variant="contained" 
              color="primary"
              onClick={handleSubmit}
              type="submit"
            >
              Add Note
            </Button>
            </form>
          </Paper>
        </Grid>
      </Grid>
    </div>
  )
}

export default Form
