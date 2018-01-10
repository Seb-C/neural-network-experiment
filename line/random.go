package line

import "math/rand"
import "math"
import "fmt"

type PixelBitMap [20][20]bool

/**
 * Returns a random pixel bit map with a random line drawn on it
 */
func NewRandomPixelBitMap() *PixelBitMap {
	var pixelBitMap PixelBitMap // All pixels are false by default
	width := float64(len(pixelBitMap))
	height := float64(len(pixelBitMap[0]))

	// starting from x=0 and y=random point that may even be outside the image
	startX := float64(0)
	startY := math.Floor(width*2*rand.Float64() - width/2)

	// Pre-calculating line angle and max length
	maxLineLength := math.Ceil(math.Sqrt(math.Pow(width, 2) + math.Pow(height*2, 2)))
	angle := rand.Float64()*math.Pi + math.Pi/2

	for i := float64(0); i < maxLineLength; i++ {
		x := startX + math.Floor(i/2*math.Cos(angle))
		y := startY + math.Floor(i/2*math.Sin(angle))

		if x >= 0 && x < width && y >= 0 && y < height {
			pixelBitMap[int(x)][int(y)] = true
		}
	}

	return &pixelBitMap
}

func (pixelBitMap *PixelBitMap) Print() {
	for y := 0; y < len(pixelBitMap); y++ {
		for x := 0; x < len(pixelBitMap[x]); x++ {
			if pixelBitMap[x][y] {
				fmt.Print("  ")
			} else {
				fmt.Print("\u2588\u2588")
			}
		}
		fmt.Println()
	}
}
