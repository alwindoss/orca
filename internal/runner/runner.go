package runner

import (
	"fmt"
	"os"
	"sync"

	"github.com/alwindoss/orca/internal/connection"
	"github.com/alwindoss/orca/internal/inventory"
	"github.com/alwindoss/orca/internal/playbook"
)

type Runner struct {
	Inventory  *inventory.Inventory
	Plays      []playbook.Play
	PrivateKey []byte
}

func NewRunner(inv *inventory.Inventory, plays []playbook.Play, privateKeyPath string) (*Runner, error) {
	key, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}
	return &Runner{Inventory: inv, Plays: plays, PrivateKey: key}, nil
}

func (r *Runner) Run() error {
	var wg sync.WaitGroup
	for _, play := range r.Plays {
		group, ok := r.Inventory.Groups[play.Hosts]
		if !ok {
			return fmt.Errorf("group %s not found in inventory", play.Hosts)
		}

		for _, host := range group.Hosts {
			wg.Add(1)
			go func(h inventory.Host) {
				defer wg.Done()
				client, err := connection.NewSSHClient(h.Host, h.User, h.Port, r.PrivateKey)
				if err != nil {
					fmt.Printf("Failed to connect to %s: %v\n", h.Host, err)
					return
				}
				defer client.Close()

				for _, task := range play.Tasks {
					output, err := client.RunCommand(task.Command, task.Become)
					if err != nil {
						fmt.Printf("Task %s on %s failed: %v\nOutput: %s\n", task.Name, h.Host, err, output)
					} else {
						fmt.Printf("Task %s on %s succeeded: %s\n", task.Name, h.Host, output)
					}
				}
			}(host)
		}
	}
	wg.Wait()
	return nil
}
