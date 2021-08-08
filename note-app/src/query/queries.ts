import { gql } from "@apollo/client"

export const GetNote = gql`
  query GetNote($noteId: String!) {
    getNote(data: $noteId){
      id
      title
      content
      checklist{
        title
        state
      }
    }
  }
`

export const CreateNote = gql`
  mutation CreateNote($data: NewNote!) {
    createNote(data: $data){
      id
      title
      content
      checklist{
        title
        state
      }
    }
  }
`

export const DeleteNote = gql`
  mutation DeleteNote($noteId: String){
    deleteNote(data: $noteId)
  }
`

export const UpdateNote = gql`
  mutation UpdateNote($data: NoteUpdate!) {
    updateNote(data: $data)
  }
`

export const GetNotes = gql`
  query GetNotes($page: Int!) {
    getNotes(data: $page) {
      id
      title
      content
      checklist{
        title
        state
      }
    }
  }
`