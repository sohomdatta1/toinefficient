package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func a(n int, k int) uint64 {
	if k > n {
		return 0
	} else if k == 0 || k == n {
		return 1
	}

	// Recur
	return a(n-1, k-1) + a(n-1, k)
}

func b() {
	fmt.Printf("USAGE: %s [name_of_file]\n", os.Args[0])
}

func d(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(0)
	}
}

func f() {
	fmt.Println("flag{I_am_not_a_flag}... Lokk deeper inside the code")
}

func g(h byte, j byte) byte {
	j = j % 8
	temp := h<<j | h>>(8-j)
	return temp
}

func e(b_arr []byte, c int) byte {
	tot := a(c, int(c/2))
	tot_str := strconv.FormatUint(tot, 10)
	offset := tot_str[rand.Intn(len(tot_str))] - '0'
	return g(b_arr[0], offset)
}

func c() {
	f, err := os.Open(os.Args[1])
	d(err)
	f_wr, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0666)
	d(err)
	cen := 10
	i := 1

	for true {
		b_arr := make([]byte, 1)
		_, err = f.Read(b_arr)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}
		d(err)
		new_b := e(b_arr, cen)
		_, err = f_wr.Write([]byte{new_b})
		d(err)
		cen = (cen + 1) % 70
		i = i + 1
	}

	f.Close()
	f_wr.Close()
}

func main() {
	if len(os.Args) < 2 {
		b()
	}
	rand.Seed(int64(0))

	fmt.Println("Decrypting the flag")
	fmt.Println("This may take a while...")
	c()
}
