package model

var RUNE_TO_DEC = map[rune]string {
	' ': "000",
	'A': "001",
	'B': "002",
	'C': "003", 
	'D': "004",
	'E': "005",
	'F': "006",
	'G': "007",
	'H': "008",
	'I': "009",
	'J': "010",
	'K': "011",
	'L': "012",
	'M': "013",
	'N': "014",
	'O': "015",
	'P': "016",
	'Q': "017",
	'R': "018",
	'S': "019",
	'T': "020",
	'U': "021",
	'V': "022",
	'W': "023",
	'X': "024",
	'Y': "025",
	'Z': "026",
	'a': "051",
	'b': "052",
	'c': "053",
	'd': "054",
	'e': "055",
	'f': "056",
	'g': "057",
	'h': "058",
	'i': "059",
	'j': "060",
	'k': "061",
	'l': "062",
	'm': "063",
	'n': "064",
	'o': "065",
	'p': "066",
	'q': "067",
	'r': "068",
	's': "069",
	't': "070",
	'u': "071",
	'v': "072",
	'w': "073",
	'x': "074",
	'y': "075",
	'z': "076",
	'.': "100",
	',': "101",
	':': "102",
	';': "103",
	'-': "104",
	'_': "105",
	'/': "106",
	'\\': "107", 
	'(': "108",
	')': "109",
	'[': "110",
	']': "111",
	'{': "112",
	'}': "113",
	'!': "114",
	'?': "115",
	'@': "116",
	'#': "117",
	'$': "118",
	'%': "119",
	'^': "120",
	'&': "121",
	'*': "122",
	'+': "123",
	'=': "124",
	'|': "125",
	'"': "126",
	'\'': "127",
	'~': "128",
	'`': "129",
}

var DEC_TO_RUNE = invertMap(RUNE_TO_DEC)

func invertMap(m map[rune]string) map[string]rune {
	var result = map[string]rune{}

	for i, c := range m {
		result[c] = i
	}

	return result
}

func GetDecimalValueForCharacter(char rune) string {
	return RUNE_TO_DEC[char]
}

func GetCharacterForDecimalValue(str string) rune {
	return DEC_TO_RUNE[str]
}

