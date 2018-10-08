package main

import (
	"fmt"
	p "github.com/prometheus/procfs"
)

func main() {
	procs, _ := p.AllProcs()
	for _, proc := range procs {
		cmd, _ := proc.Comm()
		stats, err := proc.InotifyStats()
		if err != nil {
			fmt.Println(err)
		}

		for _, s := range stats {
			if len(s.Lines) != 0 {
				fmt.Println(cmd)
				for _, i := range s.Lines {
					fmt.Printf("  %v\n", i)
				}
			}
		}
	}
}
