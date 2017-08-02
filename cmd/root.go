// Copyright © 2017 Stream

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "vg",
	Short: "Easy and powerful workspace based development for go",
	Long: `
Virtualgo (or vg for short) is a tool which provides easy and powerful
workspace based development for Go. The vendor directory provides something
similar. However, virtualgo adds features features that are either broken or
fully missing when using a vendor directory.

Below is an example of the 'vg' command in action. For more info see detailed
command info at 'vg help <command>' or look at the README on Github:
https://github.com/GetStream/vg/blob/master/README.md

To start using virtualgo for a project run the following:

	$ cd $GOPATH/src/github.com/Getstream/example
	$ vg init  # initial creation of workspace

Now all commands will be executed from within the example workspace

	(example) $ go get github.com/pkg/errors # package only present in workspace
	(example) $ vg ensure  # installs the dependencies of the example project using dep
	(example) $ vg deactivate

If you cd back into the project the workspace is now activated automatically

	$ cd ~
	$ cd $GOPATH/src/github.com/Getstream/example
	(example) $ 

See the README on Github for more info
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
