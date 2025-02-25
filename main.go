package main

import (
	"fmt"
	"time"

	"test.com/utils"
)

func main() {
	filename := "tetragrid_40k.vtk"
	start := time.Now()
	grid2d, _ := utils.Grid("test_data/" + filename)
	grid2d.Need_cell_centers()
	var solver utils.Solver
	solver.Set_bnd_type(2)
	solver.Set_grid(grid2d)
	solver.Approximate_parts()
	solver.Write_to_file(filename)
	time.Sleep(2 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("time = ", elapsed)
}
