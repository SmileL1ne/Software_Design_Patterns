package main

type HtmlDialog struct {
}

func (h *HtmlDialog) createButton() Button {
	return &HtmlButton{}
}
