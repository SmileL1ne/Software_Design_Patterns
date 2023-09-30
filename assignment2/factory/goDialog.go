package main

type GoDialog struct {
}

func (w *GoDialog) createButton() Button {
	return &GoButton{}
}
