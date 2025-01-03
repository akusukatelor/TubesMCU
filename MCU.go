package paket

import (
	"Tubes-MCU/data"
	"Tubes-MCU/utils"
	"fmt"
)

func TambahPaketMCU() {
	if data.JumlahPaket >= data.MAXDATA {
		fmt.Println("Data paket MCU penuh!")
		return
	}
	var paket data.PaketMCU
	fmt.Print("Masukkan Nama Paket MCU: ")
	utils.HandleLongInput(&paket.NamaPaket)
	fmt.Print("Masukkan Harga Paket MCU: (K)")
	fmt.Scan(&paket.Harga)
	data.PaketMCUList[data.JumlahPaket] = paket
	data.JumlahPaket++
	fmt.Println("Paket MCU berhasil ditambahkan!")
}
func EditPaketMCU() {
	var namaPaket string
	fmt.Print("Masukkan Nama Paket MCU yang Ingin Diedit: ")
	utils.HandleLongInput(&namaPaket)
	for i := 0; i < data.JumlahPaket; i++ {
		if data.PaketMCUList[i].NamaPaket == namaPaket {
			fmt.Print("Masukkan Nama Paket MCU Baru: ")
			utils.HandleLongInput(&data.PaketMCUList[i].NamaPaket)
			fmt.Print("Masukkan Harga Paket MCU Baru: (K)")
			fmt.Scan(&data.PaketMCUList[i].Harga)
			fmt.Println("Paket MCU berhasil diubah!")
			return
		}
	}
	fmt.Println("Paket MCU tidak ditemukan.")
}

func HapusPaketMCU() {
	var namaPaket string
	fmt.Print("Masukkan Nama Paket MCU yang Ingin Dihapus: ")
	utils.HandleLongInput(&namaPaket)
	for i := 0; i < data.JumlahPaket; i++ {
		if data.PaketMCUList[i].NamaPaket == namaPaket {
			for j := i; j < data.JumlahPaket-1; j++ {
				data.PaketMCUList[j] = data.PaketMCUList[j+1]
			}
			data.JumlahPaket--
			fmt.Println("Paket MCU berhasil dihapus!")
			return
		}
	}
	fmt.Println("Paket MCU tidak ditemukan.")
}
