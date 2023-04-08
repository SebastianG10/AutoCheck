package model

// Transition representa una transición en el autómata.
type Transition struct {
	from  *State
	to    *State
	input string
}

// NewTransition crea una nueva transición con el estado de origen, estado destino y símbolo de entrada especificados.
func NewTransition(from *State, to *State, input string) *Transition {
	return &Transition{
		from:  from,
		to:    to,
		input: input,
	}
}

// GetFromState retorna el estado de origen de la transición.
func (t *Transition) GetFromState() *State {
	return t.from
}

// SetFromState establece el estado de origen de la transición.
func (t *Transition) SetFromState(from *State) {
	t.from = from
}

// GetToState retorna el estado destino de la transición.
func (t *Transition) GetToState() *State {
	return t.to
}

// SetToState establece el estado destino de la transición.
func (t *Transition) SetToState(to *State) {
	t.to = to
}

// GetInput retorna el símbolo de entrada de la transición.
func (t *Transition) GetInput() string {
	return t.input
}

// SetInput establece el símbolo de entrada de la transición.
func (t *Transition) SetInput(input string) {
	t.input = input
}
