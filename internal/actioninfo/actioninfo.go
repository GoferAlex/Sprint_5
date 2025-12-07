package actioninfo

import "fmt"

type DataParser interface {
	Parse(v string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	if len(dataset) == 0 {
		fmt.Println("training data is missing")
		return
	}
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			fmt.Println(err)
		}
	}
	s, err := dp.ActionInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
