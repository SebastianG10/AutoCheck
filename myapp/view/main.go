package main

import (
	"fmt"
	"myapp/controller"
)

// main es la función principal del programa.
func main() {
	control := &controller.ControlAutomatas{}

	stateMap := control.ReadState("q0,q1")
	symbols := control.ReadSymbols("0,1")
	transitionsList := control.ReadTransitions("[q0,0,q1],[q0,1,q0],[q1,0,q1],[q1,1,q0]", stateMap)
	initialState := control.ReadInitialState("q0")
	finalStatesMap := control.ReadFinalStates("q1")

	automata := control.CreateAutomata(transitionsList, initialState, finalStatesMap, stateMap, symbols)

	// Prueba para ProcessInput
	input := "1010"
	automata.ProcessInput(input)
	fmt.Printf("Después de procesar la entrada '%s', el estado actual es: %s\n", input, automata.GetCurrentState().GetName())

	fmt.Println("Historial de estados actuales:")
	for _, state := range automata.GetHistoryCurrentState() {
		fmt.Printf("%s -> ", state.GetName())
	}

	// Prueba para IsAccepted
	if automata.IsAccepted() {
		fmt.Println("La entrada es aceptada por el autómata.")
	} else {
		fmt.Println("La entrada no es aceptada por el autómata.")
	}

}
