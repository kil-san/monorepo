import React, { useContext } from "react"
import { NoteContext } from 'context'

const useNoteContext = () => {
  const contextData = useContext(NoteContext)

  return {
    ...contextData
  }
}

export default useNoteContext