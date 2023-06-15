package movie

func FindName(imdbID string) string {
	switch imdbID {
	case "title1":
		return "Avengers: Wick"
	case "title2":
		return "Black Panther"
	}

	return "not found"
}
