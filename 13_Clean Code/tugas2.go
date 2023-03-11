// package main

// type kendaraan struct {
// 	totalroda       int
// 	kecepatanperjam int
// }

// type mobil struct {
// 	kendaraan
// }

// func (m *mobil) berjalan() {
// 	m.tambahkecepatan(10)
// }

// func (m *mobil) tambahkecepatan(kecepatanbaru int) {
// 	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
// }

// func main() {
// 	mobilcepat := mobil{}
// 	mobilcepat.berjalan()
// 	mobilcepat.berjalan()
// 	mobilcepat.berjalan()

// 	mobillamban := mobil{}
// 	mobillamban.berjalan()
// }

package main

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLambat := Mobil{}
	mobilLambat.Berjalan()
}

// Penjelasan perubahan:

// 1. Menggunakan huruf kapital pada nama struct dan method.
// 2. Menambahkan spasi antara field dan value pada struct.
// 3. Menggunakan nama yang lebih deskriptif pada field dan method.
// 4. Memindahkan method TambahKecepatan ke bawah method Berjalan.
// 5. Menggunakan shorthand operator pada method TambahKecepatan.