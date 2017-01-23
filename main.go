package main

import (
	"fmt"
	"regexp"
	"os"
	"strconv"
)

type position struct {
	x	int
	y	int
}

const (
	R	=	'R'
	L	=	'L'
	W	=	'W'
)

func main(){
	var deg = 90
	var pos position
	var input string
	input = os.Args[len(os.Args) - 1]

	for len(input) != 0{
		fString := input[0]

		if fString == R {
			turn(&deg, R)
		}else if fString == L {
			turn(&deg, L)
		}else if fString == W{
			r := regexp.MustCompile("[0-9]+")
			s := r.FindStringSubmatch(input)
			if len(s) != 0{
				step, _ := strconv.Atoi(s[0])
				walk(&pos, deg, step)
			}
		}

		input = input[1:]
	}

	fmt.Println("X:", pos.x, "Y:", pos.y, "Direction:", degToString(deg))
}

func degToString(deg int) string{
	var result string
	switch deg{
	case 0:
		result = "East"
	case 90:
		result = "North"
	case 180:
		result = "West"
	case 270:
		result = "South"
	}

	return result
}

func walk(pos *position, deg int, step int){
	switch deg{
	case 0:
		pos.x += step
	case 90:
		pos.y += step
	case 180:
		pos.x -= step
	case 270:
		pos.y -= step
	}
}

func turn(deg *int, how byte){
	if how == R {
		*deg -= 90
	}else if how == L{
		*deg += 90
	}

	if *deg >= 360{
		*deg -= 360
	}else if *deg < 0 {
		*deg += 360
	}
}
