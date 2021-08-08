package transform

import (
	"github.com/kil-san/micro-serv/note-service/pb"
	"github.com/kil-san/micro-serv/pkg/model"
)

func PbToNote(pbNote *pb.Note) model.Note {
	var checklist []model.CheckListItem
	for _, item := range pbNote.Checklist {
		checklist = append(checklist, model.CheckListItem{
			Title: item.Title,
			State: item.State,
		})
	}

	note := model.Note{
		Id:        pbNote.Id,
		Title:     pbNote.Title,
		Content:   pbNote.Content,
		Checklist: checklist,
	}

	return note
}

func PbToNoteList(pbNotes []*pb.Note) []model.Note {
	var noteList []model.Note
	for _, note := range pbNotes {
		noteList = append(noteList, PbToNote(note))
	}

	return noteList
}

func NoteToPb(note model.Note) *pb.Note {
	checklist := make([]*pb.CheckListItem, len(note.Checklist))
	for i, item := range note.Checklist {
		checklist[i] = &pb.CheckListItem{
			Title: item.Title,
			State: item.State,
		}
	}

	pbNote := &pb.Note{
		Id:        note.Id,
		Title:     note.Title,
		Content:   note.Content,
		Checklist: checklist,
	}

	return pbNote
}

func NoteToPbList(notes []model.Note) []*pb.Note {
	noteList := make([]*pb.Note, len(notes))
	for i, note := range notes {
		noteList[i] = NoteToPb(note)
	}

	return noteList
}
