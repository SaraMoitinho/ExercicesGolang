//package main

//import "fmt"

//const secondInHour = 3600

/*func main() {
	fmt.Println("Hello Sara. Você irá conseguir!!")
	distance := 60.8
	fmt.Printf("The distance in miles is %f", distance*0.621)
}*/
/////////////////////////////////
// The Typical Structure of a Go Application
// Go Playground: https://play.golang.org/p/vY_IeYBb1GN
/////////////////////////////////

// a package clause starts every source file
// main is a special name declaring an executable rather than a library (package)
package main

// import declaration declares packages used in this file
import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

type DBResult struct {
	Status     string `json:"status"`
	DBPassword string `json:"db_password"`
}

func (d DBResult) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("Status", d.Status),
		slog.String("db_password", "FECHADO"),
	)
}

// package scoped variables and constants
var x int = 100

const y = 0

// a function declaration. main is a special function name
// it is the entry point for the executable program
// Go uses brace brackets to delimit code blocks
func main() {

	// Local Scoped Variables and Constants Declarations, calling functions etc
	var a int = 7
	var b float64 = 3.5
	const c int = 10

	// Println() function prints out a line to stdout
	// It belongs to package fmt
	fmt.Println("Hello Go world!") // => Hello Go world!
	fmt.Println(a, b, c)           // => 7 3.5 10
	log.Println("\nLog antigo")
	slog.Info("Trazendo informações INFO")
	slog.Error("Trazendo informações do ERRO")
	slog.Warn("Trazendo informações do WARNING")

	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	l.Info("Trazendo informações INFO")
	l.Error("Trazendo informações do ERRO")
	l.Warn("Trazendo informações do WARNING")

	level := new(slog.LevelVar)
	l = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	level.Set(slog.LevelDebug)
	l.Debug("Trazendo informações do DEBUG")

	dBResult := DBResult{Status: "sucess", DBPassword: "123456"}
	l.Info("DBRsult", dBResult)
}
