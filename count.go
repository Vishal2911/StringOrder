package main

import (
	"fmt"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func main() {

	phrase := `English is a Germanic language of the Indo-European language family, originally spoken by the inhabitants of early medieval 
	England.[3][4][5] It is named after the Angles, one of the ancient Germanic peoples that migrated from Anglia, a peninsula on the Baltic Sea 
	(not to be confused with East Anglia), to the area of Great Britain later named after them: England. The closest living relatives of 
	English include Scots, followed by the Low Saxon and Frisian languages. English is genealogically a West Germanic language, though 
	its vocabulary is also hugely influenced by Old Norman French and Latin, as well as by Old Norse (a North Germanic language).[6][7][8]

	English has developed over the course of more than 1,400 years. The earliest forms of English, a group of West Germanic 
	(Ingvaeonic) dialects brought to Great Britain by Anglo-Saxon settlers in the 5th century and further mutated by 
	Norse-speaking Viking settlers starting in the 8th and 9th centuries, are collectively called Old English. Middle 
	English began in the late 11th century with the Norman conquest of England; this was a period in which English 
	absorbed abundant French and Latin vocabulary through Old French: in particular, its Old Norman dialect.[9][10]
	Early Modern English began in the late 15th century with the introduction of the printing press to London, the 
	printing of the King James Bible and the start of the Great Vowel Shift.`

	stringArrey := strings.Fields(strings.ToLower(phrase))
	count := map[string]int{}
	for i := 0; i < len(stringArrey); i++ {
		val, ok := count[stringArrey[i]]
		if ok {
			count[stringArrey[i]] = val + 1
		} else {
			count[stringArrey[i]] = 1
		}
	}
	list := rankByWordCount(count)
	fmt.Println()
	for i := 0 ; i < 10 ; i++ {
	
		fmt.Println(list[i].Key)
		if i+1 > len(list)-1 {
			break
		}
	}
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
