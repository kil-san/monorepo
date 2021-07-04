import React, { useState, useEffect } from 'react'
import { NoteProvider, Note } from 'context/NoteContext'

interface IUseNoteContext {
    notes?: Array<Note>
}

const UseNoteContext: React.FC<IUseNoteContext> = ({ 
    notes: noteList,
    children
}) => {
  const [notes, setNotes] = useState<Array<Note>>([])

  const addNote = (note: Note) => {
    setNotes(notes => [
      ...notes,
      note
    ])
  }

  useEffect(() => {
    noteList && setNotes(noteList)
  }, [noteList])

  return (
    <NoteProvider
      value={{
        notes,
        addNote,
        setNotes
      }}
    >
      {children}
    </NoteProvider>
  )
}

export default UseNoteContext