package commands

import (
	"fmt"
	"github.com/jingweno/gh/utils"
	"os"
)

var cmdHelp = &Command{
	Usage: "help [command]",
	Short: "Show help",
	Long:  `Shows usage for a command.`,
}

func init() {
	cmdHelp.Run = runHelp // break init loop
}

func runHelp(cmd *Command, args *Args) {
	if args.IsParamsEmpty() {
		printUsage()
		os.Exit(0)
	}

	if args.ParamsSize() > 1 {
		utils.Check(fmt.Errorf("too many arguments"))
	}

	for _, cmd := range All() {
		if cmd.Name() == args.FirstParam() {
			cmd.PrintUsage()
			os.Exit(0)
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic: %q. Run 'git help'.\n", args.FirstParam())
	os.Exit(2)
}

var helpText = `usage: git [--version] [--exec-path[=<path>]] [--html-path] [--man-path] [--info-path]
           [-p|--paginate|--no-pager] [--no-replace-objects] [--bare]
           [--git-dir=<path>] [--work-tree=<path>] [--namespace=<name>]
           [-c name=value] [--help]
           <command> [<args>]

Basic Commands:
   init       Create an empty git repository or reinitialize an existing one
   add        Add new or modified files to the staging area
   rm         Remove files from the working directory and staging area
   mv         Move or rename a file, a directory, or a symlink
   status     Show the status of the working directory and staging area
   commit     Record changes to the repository

History Commands:
   log        Show the commit history log
   diff       Show changes between commits, commit and working tree, etc
   show       Show information about commits, tags or files

Branching Commands:
   branch     List, create, or delete branches
   checkout   Switch the active branch to another branch
   merge      Join two or more development histories (branches) together
   tag        Create, list, delete, sign or verify a tag object

Remote Commands:
   clone      Clone a remote repository into a new directory
   fetch      Download data, tags and branches from a remote repository
   pull       Fetch from and merge with another repository or a local branch
   push       Upload data, tags and branches to a remote repository
   remote     View and manage a set of remote repositories

Advanced Commands:
   reset      Reset your staging area or working directory to another point
   rebase     Re-apply a series of patches in one branch onto another
   bisect     Find by binary search the change that introduced a bug
   grep       Print files with lines matching a pattern in your codebase

GitHub Commands:
   pull-request   Open a pull request on GitHub
   fork           Make a fork of a remote repository on GitHub and add as remote
   create         Create this repository on GitHub and add GitHub as origin
   browse         Open a GitHub page in the default browser
   compare        Open a compare page on GitHub
   ci-status      Show the CI status of a commit
   release        Manipulate releases (beta)
   issue          Manipulate issues (beta)

See 'git help <command>' for more information on a specific command.
`

func printUsage() {
	fmt.Print(helpText)
}
