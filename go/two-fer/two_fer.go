package twofer

// ShareWith takes a name parameter and returns a string that includes it.
// If the name parameter is not included, ShareWith simply returns "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
