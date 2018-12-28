package main

import (
	"fmt"
	"os"
	"poetry"
)

func main ()  {


	fileName := "doggie.txt"
	p, err := poetry.LoadPoem(fileName)

	if err != nil {
		fmt.Printf("An Error occured reading file %s \n", fileName)
		os.Exit(-1)
	}

	fmt.Printf("%s\n", p)



}
