package inventory

import (
	"os"
	"testing"
)

func TestLoadInventory(t *testing.T) {
	content := `
groups:
  webservers:
    hosts:
      - host: 192.168.1.10
        user: ubuntu
        port: 22
`
	tmpfile, err := os.CreateTemp("", "inventory*.yml")
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}

	inv, err := LoadInventory(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to load inventory: %v", err)
	}

	if len(inv.Groups["webservers"].Hosts) != 1 {
		t.Errorf("Expected 1 host, got %d", len(inv.Groups["webservers"].Hosts))
	}
}
