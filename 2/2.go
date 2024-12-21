package main
import "fmt"

const nMax = 7919

type Buku struct {
	id int
	judul string
	penulis string
	penerbit string
	tahun int
	rating int
}
type DaftarBuku struct {
	pustaka [nMax]Buku
	n int
}
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	pustaka.n = n
	for i := 0; i < n; i++ {
		fmt.Println("Masukkan data (1 Kata) untuk buku ke-", i+1)
		fmt.Print("ID: ")
		fmt.Scan(&pustaka.pustaka[i].id)
		fmt.Print("Judul: ")
		fmt.Scan(&pustaka.pustaka[i].judul)
		fmt.Print("Penulis: ")
		fmt.Scan(&pustaka.pustaka[i].penulis)
		fmt.Print("Penerbit: ")
		fmt.Scan(&pustaka.pustaka[i].penerbit)
		fmt.Print("Tahun: ")
		fmt.Scan(&pustaka.pustaka[i].tahun)
		fmt.Print("Rating(1-5): ")
		fmt.Scan(&pustaka.pustaka[i].rating)
	}
}
func CetakTertfavorit(pustaka DaftarBuku) {
	if pustaka.n == 0 {
		fmt.Println("Tidak ada buku dalam daftar.")
		return
	}
	maxRating := pustaka.pustaka[0].rating
	for i := 1; i < pustaka.n; i++ {
		if pustaka.pustaka[i].rating > maxRating {
			maxRating = pustaka.pustaka[i].rating
		}
	}
	fmt.Println("\nBuku dengan rating tertinggi:")
	for i := 0; i < pustaka.n; i++ {
		if pustaka.pustaka[i].rating == maxRating {
			CetakBuku(pustaka.pustaka[i])
		}
	}
}
func CetakBuku(b Buku) {
	fmt.Printf("ID       : %d\n", b.id)
	fmt.Printf("Judul    : %s\n", b.judul)
	fmt.Printf("Penulis  : %s\n", b.penulis)
	fmt.Printf("Penerbit : %s\n", b.penerbit)
	fmt.Printf("Tahun    : %d\n", b.tahun)
	fmt.Printf("Rating   : %d\n", b.rating)
	fmt.Println("-------------------------------")
}

func CetakBukuRating(pustaka DaftarBuku, rating int) {
	found := false
	fmt.Printf("\nBuku dengan rating %d:\n", rating)
	for i := 0; i < pustaka.n; i++ {
		if pustaka.pustaka[i].rating == rating {
			found = true
			CetakBuku(pustaka.pustaka[i])
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku
	var jumlahBuku int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&jumlahBuku)

	DaftarkanBuku(&pustaka, jumlahBuku)
	CetakTertfavorit(pustaka)

	var rating int
	fmt.Print("Masukkan rating untuk mencari buku: ")
	fmt.Scan(&rating)

	CetakBukuRating(pustaka, rating)
}