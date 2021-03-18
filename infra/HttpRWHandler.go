package infra

type HttpRWHandler struct {
	Context string
}

func (h *HttpRWHandler) Download(url string) string {
	return "read writer handler"
}

func (h *HttpRWHandler) Write(context string) bool {
	h.Context = context
	return true
}
