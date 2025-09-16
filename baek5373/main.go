package main

import (
	"bufio"
	"fmt"
	"os"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

func main() {
	defer std.Flush()

	t := 0
	fmt.Fscan(std, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(std, &n)
		r := make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(std, &r[j])
		}
		solve(r)
	}
}

type Cube map[byte][]byte

func (cube Cube) rotateFace(r string) {
	face := cube[r[0]]
	if r[1] == '+' {
		temp := face[0]
		face[0] = face[6]
		face[6] = face[8]
		face[8] = face[2]
		face[2] = temp
		temp = face[1]
		face[1] = face[3]
		face[3] = face[7]
		face[7] = face[5]
		face[5] = temp
	} else {
		temp := face[0]
		face[0] = face[2]
		face[2] = face[8]
		face[8] = face[6]
		face[6] = temp
		temp = face[1]
		face[1] = face[5]
		face[5] = face[7]
		face[7] = face[3]
		face[3] = temp
	}
}

func solve(rotations []string) {
	cube := Cube{
		'U': []byte("wwwwwwwww"),
		'F': []byte("rrrrrrrrr"),
		'R': []byte("bbbbbbbbb"),
		'B': []byte("ooooooooo"),
		'L': []byte("ggggggggg"),
		'D': []byte("yyyyyyyyy"),
	}

	u := cube['U']
	f := cube['F']
	r := cube['R']
	b := cube['B']
	l := cube['L']
	d := cube['D']

	for _, v := range rotations {
		cube.rotateFace(v)

		switch v {
		case "U+":
			for i := range 3 {
				temp := b[i]
				b[i] = l[i]
				l[i] = f[i]
				f[i] = r[i]
				r[i] = temp
			}

		case "U-":
			for i := range 3 {
				temp := b[i]
				b[i] = r[i]
				r[i] = f[i]
				f[i] = l[i]
				l[i] = temp
			}

		case "F+":
			for i := range 3 {
				temp := u[6+i]
				u[6+i] = l[8-3*i]
				l[8-3*i] = d[2-i]
				d[2-i] = r[3*i]
				r[3*i] = temp
			}

		case "F-":
			for i := range 3 {
				temp := u[6+i]
				u[6+i] = r[3*i]
				r[3*i] = d[2-i]
				d[2-i] = l[8-3*i]
				l[8-3*i] = temp
			}

		case "R+":
			for i := range 3 {
				temp := u[8-3*i]
				u[8-3*i] = f[8-3*i]
				f[8-3*i] = d[8-3*i]
				d[8-3*i] = b[3*i]
				b[3*i] = temp
			}

		case "R-":
			for i := range 3 {
				temp := u[8-3*i]
				u[8-3*i] = b[3*i]
				b[3*i] = d[8-3*i]
				d[8-3*i] = f[8-3*i]
				f[8-3*i] = temp
			}

		case "B+":
			for i := range 3 {
				temp := u[2-i]
				u[2-i] = r[8-3*i]
				r[8-3*i] = d[6+i]
				d[6+i] = l[3*i]
				l[3*i] = temp
			}
		case "B-":
			for i := range 3 {
				temp := u[2-i]
				u[2-i] = l[3*i]
				l[3*i] = d[6+i]
				d[6+i] = r[8-3*i]
				r[8-3*i] = temp
			}
		case "L+":
			for i := range 3 {
				temp := u[3*i]
				u[3*i] = b[8-3*i]
				b[8-3*i] = d[3*i]
				d[3*i] = f[3*i]
				f[3*i] = temp
			}
		case "L-":
			for i := range 3 {
				temp := u[3*i]
				u[3*i] = f[3*i]
				f[3*i] = d[3*i]
				d[3*i] = b[8-3*i]
				b[8-3*i] = temp
			}

		case "D+":
			for i := range 3 {
				temp := f[6+i]
				f[6+i] = l[6+i]
				l[6+i] = b[6+i]
				b[6+i] = r[6+i]
				r[6+i] = temp
			}
		case "D-":
			for i := range 3 {
				temp := f[6+i]
				f[6+i] = r[6+i]
				r[6+i] = b[6+i]
				b[6+i] = l[6+i]
				l[6+i] = temp
			}
		}
	}
	fmt.Fprintln(std, string(cube['U'][0:3]))
	fmt.Fprintln(std, string(cube['U'][3:6]))
	fmt.Fprintln(std, string(cube['U'][6:9]))
}
