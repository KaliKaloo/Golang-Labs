package main
 
func calculateNextState(p golParams, world [][]byte) [][]byte {

	newWorld := make([][]byte, p.imageHeight)
	for i := range newWorld {
		newWorld[i] = make([]byte, p.imageWidth)
	}

	for h := 0; h < p.imageHeight; h++ {
		for w := 0; w < p.imageWidth; w++ {
			neig := calculateNeighbours(p, w, h, world)
			if world[h][w] == 255 {
				if neig == 3 || neig == 2 {
					newWorld[h][w] = 255
				} else {
					newWorld[h][w] = 0
				}
			} else {
				if neig == 3 {
					newWorld[h][w] = 255
				}else{
					newWorld[h][w] = 0
				}
			}
		}
	}
	return newWorld 
}


//looking at neighbours to left and right; all combinations.
//e.g. (h-1,w-1) (h-1,w) (h-1,w+1) (h,w-1) (h.w+1) (h+1,w-1) (h+1,w) (h+1.w+1)
func calculateNeighbours(p golParams, w, h int, world [][]byte) int {
	neighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 { // NOT (h,w)
				a:= ((h+i)+p.imageHeight)%p.imageHeight
				b:= ((w+j)+p.imageWidth) % p.imageWidth
				if world[a][b] == 255 {
					neighbours++
				}
			}
		}
	}
	return neighbours
}

//takes the world as input and returns the (x,y) corordinates of all the cells that are alive
func calculateAliveCells(p golParams, world [][]byte) []cell {
	aliveList := []cell{}

	for h := 0; h < p.imageHeight; h++ {
		for w := 0; w < p.imageWidth; w++ {
			if world[h][w] == 255 {
				aliveList = append(aliveList, cell{x: w, y: h})
			}
		}
	}
	return aliveList
}
