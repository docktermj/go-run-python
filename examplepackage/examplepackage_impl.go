package examplepackage

import (
	"context"
	"log"
	"os"
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

	// Set environment variables.

	os.Setenv("LD_LIBRARY_PATH", "/opt/senzing/g2/lib/")
	senzingEngineConfigurationJson := `
	        {
          "PIPELINE": {
            "CONFIGPATH": "/etc/opt/senzing",
            "RESOURCEPATH": "/opt/senzing/g2/resources",
            "SUPPORTPATH": "/opt/senzing/data"
          },
          "SQL": {
            "BACKEND": "SQL",
            "CONNECTION": "sqlite3://na:na@/tmp/sqlite/G2C.db"
          }
        }`
	os.Setenv("SENZING_ENGINE_CONFIGURATION_JSON", senzingEngineConfigurationJson)

	// Arguments

	arguments := []string{}

	// Run command.

	cmd := exec.CommandContext(ctx, "/opt/senzing/g2/python/G2ConfigTool.py", arguments...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf(">>>>>>>> x: %+v\n", x)

	// out, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(">>>>>> 2")
	// }
	// fmt.Println(string(out))
	return nil
}
