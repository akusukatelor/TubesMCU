package pasien

import (
	"Tubes-MCU/data"
	"Tubes-MCU/paket"
	"Tubes-MCU/utils"
	"fmt"
)

func tambahDataPasien() {
	if data.JumlahPasien >= data.MAXDATA {
		fmt.Println("Data pasien penuh!")
		return
	}
	var pasien data.Pasien
	var namaPaket string
	fmt.Print("Masukkan Nama Pasien: ")
	utils.HandleLongInput(&pasien.Nama)
	fmt.Print("Masukkan Jenis Kelamin (L/K): ")
	fmt.Scan(&pasien.JenisKelamin)
	fmt.Print("Masukkan Umur Pasien: ")
	fmt.Scan(&pasien.Umur)
	fmt.Print("Masukkan Waktu MCU (1-31): ")
	fmt.Scan(&pasien.WaktuMCU)
	fmt.Print("Masukkan Nama Paket MCU yang Dipilih: ")
	utils.HandleLongInput(&namaPaket)
	for i := 0; i < data.JumlahPaket; i++ {
		if data.PaketMCUList[i].NamaPaket == namaPaket {
			pasien.PaketMCU = data.PaketMCUList[i]
			break
		}
	}
	fmt.Print("Masukkan Rekap Hasil MCU: ")
	utils.HandleLongInput(&pasien.HasilRekap)
	data.PasienList[data.JumlahPasien] = pasien
	data.JumlahPasien++
	fmt.Println("Data pasien berhasil ditambahkan!")
}

