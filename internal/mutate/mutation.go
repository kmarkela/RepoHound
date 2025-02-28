package mutatetion

// var separatos = [...]string{".", "-", ""}
const seperator = "-"

func Mutate(ulist []string, kw string) []string {

	var rList = []string{}

	for _, u := range ulist {
		rList = append(rList, u)
		rList = append(rList, u+seperator+kw)
		rList = append(rList, kw+seperator+u)

	}

	return rList
}
