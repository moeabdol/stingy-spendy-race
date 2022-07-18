/*
This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; version 2.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, see <http://www.gnu.org/licenses/>.

Copyright (C) Mohammed Saed, 2022
*/
package main

import (
	"fmt"
	"time"
)

var money = 100

func stingy() {
	for i := 0; i < 1000; i++ {
		money += 10
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy done!")
}

func spendy() {
	for i := 0; i < 1000; i++ {
		money -= 10
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Spendy done!")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Money =", money)
}
