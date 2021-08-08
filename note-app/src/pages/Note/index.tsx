import React, { useEffect, useContext } from 'react'
import { Link, useHistory } from 'react-router-dom'
import useStyles from './style'
import { Grid, Paper, Typography, Checkbox, TextField, FormControlLabel } from '@material-ui/core'
import { ArrowBack } from '@material-ui/icons';
import { Container, HeaderBar, NoteList } from 'components'
import { NoteContext } from 'context/NoteContext';

const Note = () => {
  const history = useHistory()
  const classes = useStyles()
  const { currentNote: note } = useContext(NoteContext)

  // useEffect(() => {
  //   if (!note) {
  //     history.push("/")
  //   }
  // }, [note, history])

  return  (
    <NoteList>
      <HeaderBar showFooter>
        <Container maxWidth>
          <div className={classes.root}>
            <Grid container justify="center">
              <Grid item xs={12} sm={12} md={6}>
                <Paper elevation={3} className={classes.form}>
                  <div className={classes.title_container}>
                    <Link to="/" className={classes.back_button}>
                      <ArrowBack color="primary"/>
                    </Link>
                    <Typography 
                      variant="h3" 
                      gutterBottom
                      className={classes.title_text}
                    >
                      {note?.title}
                    </Typography>
                  </div>
                  <div>
                    <Typography 
                      variant="body1" 
                      gutterBottom
                      className={classes.body_text}
                    >
                      {note?.content}
                    </Typography>
                  </div>
                  <div>
                    {note?.checklist?.map((item, index) => {
                      return (
                        <div key={index}>
                          <FormControlLabel 
                            control={<Checkbox defaultChecked={item.state} disabled />}
                            label={item.title}
                          />
                        </div>
                      )
                    })}
                  </div>
                </Paper>
              </Grid>
            </Grid>
          </div>
        </Container>
      </HeaderBar>
    </NoteList>
  )
}


export default Note
