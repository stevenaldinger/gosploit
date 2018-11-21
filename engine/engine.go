package engine

import (
	"fmt"
	"os/exec"
	tm "github.com/buger/goterm"
	"sync"
	"time"
	"github.com/sethgrid/multibar"
	"plugin"
	"strings"
	"github.com/fatih/color"
	"github.com/abiosoft/ishell"
	"github.com/stevenaldinger/gosploit/auxiliary"
	"github.com/stevenaldinger/gosploit/modules/exploits/www/wordpress"
)

func RunShell() {
	//Print Welcome Info
	multiline := `
	───▄▀▀▀▄▄▄▄▄▄▄▀▀▀▄───
	───█▒▒░░░░░░░░░▒▒█───
	────█░░█░░░░░█░░█────
	─▄▄──█░░░▀█▀░░░█──▄▄─
	█░░█─▀▄░░░░░░░▄▀─█░░█
	`
	color.Red(multiline)

	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	color.Red("GOSPLOIT V1")

	shell.SetPrompt("gosploit > ")

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "greet",
		Help: "greet user",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})

	// Test shell function
	shell.AddCmd(&ishell.Cmd{
		Name: "auxiliary/scanner/xss",
		Help: "scan targets for xss vulnerabilties",
		Func: func(c *ishell.Context) {
			// get RHOST
			c.Print("RHOST: ")
			domain := c.ReadLine()
			auxiliary.XSS_Scan(domain)
		},
	})

	// Test shell function
	shell.AddCmd(&ishell.Cmd{
		Name: "exploits/www/wordpress/content_injection",
		Help: "Test WP for < 4.7 Content Injection",
		Func: func(c *ishell.Context) {
			// get RHOST
			c.Print("RHOST: ")
			domain := c.ReadLine()
			wordpress.ContentInjection_4_7(domain)
		},
	})

	// run shell
	shell.Run()
}

func RunGoSploit() {
	//tm.Clear() // Clear current screen

	app := "nmap"
	//app := "buah"

	arg0 := "-sV"
	arg1 := "-sC"
	arg2 := "127.0.0.1"
	arg3 := "-p 443"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(100|tm.PCT, 6, 0)
	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box, string(stdout))
	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box.String(), 0|tm.PCT, 40|tm.PCT))
	tm.Flush()

}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ProgressBar() {
	// create the multibar container
	// this allows our bars to work together without stomping on one another
	progressBars, _ := multibar.New()

	// some arbitrary totals for our  progress bars
	// in practice, these could be file sizes or similar
	mediumTotal, smallTotal, largerTotal := 150, 100, 200

	// we will update the progress down below in the mock work section with barProgress1(int)
	barProgress1 := progressBars.MakeBar(mediumTotal, "1st")

	progressBars.Println()
	progressBars.Println("We can separate bars with blocks of text, or have them grouped.\n")

	barProgress2 := progressBars.MakeBar(smallTotal, "2nd - with description:")
	barProgress3 := progressBars.MakeBar(largerTotal, "3rd")
	barProgress4 := progressBars.MakeBar(mediumTotal, "4th")
	barProgress5 := progressBars.MakeBar(smallTotal, "5th")
	barProgress6 := progressBars.MakeBar(largerTotal, "6th")

	progressBars.Println("And we can have blocks of text as we wait for progress bars to complete...")

	// listen in for changes on the progress bars
	// I should be able to move this into the constructor at some point
	go progressBars.Listen()

	/*

		*** mock work ***
		spawn some goroutines to do arbitrary work, updating their
		respective progress bars as they see fit

	*/
	wg := &sync.WaitGroup{}
	wg.Add(6)
	go func() {
		// do something asyn that we can get updates upon
		// every time an update comes in, tell the bar to re-draw
		// this could be based on transferred bytes or similar
		for i := 0; i <= mediumTotal; i++ {
			barProgress1(i)
			time.Sleep(time.Millisecond * 15)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= smallTotal; i++ {
			barProgress2(i)
			time.Sleep(time.Millisecond * 25)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= largerTotal; i++ {
			barProgress3(i)
			time.Sleep(time.Millisecond * 12)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= mediumTotal; i++ {
			barProgress4(i)
			time.Sleep(time.Millisecond * 10)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i <= smallTotal; i++ {
			barProgress5(i)
			time.Sleep(time.Millisecond * 20)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i <= largerTotal; i++ {
			barProgress6(i)
			time.Sleep(time.Millisecond * 10)
		}
		wg.Done()
	}()
	wg.Wait()

	// continue doing other work
	fmt.Println("All Bars Complete")
}

type GosploitModule interface {
	Exploit()
}

func LoadModule(s string) {

	modulepath := strings.TrimSuffix(s, "\n")
	modulepath = strings.TrimPrefix(modulepath, "use ")
	var mod string

	switch modulepath {
	case "test/chi/chi":
		mod = "./modules/test/chi/chi.so"
	case "test/eng/eng":
		mod = "./modules/test/eng/eng.so"
	default:
		fmt.Println("can't find module")
	}

	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable GosploitModule
	symGosploitModule, err := plug.Lookup("GosploitModule")
	if err != nil {
		fmt.Println(err)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type GosploitModule (defined above)
	var module GosploitModule
	module, ok := symGosploitModule.(GosploitModule)
	if !ok {
		fmt.Println("unexpected type from module symbol")
	}

	// 4. use the module
	module.Exploit()

}
