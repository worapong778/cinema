package movie

func FindName(imdbID string) string {
	switch imdbID {
	case "M001":
		return "Avenger Endgame"
	case "M002":
		return "Black Panther"
	case "M003":
		return "SpiderMan"
	}
	return "not found."
}
