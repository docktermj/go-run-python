package examplepackage

import (
	"context"
	"fmt"
	"os/exec"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// ExamplePackageImpl is an example type-struct.
type ExamplePackageImpl struct {
	Something string
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const exampleConstant = "examplePackage"

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The SaySomething method simply prints the 'Something' value in the type-struct.

Input
  - ctx: A context to control lifecycle.

Output
  - Nothing is returned, except for an error.  However, something is printed.
    See the example output.
*/
func (examplepackage *ExamplePackageImpl) SaySomething(ctx context.Context) error {
	fmt.Printf("%s: %s\n", exampleConstant, examplepackage.Something)

	fmt.Println(">>>>>> 1")

	err := exec.Command("/opt/senzing/g2/python/G2ConfigTool.py").Run()
	if err != nil {
		fmt.Println(">>>>>> 1")

		panic(err)
		fmt.Println(">>>>>> 1")

	}
	return nil
}
