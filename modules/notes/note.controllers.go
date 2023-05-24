package notes

type Note_Controller struct {
	svc *Note_Service
}

func NewNoteController(svc *Note_Service) Note_Controller {
	return Note_Controller{svc}
}