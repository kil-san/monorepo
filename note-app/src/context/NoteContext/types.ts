export type Note = {
  id?: string
  title?: string
  content?: string
}

export interface INoteContext {
  notes: Array<Note>
  currentNote?: Note
  addNote: (note: Note)=> void
  setNotes: (notes: Array<Note>) => void
}