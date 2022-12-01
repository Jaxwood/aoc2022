package main


import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
)

type Calorie struct {
    calories []float64
}

func read(filename string) []Calorie {
    readFile, err := os.Open(filename)
    defer readFile.Close()

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    var fileLines []float64
    var calories []Calorie

    for fileScanner.Scan() {
        line := fileScanner.Text()
        if line == "" {
            calories = append(calories, Calorie{calories: fileLines})
            fileLines = make([]float64, 0)
            continue
        }
        num, err := strconv.ParseFloat(line, 64)
        if err != nil {
            fmt.Println(err)
        }
        fileLines = append(fileLines, num)
    }
    calories = append(calories, Calorie{calories: fileLines})


    return calories
}

func day01a(filename string) float64 {

    lines := read(filename)
    var best = 0.0
    for _, line := range lines {
        var sum = 0.0
        for _, calorie := range line.calories {
            sum += calorie
        }
        best = math.Max(best, sum)
    }

    return best
}

func day01b(filename string) float64 {

    lines := read(filename)
    var best = 0.0
    for _, line := range lines {
        var sum = 0.0
        for _, calorie := range line.calories {
            sum += float64(calorie)
        }
        best = math.Max(best, sum)
    }

    return best
}
