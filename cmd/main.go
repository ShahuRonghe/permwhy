package main

import (
	"flag"
	"fmt"
	"os"

	"permwhy/internal"
)

/**
 * @brief Main entry point for the permwhy command line tool.
 * @details Parses command line arguments and executes the appropriate command.
 *
 * @usage permwhy <command> <args>
 * 1. permwhy trace <path>
 *    Trace the permissions of a file or directory
 *
 * @return void
 */
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		exitWithError(fmt.Errorf("no arguments provided"))
	}
	util := args[0]
	switch util {
	case "trace":
		pathTrace(args)
	default:
		exitWithError(fmt.Errorf("invalid command: %s", util))
	}
}

/**
 * @brief Print an error message to stderr.
 * @param err The error to print.
 * @return void
 */
func printError(err error) {
	fmt.Fprintln(os.Stderr, err)
}

/**
 * @brief Print an error message to stderr and exit with code 1.
 * @param err The error to print.
 * @return void
 */
func exitWithError(err error) {
	printError(err)
	os.Exit(1)
}

/**
 * @brief Trace the permissions of a file or directory.
 * @param args The command line arguments.
 * @return void
 */
func pathTrace(args []string) {
	if len(args) < 2 {
		exitWithError(fmt.Errorf("[usage] permwhy trace <path>"))
	}
	result, err := internal.Trace(args[1])
	if err != nil {
		exitWithError(err)
	}
	fmt.Println(result)
	os.Exit(0)
}
