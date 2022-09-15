package cobman

var (
	seeAlso = map[string]int{}
)

func SeeAlso(item string, section int) {
	seeAlso[item] = section
}
