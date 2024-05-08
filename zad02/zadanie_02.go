package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"

    "gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func lightningStrike(forest [][]rune) bool {
    rows := len(forest)
    cols := len(forest[0])
    row := rand.Intn(rows)
    col := rand.Intn(cols)

    if forest[row][col] == '🌳' {
        fmt.Printf("Piorun uderzył w drzewo na pozycji (%d, %d)!\n", row, col)
        forest[row][col] = '🔥'
        startFireSides(forest, row, col)
        return true
    } else {
        fmt.Printf("Piorun trafił w puste pole!\n")
        return false
    }
}

func startFire(forest [][]rune, row, col int) {
    rows := len(forest)
    cols := len(forest[0])
    for i := row - 1; i <= row+1; i++ {
        for j := col - 1; j <= col+1; j++ {
            if i >= 0 && i < rows && j >= 0 && j < cols {
                if forest[i][j] == '🌳' {
                    forest[i][j] = '🔥' 
                    startFire(forest, i, j)
                }
            }
        }
    }
}

func startFireCorners(forest [][]rune, row, col int) {
    rows := len(forest)
    cols := len(forest[0])

    for i := row - 1; i <= row+1; i += 2 {
        for j := col - 1; j <= col+1; j += 2 {
            if i >= 0 && i < rows && j >= 0 && j < cols {
                if forest[i][j] == '🌳' {
                    forest[i][j] = '🔥' 
                    startFireCorners(forest, i, j)
                }
            }
        }
    }
}

func startFireSides(forest [][]rune, row, col int) {
    rows := len(forest)
    cols := len(forest[0])
    for _, offset := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
        i, j := row+offset[0], col+offset[1]
        if i >= 0 && i < rows && j >= 0 && j < cols {
            if forest[i][j] == '🌳' {
                forest[i][j] = '🔥' 
                startFireSides(forest, i, j)
            }
        }
    }
}

func simulateOptimalForestDensity(trials, rows, cols, minTreeDensity, maxTreeDensity int) map[int]float64 {
    results := make(map[int]float64)

    for density := minTreeDensity; density <= maxTreeDensity; density++ {
        totalBurnedPercentage := 0.0
        totalTreeStrikes := 0

        for i := 0; i < trials; i++ {
            forest := generateForest(rows, cols, density)
            if lightningStrike(forest) {
                totalTreeStrikes++
                totalBurnedPercentage += calculateBurnedPercentage(forest)
            }
        }

        if totalTreeStrikes > 0 {
            averageBurnedPercentage := totalBurnedPercentage / float64(totalTreeStrikes)
            results[density] = averageBurnedPercentage
        } else {
            results[density] = 0.0
        }
    }

    return results
}

func calculateBurnedPercentage(forest [][]rune) float64 {
    totalTrees := 0
    burnedTrees := 0

    for _, row := range forest {
        for _, cell := range row {
            if cell == '🌳' {
                totalTrees++
            } else if cell == '🔥' {
                burnedTrees++
                totalTrees++
            }
        }
    }

    if totalTrees == 0 {
        return 0.0
    }
    return (float64(burnedTrees) / float64(totalTrees)) * 100.0
}

func displayForest(forest [][]rune) {
    for _, row := range forest {
        fmt.Println(string(row))
    }
}

func generateForest(rows, cols, treeDensity int) [][]rune {
    forest := make([][]rune, rows)
    for i := range forest {
        forest[i] = make([]rune, cols)
    }

    for i := range forest {
        for j := range forest[i] {
            forest[i][j] = '🟫'
        }
    }

    totalCells := rows * cols
    treeCells := totalCells * treeDensity / 100

    for treeCells > 0 {
        row := rand.Intn(rows)
        col := rand.Intn(cols)
        if forest[row][col] == '🟫' {
            forest[row][col] = '🌳'
            treeCells--
        }
    }

    return forest
}

func wykonajProbyLosoweIStworzWykres(rows, cols int) {

    trials := 1000    
    minTreeDensity := 1
    maxTreeDensity := 100 
    results := simulateOptimalForestDensity(trials, rows, cols, minTreeDensity, maxTreeDensity)

    pts := make(plotter.XYs, len(results))
	i := 0
	for density, burnedPercentage := range results {
		pts[i].X = float64(density)
		pts[i].Y = burnedPercentage
		i++
	}

	p := plot.New()

	p.Title.Text = "Zależność procentu spalonego lasu od procentu zalesienia"
	p.X.Label.Text = "Procent zalesienia lasu"
	p.Y.Label.Text = "Procent spalonego lasu"

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(s)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "wykres.jpg"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wykres został wygenerowany i zapisany jako wykres.jpg")
}


func main() {
    rand.Seed(time.Now().UnixNano())
    
    var option int
    var rows int
    var cols int
    var density int

    fmt.Println("Mamy do wyboru dwie opcje:\n1) Program symulujący spalanie lasu, wyświetla wizualizację spalonego lasu\n2) Wykonanie prób losowych i stworzenie wykresu")
    fmt.Scanln(&option)
	if option == 1 {
        fmt.Println("Podaj liczbę rzędów:")
        fmt.Scanln(&rows)
        fmt.Println("Podaj liczbę kolumn:")
        fmt.Scanln(&cols)
        fmt.Println("Podaj procent zalesienia:")
        fmt.Scanln(&density)

        forest := generateForest(rows, cols, density)
        displayForest(forest)
        lightningStrike(forest)
        fmt.Println("\nLas po uderzeniu pioruna:\n")
        displayForest(forest)
        burnedPercentage := calculateBurnedPercentage(forest)
        fmt.Printf("Procent spalonych drzew: %.2f%%\n", burnedPercentage)

    } else if option == 2 {
        fmt.Println("Podaj liczbę rzędów:")
        fmt.Scanln(&rows)
        fmt.Println("Podaj liczbę kolumn:")
        fmt.Scanln(&cols)

        wykonajProbyLosoweIStworzWykres(rows, cols)
    }
}