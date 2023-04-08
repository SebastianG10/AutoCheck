package main

import (
	"fmt"
	"myapp/controller"

	// Imports para interfaz gráfica
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// main es la función principal del programa.
func main() {

	interfaz()

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

func interfaz() {
	myApp := app.New()
	welcomeWindow := myApp.NewWindow("Bienvenido")
	welcomeWindow.Resize(fyne.NewSize(900, 750))

	welcomeWindow.SetContent(welcomeContent(myApp))
	welcomeWindow.Show()

	myApp.Run()
}

func logicContent() {

}

func welcomeContent(app fyne.App) *fyne.Container {
	blue := color.NRGBA{R: 50, G: 119, B: 168, A: 0xff}

	infinite := widget.NewProgressBarInfinite()

	welcomeLabel := canvas.NewText("¡AUTÓMATAS!", blue)
	welcomeLabel.TextStyle = fyne.TextStyle{Bold: true}
	welcome := container.NewCenter(welcomeLabel)

	image := canvas.NewImageFromFile("../resources/welcomeimage.png")
	image.SetMinSize(fyne.NewSize(458, 210))
	imageCont := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), image, layout.NewSpacer())

	descriptionLabel := widget.NewLabelWithStyle("Bienvenido al software donde podrá probar entradas para algunos autómatas básicos.\n", fyne.TextAlignCenter, fyne.TextStyle{})
	instructionsLabel := widget.NewLabel("Pulse el botón en cualquier momento para empezar!")

	// la función del parámetro del botón abre la ventana donde se desarrolla la lógica de la aplicaciónx
	startButton := widget.NewButtonWithIcon("Empezar", theme.MediaPlayIcon(), func() {
		var logicWindow = app.NewWindow("Autómatas")
		// logicWindow.SetContent(logicContent())
		logicWindow.Show()
	})

	developed := widget.NewLabel("Desarrollado por:")
	//datos para el card de Garcias
	gUrl, _ := url.Parse("https://github.com/SebastianG10")
	gLink := widget.NewHyperlink("https://github.com/SebastianG10", gUrl)
	gimage := canvas.NewImageFromFile("../resources/garcias.png")
	gcard := widget.NewCard("Sebastían", "García", gLink)
	gcard.SetImage(gimage)
	//datos para el card de Seider
	sUrl, _ := url.Parse("https://github.com/vanebrina")
	sLink := widget.NewHyperlink("https://github.com/vanebrina", sUrl)
	sImage := canvas.NewImageFromFile("../resources/seider.png")
	sCard := widget.NewCard("Seider", "Sanín", sLink)
	sCard.SetImage(sImage)
	//datos para el card de Jhon
	jUrl, _ := url.Parse("https://github.com/ElJhones18")
	jLink := widget.NewHyperlink("https://github.com/ElJhones18", jUrl)
	jImage := canvas.NewImageFromFile("../resources/jhon.png")
	jCard := widget.NewCard("Jhon Esteban", "Rodríguez", jLink)
	jCard.SetImage(jImage)

	cards := container.NewHBox(layout.NewSpacer(), gcard, sCard, jCard, layout.NewSpacer())

	content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		layout.NewSpacer(),
		infinite,
		welcome,
		imageCont,
		descriptionLabel,
		instructionsLabel,
		startButton,
		developed,
		cards,
		layout.NewSpacer(),
	)
	return content
}
