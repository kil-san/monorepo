import React, { useContext } from 'react'
import { HeaderBar, Container, Form, NoteList } from 'components'
import useStyles from './style'

const Home = () => {
  const classes = useStyles()

  return (
    <NoteList>
      <HeaderBar showFooter>
        <Container maxWidth>
          <Form />
        </Container>
      </HeaderBar>
    </NoteList>
  )
}

export default Home;
