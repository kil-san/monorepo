import React, { useContext } from 'react';
import { Drawer, List, ListItem, ListItemIcon, ListItemText}  from '@material-ui/core';
import { MoveToInbox as InboxIcon, Mail as MailIcon }  from '@material-ui/icons';
import useStyles from './style';
import { NoteContext, Note } from 'context'
import { useHistory } from 'react-router-dom'

const drawerWidth = 240;

const NoteList: React.FC<any> = ({ children }) => {
  const classes = useStyles({ drawerWidth });
  const { notes, setCurrentNote } = useContext(NoteContext)
  const history = useHistory()

  const selectNote = (note: Note) => {
    setCurrentNote(note)
    history.push('/note')
  }

  return (
    <div className={classes.root}>
      <Drawer
        className={classes.drawer}
        variant="permanent"
        classes={{
          paper: classes.drawerPaper,
        }}
        anchor="left"
      >
        <div className={classes.toolbar} />
        <List>
          {notes.map((note, index) => (
            <ListItem button key={note.title} onClick={() => selectNote(note)}>
              <ListItemIcon>{index % 2 === 0 ? <InboxIcon /> : <MailIcon />}</ListItemIcon>
              <ListItemText primary={note.title} />
            </ListItem>
          ))}
        </List>
      </Drawer>
      <main className={classes.content}>
        {children}
      </main>
    </div>
  )
}

export default NoteList
