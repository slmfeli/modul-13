package main

import "fmt"

func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func cekBerjarak(arr []int, n int) string {
	if n < 2 {
		return "Data berjarak tetap"
	}

	selisih := arr[1] - arr[0]
	if selisih < 0 {
		selisih = -selisih
	}

	for i := 1; i < n-1; i++ {
		currSelisih := arr[i+1] - arr[i]
		if currSelisih < 0 {
			currSelisih = -currSelisih
		}
		if currSelisih != selisih {
			return "Data berjarak tidak tetap"
		}
	}

	return fmt.Sprintf("Data berjarak %d", selisih)
}

func main() {
	var input [100]int
	var n int

	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")
	for {
		var bil int
		fmt.Scan(&bil)
		if bil < 0 {
			break
		}
		input[n] = bil
		n++
	}

	if n == 0 {
		fmt.Println("Tidak ada data untuk diproses.")
		return
	}

	insertionSort(input[:], n)

	status := cekBerjarak(input[:], n)

	fmt.Print("Array setelah diurutkan: ")
	for i := 0; i < n; i++ {
		fmt.Print(input[i], " ")
	}
	fmt.Println()
	fmt.Println(status)
}