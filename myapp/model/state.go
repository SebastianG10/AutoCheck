package model

// State representa un estado en el autómata.
type State struct {
	name string
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
