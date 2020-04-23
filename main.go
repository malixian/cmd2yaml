package main

import (
	"cmd2yaml/cmd"

)

func main() {
	//dockerCmd := `docker run --network host -v /tmp:/tmp -p 12345:12345 --device /dev/ your-container bash -c "python test.py"`
	//yaml := cmd2yaml(dockerCmd)
	//fmt.Printf("generate yaml is %s", yaml)

	cmd.Execute()
}
