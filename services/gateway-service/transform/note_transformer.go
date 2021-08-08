package transform

import (
	"github.com/kil-san/micro-serv/gateway-service/graph/model"
	"github.com/kil-san/micro-serv/note-service/pb"
)

func PbToNote(pbNote *pb.Note) *model.Note {
	checklist := make([]*model.CheckListItem, len(pbNote.Checklist))
	for i, item := range pbNote.Checklist {
		checklist[i] = &model.CheckListItem{
			Title: item.Title,
			State: item.State,
		}
	}

	note := &model.Note{
		ID:        pbNote.Id,
		Title:     pbNote.Title,
		Content:   pbNote.Content,
		Checklist: checklist,
	}

	return note
}

func PbToNoteList(pbNotes []*pb.Note) []*model.Note {
	noteList := make([]*model.Note, len(pbNotes))
	for i, note := range pbNotes {
		noteList[i] = PbToNote(note)
	}

	return noteList
}

func NoteUpdateToPb(note model.NoteUpdate) *pb.Note {
	checklist := make([]*pb.CheckListItem, len(note.Checklist))
	for i, item := range note.Checklist {
		checklist[i] = &pb.CheckListItem{
			Title: item.Title,
			State: item.State,
		}
	}

	pbNote := &pb.Note{
		Id:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		Checklist: checklist,
	}

	return pbNote
}

func NewNoteToPb(note model.NewNote) *pb.Note {
	checklist := make([]*pb.CheckListItem, len(note.Checklist))
	for i, item := range note.Checklist {
		checklist[i] = &pb.CheckListItem{
			Title: item.Title,
			State: item.State,
		}
	}

	pbNote := &pb.Note{
		Title:     note.Title,
		Content:   note.Content,
		Checklist: checklist,
	}

	return pbNote
}
