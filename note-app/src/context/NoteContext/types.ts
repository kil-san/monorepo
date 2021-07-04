export type Note = {
  id?: string
  title?: string
  status?: string
}

export interface INoteContext {
  notes: Array<Note>
  addNote: (note: Note)=> void
  setNotes: (notes: Array<Note>) => void
}