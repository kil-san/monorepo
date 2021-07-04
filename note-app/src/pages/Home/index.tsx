import React, { useContext } from 'react'
import { HeaderBar, Container, Form, NoteList } from 'components'
import { UseNoteContext } from 'hooks'
import useStyles from './style'

const Home = () => {
  const classes = useStyles()

  return (
    <UseNoteContext>
      <NoteList>
        <HeaderBar showFooter>
          <Container maxWidth>
            <Form />
          </Container>
        </HeaderBar>
      </NoteList>
    </UseNoteContext>
  )
}

export default Home;
