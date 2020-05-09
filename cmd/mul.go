package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var mulCommand = &cobra.Command{
	Use:   "mul",
	Short: "Multiply two or more numbers",
	Long: `Multiply two or more numbers, space seperated. 
Syntax: calcy mul num1 num2 ...
Example: calcy mul 12 2
// 24
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Printf("Minimum two numbers needed for multiplication.\n")
		} else {
			fmt.Printf("Multiplication: %f\n", getMul(args))
		}
	},
}

func getMul(array []string) float64 {
	var result float64 = 1
	for _, v := range array {
		num, _ := strconv.ParseFloat(v, 8)
		result = result * num
	}
	return result
}

func init() {
	rootCmd.AddCommand(mulCommand)
}
