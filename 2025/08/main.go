package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y, z float64
	circuit []*point
}

func StringArrayToPoint(strArray []string) (*point, error) {
	p := new(point)
	floatValue, err := strconv.ParseFloat(strArray[0], 64)
	if err != nil {
		return nil, err
	}
	p.x = floatValue
	floatValue, err = strconv.ParseFloat(strArray[1], 64)
	if err != nil {
		return nil, err
	}
	p.y = floatValue
	floatValue, err = strconv.ParseFloat(strArray[2], 64)
	if err != nil {
		return nil, err
	}
	p.z = floatValue
	p.circuit = []*point{p}
	return p, nil
}

func ReadFileAsPointArray(fileName string) ([]*point, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var pointArray []*point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strNumbers := strings.Split(scanner.Text(), ",")
		point, parseErr := StringArrayToPoint(strNumbers)
		if parseErr != nil {
			return nil, parseErr
		}
		pointArray = append(pointArray, point)
	}

	return pointArray, nil
}

func (p1 point) distance(p2 *point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2) + math.Pow(p1.z-p2.z, 2))
}

func (p point) print() {
	fmt.Printf("point [x=%f][y=%f][z=%f]\n", p.x, p.y, p.z)
}

func (p point) nearestPoint(points []*point, thisPoint int) (int, float64) {
	nearestPoint := 0
	lowestDistance := 100000000.0
	for i := 0; i < len(points); i++ {
		if i == thisPoint {
			continue
		}
		distance := p.distance(points[i])
		if distance < lowestDistance {
			lowestDistance = distance
			nearestPoint = i
		}
	}
	return nearestPoint, lowestDistance
}

func CalculateNearestPair(points []*point, pointsToSkip []bool) (int, int) {
	for i := 0; i < len(points); i++ {
		if pointsToSkip[i] {
			continue
		}
		for j := i + 1; j < len(points); j++ {
			if pointsToSkip[j] {
				continue
			}
			points[i].distance(points[j])
		}
	}
}

func main() {
	points, err := ReadFileAsPointArray("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(points); i++ {

	}
}