func editDataPasien() {
	var namaPasien string
	fmt.Print("Masukkan Nama Pasien yang Ingin Diedit: ")
	utils.HandleLongInput(&namaPasien)
	for i := 0; i < data.JumlahPasien; i++ {
		if data.PasienList[i].Nama == namaPasien {
			fmt.Print("Masukkan Nama Pasien Baru: ")
			utils.HandleLongInput(&data.PasienList[i].Nama)
			fmt.Print("Masukkan Jenis Kelamin Baru (L/K): ")
			fmt.Scan(&data.PasienList[i].JenisKelamin)
			fmt.Print("Masukkan Umur Baru: ")
			fmt.Scan(&data.PasienList[i].Umur)
			fmt.Print("Masukkan Waktu MCU Baru (1-31): ")
			fmt.Scan(&data.PasienList[i].WaktuMCU)
			fmt.Print("Masukkan Rekap Hasil MCU Baru: ")
			utils.HandleLongInput(&data.PasienList[i].HasilRekap)
			fmt.Println("Data pasien berhasil diubah!")
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan.")
}

func hapusDataPasien() {
	var namaPasien string
	fmt.Print("Masukkan Nama Pasien yang Ingin Dihapus: ")
	utils.HandleLongInput(&namaPasien)
	for i := 0; i < data.JumlahPasien; i++ {
		if data.PasienList[i].Nama == namaPasien {
			for j := i; j < data.JumlahPasien-1; j++ {
				data.PasienList[j] = data.PasienList[j+1]
			}
			data.JumlahPasien--
			fmt.Println("Data pasien berhasil dihapus!")
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan.")
}

func laporanPemasukan(tanggal int) int {
	var total int

	total = 0
	for i := 0; i < data.JumlahPasien; i++ {
		if data.PasienList[i].WaktuMCU == tanggal {
			total += data.PasienList[i].PaketMCU.Harga
		}
	}
	return total
}

func cariPasienBerdasarkanPaket(jumlahPasien *int) {
	var namaPaket string
	fmt.Print("Masukkan Nama Paket MCU untuk Mencari Pasien: ")
	utils.HandleLongInput(&namaPaket)
	found := false

	for i := 0; i < *jumlahPasien; i++ {
		if data.PasienList[i].PaketMCU.NamaPaket == namaPaket {
			fmt.Printf(
				"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
				data.PasienList[i].Nama,
				data.PasienList[i].JenisKelamin,
				data.PasienList[i].Umur,
				data.PasienList[i].PaketMCU.NamaPaket,
				float64(data.PasienList[i].PaketMCU.Harga),
				data.PasienList[i].WaktuMCU,
				data.PasienList[i].HasilRekap,
			)

			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada pasien yang memilih paket tersebut.")
	}
}

func cariPasienBerdasarkanWaktu() {
	var tanggal int
	fmt.Print("Masukkan Tanggal MCU (1-31): ")
	fmt.Scan(&tanggal)
	found := false

	for i := 0; i < data.JumlahPasien; i++ {
		if data.PasienList[i].WaktuMCU == tanggal {
			fmt.Printf(
				"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
				data.PasienList[i].Nama,
				data.PasienList[i].JenisKelamin,
				data.PasienList[i].Umur,
				data.PasienList[i].PaketMCU.NamaPaket,
				float64(data.PasienList[i].PaketMCU.Harga),
				data.PasienList[i].WaktuMCU,
				data.PasienList[i].HasilRekap,
			)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada pasien pada tanggal tersebut.")
	}
}

func cariPasienBerdasarkanNamaRekursif(nama string, low, high int) {
	if low > high {
		fmt.Println("Pasien tidak ditemukan.")
		return
	}
	mid := (low + high) / 2
	if data.PasienList[mid].Nama == nama {
		fmt.Printf(
			"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
			data.PasienList[mid].Nama,
			data.PasienList[mid].JenisKelamin,
			data.PasienList[mid].Umur,
			data.PasienList[mid].PaketMCU.NamaPaket,
			float64(data.PasienList[mid].PaketMCU.Harga),
			data.PasienList[mid].WaktuMCU,
			data.PasienList[mid].HasilRekap,
		)
		return
	} else if data.PasienList[mid].Nama < nama {
		cariPasienBerdasarkanNamaRekursif(nama, mid+1, high)
	} else {
		cariPasienBerdasarkanNamaRekursif(nama, low, mid-1)
	}
}

func cariPasienBerdasarkanNama() {
	selectionSortNama()
	var nama string
	fmt.Print("Masukkan Nama Pasien: ")
	utils.HandleLongInput(&nama)
	cariPasienBerdasarkanNamaRekursif(nama, 0, data.JumlahPasien-1)
}
func selectionSortNama() {
	for i := 0; i < data.JumlahPasien-1; i++ {
		minIndex := i
		for j := i + 1; j < data.JumlahPasien; j++ {
			if data.PasienList[j].Nama < data.PasienList[minIndex].Nama {
				minIndex = j
			}
		}
		data.PasienList[i], data.PasienList[minIndex] = data.PasienList[minIndex], data.PasienList[i]
	}
}

func insertionSortWaktuAscending() {
	for i := 1; i < data.JumlahPasien; i++ {
		key := data.PasienList[i]
		j := i - 1
		for j >= 0 && data.PasienList[j].WaktuMCU > key.WaktuMCU {
			data.PasienList[j+1] = data.PasienList[j]
			j--
		}
		data.PasienList[j+1] = key
	}
	fmt.Println("Data Pasien Ascending Berdasarkan Waktu MCU:")
	for i := 0; i < data.JumlahPasien; i++ {
		fmt.Printf(
			"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
			data.PasienList[i].Nama,
			data.PasienList[i].JenisKelamin,
			data.PasienList[i].Umur,
			data.PasienList[i].PaketMCU.NamaPaket,
			float64(data.PasienList[i].PaketMCU.Harga),
			data.PasienList[i].WaktuMCU,
			data.PasienList[i].HasilRekap,
		)
	}
}

func insertionSortWaktuDescending() {
	for i := 1; i < data.JumlahPasien; i++ {
		key := data.PasienList[i]
		j := i - 1
		for j >= 0 && data.PasienList[j].WaktuMCU < key.WaktuMCU {
			data.PasienList[j+1] = data.PasienList[j]
			j--
		}
		data.PasienList[j+1] = key
	}
	fmt.Println("Data Pasien Descending Berdasarkan Waktu MCU:")
	for i := 0; i < data.JumlahPasien; i++ {
		fmt.Printf(
			"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
			data.PasienList[i].Nama,
			data.PasienList[i].JenisKelamin,
			data.PasienList[i].Umur,
			data.PasienList[i].PaketMCU.NamaPaket,
			float64(data.PasienList[i].PaketMCU.Harga),
			data.PasienList[i].WaktuMCU,
			data.PasienList[i].HasilRekap,
		)
	}
}

func insertionSortPaketAscending() {
	for i := 1; i < data.JumlahPasien; i++ {
		key := data.PasienList[i]
		j := i - 1
		for j >= 0 && data.PasienList[j].PaketMCU.Harga > key.PaketMCU.Harga {
			data.PasienList[j+1] = data.PasienList[j]
			j--
		}
		data.PasienList[j+1] = key
	}
	fmt.Println("Data Pasien Ascending Berdasarkan Harga Paket MCU:")
	for i := 0; i < data.JumlahPasien; i++ {
		fmt.Printf(
			"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
			data.PasienList[i].Nama,
			data.PasienList[i].JenisKelamin,
			data.PasienList[i].Umur,
			data.PasienList[i].PaketMCU.NamaPaket,
			float64(data.PasienList[i].PaketMCU.Harga),
			data.PasienList[i].WaktuMCU,
			data.PasienList[i].HasilRekap,
		)
	}
}

func insertionSortPaketDescending() {
	for i := 1; i < data.JumlahPasien; i++ {
		key := data.PasienList[i]
		j := i - 1
		for j >= 0 && data.PasienList[j].PaketMCU.Harga < key.PaketMCU.Harga {
			data.PasienList[j+1] = data.PasienList[j]
			j--
		}
		data.PasienList[j+1] = key
	}
	fmt.Println("Data Pasien Descending Berdasarkan Harga Paket MCU:")
	for i := 0; i < data.JumlahPasien; i++ {
		fmt.Printf(
			"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
			data.PasienList[i].Nama,
			data.PasienList[i].JenisKelamin,
			data.PasienList[i].Umur,
			data.PasienList[i].PaketMCU.NamaPaket,
			float64(data.PasienList[i].PaketMCU.Harga),
			data.PasienList[i].WaktuMCU,
			data.PasienList[i].HasilRekap,
		)
	}
}

func PenggunaRumahSakit() {
	for {
		var pil int
		fmt.Println("--- Menu MCU ---")
		fmt.Println("1. Tambah Paket MCU")
		fmt.Println("2. Edit Paket MCU")
		fmt.Println("3. Hapus Paket MCU")
		fmt.Println("4. Tambah Data Pasien")
		fmt.Println("5. Edit Data Pasien")
		fmt.Println("6. Hapus Data Pasien")
		fmt.Println("7. Laporan Pemasukan")
		fmt.Println("8. Cari berdasarkan Paket")
		fmt.Println("9. Cari berdasarkan waktu")
		fmt.Println("10. Cari berdasarkan nama")
		fmt.Println("11. Tampil ascending berdasarkan waktu")
		fmt.Println("12. Tampil descending berdasarkan waktu")
		fmt.Println("13. Tampil ascending berdasarkan paket")
		fmt.Println("14. Tampil descending berdasarkan paket")
		fmt.Println("15. keluar program")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pil)
		switch pil {
		case 1:
			paket.TambahPaketMCU()
		case 2:
			paket.EditPaketMCU()
		case 3:
			paket.HapusPaketMCU()
		case 4:
			tambahDataPasien()
		case 5:
			editDataPasien()
		case 6:
			hapusDataPasien()
		case 7:
			var tanggal int
			fmt.Print("Masukkan Tanggal (1-31) untuk Laporan Pemasukan: ")
			fmt.Scan(&tanggal)
			fmt.Print("Pemasukan pada tanggal ", tanggal, ": ")
			fmt.Print(laporanPemasukan(tanggal))
		case 8:
			cariPasienBerdasarkanPaket(&data.JumlahPasien)
		case 9:
			cariPasienBerdasarkanWaktu()
		case 10:
			cariPasienBerdasarkanNama()
		case 11:
			insertionSortWaktuAscending()
		case 12:
			insertionSortWaktuDescending()
		case 13:
			insertionSortPaketAscending()
		case 14:
			insertionSortPaketDescending()
		case 15:
			fmt.Println("Kembali ke menu utama.")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
