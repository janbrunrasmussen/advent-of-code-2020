package q

import(
	"strings"
)

func ParseAnswers(input string) []string {
	groups := strings.Split(input, "\n\n")
	answers := make([]string,0)

	for _, g := range(groups){
		answers=append(answers, strings.Replace(g,"\n","",-1))
	}

	return answers
}

func ParseAnswersPerPerson(input string) [][]string {
	groups := strings.Split(input, "\n\n")
	answers := make([][]string,0)

	for _, g := range(groups){
		answers=append(answers, strings.Split(g,"\n"))
	}

	return answers
}

func CountAnswers(answers []string) int {
	count := 0
	for _, a := range(answers){
		set := make(map[string]bool,0)
		for _, c := range(a){
			// fmt.Println(string(c))
			set[string(c)]=true
		}
		// fmt.Println(len(set))
		count += len(set)
	}

	return count
}

func CountUbiquitousAnswers(answers [][]string) int {
	count := 0
	for _, a := range(answers){
		set := make(map[string]bool,0)
		// for each answer c person 1
		for _, c := range(a[0]){
			u :=true
			// for each other person o except first person
			for _, o := range(a[0:]){
				// if first persons answer c is not in other person o anwers 
				// then it's not ubiquitous
				if !strings.Contains(o,string(c)){
					u = false
				}
			}

			if u == true {
				set[string(c)]=true
			}
		}
		count += len(set)
	}
	return count
}