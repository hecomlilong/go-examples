package main

import "fmt"

func main() {
	fmt.Println(cook("noodles"))
	fmt.Println(cook("dumpling", "egg", "Chinese chive"))
}

func cook(food string, options ...string) string {
	if len(options) == 0 {
		return fmt.Sprintf("cook %s with nothing.\n", food)
	}
	var res string
	for i, option := range options {
		if i == 0 {
			res = fmt.Sprintf("%s%s ", res, option)
		} else {
			res = fmt.Sprintf("%sand %s ", res, option)
		}
	}
	return fmt.Sprintf("cook %s with %s.", food, res)
}
