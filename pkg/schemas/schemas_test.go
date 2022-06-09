package schemas

import "testing"

func TestGenerateFromXSD(t *testing.T) {
	GenerateFromXSD("xunit", "xunit.xsd", "xunit", "xunit")
}
