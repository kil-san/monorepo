import React from 'react'
import { INoteContext, Note } from './types'

const DefaultNoteContext: INoteContext = {
  notes: [],
  addNote: (note: Note): void => {},
  setNotes: (notes: Array<Note>): void => {}
}

const NoteContext = React.createContext(DefaultNoteContext)

const { Provider, Consumer } = NoteContext

export {
  Provider as NoteProvider,
  Consumer as NoteConsumer,
  DefaultNoteContext
}

export default NoteContext