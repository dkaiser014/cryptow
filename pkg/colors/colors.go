package colors

// Color returns a color to be later used
// by the application
func Color(color string) string {
	var selectedcolor string

	switch color {
	case "red":
		selectedcolor = "\033[31m"
	case "green":
		selectedcolor = "\033[32m"
	case "yellow":
		selectedcolor = "\033[33m"
	case "blue":
		selectedcolor = "\033[34m"
	case "purple":
		selectedcolor = "\033[35m"
	case "cyan":
		selectedcolor = "\033[36m"
	case "white":
		selectedcolor = "\033[37m"
	case "default":
		selectedcolor = "\033[0m"
	}

	return selectedcolor
}
