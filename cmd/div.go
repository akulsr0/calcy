package cmd

import (
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

var divCommand = &cobra.Command{
	Use:   "div",
	Short: "Divide two numbers",
	Long: `Divide two numbers, space seperated. 
Syntax: calcy diff num1 num2
Example: calcy diff 12 6
// 2
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("There should be only two numbers")
		} else {
			q, r := getDivison(args)
			fmt.Printf("Quotient: %f and Remainder: %f\n", q, r)
		}
	},
}

func getDivison(array []string) (float64, float64) {
	x, y := array[0], array[1]
	xNum, _ := strconv.ParseFloat(x, 8)
	yNum, _ := strconv.ParseFloat(y, 8)
	q := xNum / yNum
	r := math.Mod(xNum, yNum)
	return q, r
}

func init() {
	rootCmd.AddCommand(divCommand)
}
