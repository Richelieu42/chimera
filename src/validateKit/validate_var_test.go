package validateKit

import (
	"fmt"
	"testing"
	"time"
)

func TestVar(t *testing.T) {
	fmt.Println(Var("", "required"))  // Key: '' Error:Field validation for '' failed on the 'required' tag
	fmt.Println(Var("$", "required")) // <nil>
}

func TestVarWithValue(t *testing.T) {
	// ltcsfield: Less Than Another Relative Field
	fmt.Println(VarWithValue(time.Hour, time.Hour-time.Minute, "ltcsfield")) // Key: '' Error:Field validation for '' failed on the 'ltcsfield' tag
	fmt.Println(VarWithValue(time.Hour, time.Hour+time.Minute, "ltcsfield")) // <nil>

	fmt.Println(VarWithValue(time.Duration(0), -time.Minute, "omitempty,ltcsfield")) // <nil>
}
