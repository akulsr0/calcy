package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var diffCommand = &cobra.Command{
	Use:   "diff",
	Short: "Subtract two numbers",
	Long: `Subtract two numbers, space seperated. 
Syntax: calcy diff num1 num2
Example: calcy diff 12 10
// 2
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("There should be only two numbers")
		} else {
			fmt.Printf("Difference: %f\n", getDifference(args))
		}
	},
}

func getDifference(array []string) float64 {
	x, y := array[0], array[1]
	xNum, _ := strconv.ParseFloat(x, 8)
	yNum, _ := strconv.ParseFloat(y, 8)
	diff := xNum - yNum
	return diff
}

func init() {
	rootCmd.AddCommand(diffCommand)
}
