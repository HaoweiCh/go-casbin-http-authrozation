package User

import (
	"fmt"
	"testing"
)

func TestIterPolicyRule(t *testing.T) {
	tmp := AuthEnforcer.GetPolicy()
	fmt.Printf("%d, %v", len(tmp), tmp)
}
