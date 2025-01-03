package main

import (
	"Tubes-MCU/dummy"
	"Tubes-MCU/pasien"
	"fmt"
)

func main() {
	dummy.InitDummyData()
	var pilihan int
	for {
		fmt.Println("Selamat datang di Layanan Medical Check Up Ahlul Jannah")
		fmt.Println("1. Masuk sebagai Pengguna Rumah Sakit")
		fmt.Println("2. Masuk sebagai Pengguna Pasien")
		fmt.Println("3. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			pasien.PenggunaRumahSakit()
		case 2:
			dummy.PenggunaPasien()
		case 3:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
