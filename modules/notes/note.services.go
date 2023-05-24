package notes

type Note_Service struct {
	repo Note_Repo
}

func NewNoteService(repo Note_Repo) *Note_Service {
	return &Note_Service{repo}
}