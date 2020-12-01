package expense_report

func Checksum2(input []int) int{

	for _, n1 := range(input){
		for _, n2 := range(input){
			if n1 + n2 == 2020 {
				return n1 * n2
			}
		}
	}

	return 0
}

func Checksum3(input []int) int{

	for _, n1 := range(input){
		for _, n2 := range(input){
			for _, n3 := range(input){
				if n1 + n2 + n3 == 2020 {
					return n1 * n2 * n3
				}	
			}
		}
	}

	return 0
}
