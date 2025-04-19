# orca
Orchestrate automation

## How to run
### Build the project
Run the command `make build`

### Run the application
Run the command `./bin/orca run --inventory inventory.yml --playbook playbook.yml --key ~/.ssh/id_rsa`

## Prepare target systems
### Generate SSH Keys
Run the command to generate SSH keys `ssh-keygen -t rsa -b 4096 -f ~/.ssh/id_rsa`

### Certificate based login
Run the command to copy the public key to the target host `ssh-copy-id <username>@<IP Address>`