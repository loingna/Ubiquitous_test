package ErasureCode

import (
	"fmt"
	"log"
	"testing"
)

func TestGenerateExponentAndLogarithmTable(t *testing.T) {
	test := ExponentAndLogTable{}
	test.GenerateExponentAndLogTable()
	fmt.Println("exponent table:")
	err := print16(test.ExponentTable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("log table:")
	err = print16(test.LogTable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nCalculate by table 3 * 2 = %d\n", test.ExponentTable[(test.LogTable[3] + test.LogTable[2])])
}

func print16(array [256]int) error {
	i := 0
	for _, j := range(array) {
		fmt.Printf("%02x  ", j)
		i += 1
		if i == 16 {
			fmt.Println()
			i = 0
		}
	}
	return nil
}
