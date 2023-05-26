package history

import "net/http"

type History_Controller struct {
	svc History_Service
}

func NewHistoryController(svc History_Service) History_Controller {
	return History_Controller{svc}
}

func (c *History_Controller) GetAllHistories(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	c.svc.GetAllHistories().Send(w)
}