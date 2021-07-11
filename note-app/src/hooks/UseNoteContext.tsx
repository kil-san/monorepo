import React, { useState, useEffect } from 'react'
import { useQuery, useMutation } from "@apollo/client";
import { NoteProvider, Note } from 'context/NoteContext'
import { GetNotes, CreateNote } from 'query'

interface IUseNoteContext {
    notes?: Array<Note>
}

const UseNoteContext: React.FC<IUseNoteContext> = ({ 
    notes: noteList,
    children
}) => {
  const [notes, setNotes] = useState<Array<Note>>([])
  const { data: getNotesData, error: getNotesError } = useQuery(GetNotes{
    variables: {
      data: 1
    }
  })
  const [ createNote, createNoteMutation ] = useMutation(CreateNote)

  useEffect(() => {
    if (getNotesError) {
      console.log(getNotesError)
    }

    if (getNotesData) {
      setNotes(getNotesData.getNotes)
    }
  }, [getNotesData, getNotesError])

  useEffect(() => {
    noteList && setNotes(noteList)
  }, [noteList])

  useEffect(() => {
    const { data, error } = createNoteMutation

    if (error) {
      console.log(error)
    }

    if (data) {
      console.log(data.createNote)
    }

  }, [createNoteMutation])

  const addNote = (note: Note) => {
    setNotes(notes => [
      ...notes,
      note
    ])
    createNote({
      variables: {
        title: note.title,
        content: note.content
      }
    })
  }

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