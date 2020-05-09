package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add two or more numbers",
	Long: `Add two or more numbers, space seperated. 
Syntax: calcy add num1 num2 ...
Example: calcy add 12 43 64 54 23
// 196
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Printf("Minimum two numbers needed for addition.\n")
		} else {
			fmt.Printf("Sum: %f\n", getSum(args))
		}
	},
}

func getSum(array []string) float64 {
	var result float64 = 0
	for _, v := range array {
		num, _ := strconv.ParseFloat(v, 8)
		result += num
	}
	return result
}

func init() {
	rootCmd.AddCommand(addCommand)
}
