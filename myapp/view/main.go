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
	fmt.Print(automata.ToString())

}
