import React, { useState, useEffect } from 'react'
import { useQuery, useMutation } from "@apollo/client";
import { NoteProvider, Note } from 'context/NoteContext'
import { GetNotes, CreateNote } from 'query'

interface IUseNoteContext {
    notes?: Array<Note>
}

const Provider: React.FC<IUseNoteContext> = ({ 
    notes: noteList,
    children
}) => {
  const [notes, setNotes] = useState<Array<Note>>([])
  const [currentNote, setCurrentNote] = useState<Note>()
  const [ createNote, createNoteMutation ] = useMutation(CreateNote)
  const { data: getNotesData, error: getNotesError } = useQuery(GetNotes, {
    variables: {
      page: 1
    }
  })

  useEffect(() => {
    if (getNotesError) {
      console.log(getNotesError)
    }

    if (getNotesData) {
      setNotes(getNotesData.getNotes)
    }

    if (!currentNote && getNotesData?.getNotes?.length > 0) {
      setCurrentNote(getNotesData.getNotes[0])
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
        data: {
          title: note.title,
          content: note.content,
          checklist: note.checklist
        }
      }
    })
  }

  return (
    <NoteProvider
      value={{
        notes,
        addNote,
        setNotes,
        currentNote,
        setCurrentNote
      }}
    >
      {children}
    </NoteProvider>
  )
}

export default Provider