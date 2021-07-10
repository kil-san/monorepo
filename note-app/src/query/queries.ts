import { gql } from "@apollo/client"

export const GetNote = gql`
  query GetNote($noteId: String!) {
    getNote(data: $noteId){
      id
      title
      content
    }
  }
`

export const CreateNote = gql`
  mutation CreateNote($title: String!, $content: String) {
    createNote(data: {
      title: $title
      content: $content
    }){
      id
      title
      content
    }
  }
`

export const DeleteNote = gql`
  mutation DeleteNote($noteId: String){
    deleteNote(data: $noteId)
  }
`

export const UpdateNote = gql`
  mutation UpdateNote($id: String!, $title: String!, $content: String) {
    updateNote(data: {
      id: $id
      title: $title
      content: $content
    })
  }
`

export const GetNotes = gql`
  query GetNotes {
    getNotes {
      id
      title
      content
    }
  }
`