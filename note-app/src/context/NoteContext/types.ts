export type CheckItem = {
  index: number
  state: boolean
  title: string
}

export type Note = {
  id?: string
  title?: string
  content?: string
  checklist?: CheckItem[]
}

export interface INoteContext {
  notes: Note[]
  currentNote?: Note
  addNote: (note: Note) => void
  setNotes: (notes: Note[]) => void
  setCurrentNote: (note: Note) => void
}