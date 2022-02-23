
# Fare Estimation Calculator
#### Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

#### Requirements

- Go 1.16

### Execute
**First Time:**

    go mod install
**For Run:**

    go build
    ./fare-estimation execute

### Tests

    go test ./...
### Project Structure

    |-assets (input, output files)
    |-cmd
	    |-cli (cli commands to execute in project)
	|-internal (internal codes does not to be shared)
	|-pkg (common codes like service)
	




