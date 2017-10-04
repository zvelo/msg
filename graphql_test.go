package msg

import "testing"

func TestGraphQL(t *testing.T) {
	if _, err := GraphQLHandler(nil); err != nil {
		t.Fatal(err)
	}
}
