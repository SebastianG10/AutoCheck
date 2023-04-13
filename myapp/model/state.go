package model

import "fyne.io/fyne/v2"

// State representa un estado en el autómata.
type State struct {
	name string
	position fyne.Position
}

// NewState crea un nuevo objeto State con el nombre especificado.
func NewState(name string) *State {
	return &State{name: name}
}

// GetName retorna el nombre del estado.
func (s *State) GetName() string {
	return s.name
}

// SetName establece el nombre del estado.
func (s *State) SetName(name string) {
	s.name = name
}

// SetName establece el nombre del estado.
func (s *State) SetPosition(position fyne.Position) {
	s.position = position
}

// GetPosition retorna la posición del estado.
func (s *State) GetPosition() fyne.Position {
	return s.position
}