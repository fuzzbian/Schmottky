package matplotlib

import (
	raw "Schmottky/lib/Receivers/Raw"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const tmpFile = "/dev/shm/SchmottkyTmp"

func Plot() {
	raw.ToFile(tmpFile)
	defer os.Remove(tmpFile)

	x_range := "1"
	y_range := "1"

	out, err := exec.Command("python3", "./lib/Receivers/MatPlotLib/plt.py", x_range, y_range).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", out)
}
