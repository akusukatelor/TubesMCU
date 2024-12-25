package main

import "fmt"

const maxData = 100

type PaketMCU struct {
	NamaPaket string
	Harga     int
}

type Pasien struct {
	Nama         string
	JenisKelamin string
	Umur         int
	PaketMCU     PaketMCU
	WaktuMCU     int 
	HasilRekap   string
}

var (
	paketMCUList [maxData]PaketMCU
	pasienList   [maxData]Pasien
	jumlahPaket  int
	jumlahPasien int
)

func tambahPaketMCU() {
	if jumlahPaket >= maxData {
		fmt.Println("Data paket MCU penuh!")
		return
	}
	var paket PaketMCU
	fmt.Print("Masukkan Nama Paket MCU: ") 
	fmt.Scan(&paket.NamaPaket)
	fmt.Print("Masukkan Harga Paket MCU: ")
	fmt.Scan(&paket.Harga)
	paketMCUList[jumlahPaket] = paket
	jumlahPaket++
	fmt.Println("Paket MCU berhasil ditambahkan!")
}

func editPaketMCU() {
	var namaPaket string
	fmt.Print("Masukkan Nama Paket MCU yang Ingin Diedit: ")
	fmt.Scan(&namaPaket)
	for i := 0; i < jumlahPaket; i++ {
		if paketMCUList[i].NamaPaket == namaPaket {
			fmt.Print("Masukkan Nama Paket MCU Baru: ")
			fmt.Scan(&paketMCUList[i].NamaPaket)
			fmt.Print("Masukkan Harga Paket MCU Baru: ")
			fmt.Scan(&paketMCUList[i].Harga)
			fmt.Println("Paket MCU berhasil diubah!")
			return
		}
	}
	fmt.Println("Paket MCU tidak ditemukan.")
}

func hapusPaketMCU() {
	var namaPaket string
	fmt.Print("Masukkan Nama Paket MCU yang Ingin Dihapus: ")
	fmt.Scan(&namaPaket)
	for i := 0; i < jumlahPaket; i++ {
		if paketMCUList[i].NamaPaket == namaPaket {
			for j := i; j < jumlahPaket-1; j++ {
				paketMCUList[j] = paketMCUList[j+1]
			}
			jumlahPaket--
			fmt.Println("Paket MCU berhasil dihapus!")
			return
		}
	}
	fmt.Println("Paket MCU tidak ditemukan.")
}

func tambahDataPasien() {
	if jumlahPasien >= maxData {
		fmt.Println("Data pasien penuh!")
		return
	}
	var pasien Pasien
	var namaPaket string
	fmt.Print("Masukkan Nama Pasien: ")
	fmt.Scan(&pasien.Nama)
	fmt.Print("Masukkan Jenis Kelamin (L/K): ")
	fmt.Scan(&pasien.JenisKelamin)
	fmt.Print("Masukkan Umur Pasien: ")
	fmt.Scan(&pasien.Umur)
	fmt.Print("Masukkan Waktu MCU (1-31): ")
	fmt.Scan(&pasien.WaktuMCU)
	fmt.Print("Masukkan Nama Paket MCU yang Dipilih: ")
	fmt.Scan(&namaPaket)
	for i := 0; i < jumlahPaket; i++ {
		if paketMCUList[i].NamaPaket == namaPaket {
			pasien.PaketMCU = paketMCUList[i]
			break
		}
	}
	fmt.Print("Masukkan Rekap Hasil MCU: ")
	fmt.Scan(&pasien.HasilRekap)
	pasienList[jumlahPasien] = pasien
	jumlahPasien++
	fmt.Println("Data pasien berhasil ditambahkan!")
}

func editDataPasien() {
	var namaPasien string
	fmt.Print("Masukkan Nama Pasien yang Ingin Diedit: ")
	fmt.Scan(&namaPasien)
	for i := 0; i < jumlahPasien; i++ {
		if pasienList[i].Nama == namaPasien {
			fmt.Print("Masukkan Nama Pasien Baru: ")
			fmt.Scan(&pasienList[i].Nama)
			fmt.Print("Masukkan Jenis Kelamin Baru (L/K): ")
			fmt.Scan(&pasienList[i].JenisKelamin)
			fmt.Print("Masukkan Umur Baru: ")
			fmt.Scan(&pasienList[i].Umur)
			fmt.Print("Masukkan Waktu MCU Baru (1-31): ")
			fmt.Scan(&pasienList[i].WaktuMCU)
			fmt.Print("Masukkan Rekap Hasil MCU Baru: ")
			fmt.Scan(&pasienList[i].HasilRekap)
			fmt.Println("Data pasien berhasil diubah!")
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan.")
}

func hapusDataPasien() {
	var namaPasien string
	fmt.Print("Masukkan Nama Pasien yang Ingin Dihapus: ")
	fmt.Scan(&namaPasien)
	for i := 0; i < jumlahPasien; i++ {
		if pasienList[i].Nama == namaPasien {
			for j := i; j < jumlahPasien-1; j++ {
				pasienList[j] = pasienList[j+1]
			}
			jumlahPasien--
			fmt.Println("Data pasien berhasil dihapus!")
			return
		}
	}
	fmt.Println("Pasien tidak ditemukan.")
}
func laporanPemasukan() {
	var tanggal, total int
	fmt.Println("Masukkan Tanggal (1-31) untuk Laporan Pemasukan: ")
	fmt.Scan(&tanggal)

	total = 0
	for i := 0; i < jumlahPasien; i++ {
		if pasienList[i].WaktuMCU == tanggal {
			total += pasienList[i].PaketMCU.Harga
		}
	}
	fmt.Printf("Total Pemasukan pada Tanggal %d: %d\n", tanggal, total)
}

func cariPasienBerdasarkanPaket() {
	var namaPaket string
	fmt.Println("Masukkan Nama Paket MCU untuk Mencari Pasien: ")
	fmt.Scan(&namaPaket)
	found := false

	for i := 0; i < jumlahPasien; i++ {
		if pasienList[i].PaketMCU.NamaPaket == namaPaket {
			fmt.Println("Nama: ", pasienList[i].Nama, " | Tanggal MCU: ", pasienList[i].WaktuMCU)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada pasien yang memilih paket tersebut.")
	}
}

func cariPasienBerdasarkanWaktu() {
	var tanggal int
	fmt.Println("Masukkan Tanggal MCU (1-31): ")
	fmt.Scan(&tanggal)
	found := false

	for i := 0; i < jumlahPasien; i++ {
		if pasienList[i].WaktuMCU == tanggal {
			fmt.Printf("Nama: %s, Paket: %s\n", pasienList[i].Nama, pasienList[i].PaketMCU.NamaPaket)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada pasien pada tanggal tersebut.")
	}
}

func cariPasienBerdasarkanNama() {
	selectionSortNama()
	var nama string
	fmt.Println("Masukkan Nama Pasien: ")
	fmt.Scan(&nama)

	low, high := 0, jumlahPasien-1
	for low <= high {
		mid := (low + high) / 2
		if pasienList[mid].Nama == nama {
			fmt.Printf("Ditemukan: Nama: %s, Paket: %s\n", pasienList[mid].Nama, pasienList[mid].PaketMCU.NamaPaket)
			return
		} else if pasienList[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Pasien tidak ditemukan.")
}

func selectionSortNama() {
	for i := 0; i < jumlahPasien-1; i++ {
		minIndex := i
		for j := i + 1; j < jumlahPasien; j++ {
			if pasienList[j].Nama < pasienList[minIndex].Nama {
				minIndex = j
			}
		}
		pasienList[i], pasienList[minIndex] = pasienList[minIndex], pasienList[i]
	}
}

func insertionSortWaktuAscending() {
	for i := 1; i < jumlahPasien; i++ {
		key := pasienList[i]
		j := i - 1
		for j >= 0 && pasienList[j].WaktuMCU > key.WaktuMCU {
			pasienList[j+1] = pasienList[j]
			j--
		}
		pasienList[j+1] = key
	}
	fmt.Println("Data Pasien Ascending Berdasarkan Waktu MCU:")
	for i := 0; i < jumlahPasien; i++ {
		fmt.Printf("Nama: %s, Waktu MCU: %d\n", pasienList[i].Nama, pasienList[i].WaktuMCU)
	}
}

func insertionSortWaktuDescending() {
	for i := 1; i < jumlahPasien; i++ {
		key := pasienList[i]
		j := i - 1
		for j >= 0 && pasienList[j].WaktuMCU < key.WaktuMCU {
			pasienList[j+1] = pasienList[j]
			j--
		}
		pasienList[j+1] = key
	}
	fmt.Println("Data Pasien Descending Berdasarkan Waktu MCU:")
	for i := 0; i < jumlahPasien; i++ {
		fmt.Printf("Nama: %s, Waktu MCU: %d\n", pasienList[i].Nama, pasienList[i].WaktuMCU)
	}
}

func insertionSortPaketAscending() {
	for i := 1; i < jumlahPasien; i++ {
		key := pasienList[i]
		j := i - 1
		for j >= 0 && pasienList[j].PaketMCU.Harga > key.PaketMCU.Harga {
			pasienList[j+1] = pasienList[j]
			j--
		}
		pasienList[j+1] = key
	}
	fmt.Println("Data Pasien Ascending Berdasarkan Harga Paket MCU:")
	for i := 0; i < jumlahPasien; i++ {
		fmt.Printf("Nama: %s, Harga Paket: %d\n", pasienList[i].Nama, pasienList[i].PaketMCU.Harga)
	}
}

func insertionSortPaketDescending() {
	for i := 1; i < jumlahPasien; i++ {
		key := pasienList[i]
		j := i - 1
		for j >= 0 && pasienList[j].PaketMCU.Harga < key.PaketMCU.Harga {
			pasienList[j+1] = pasienList[j]
			j--
		}
		pasienList[j+1] = key
	}
	fmt.Println("Data Pasien Descending Berdasarkan Harga Paket MCU:")
	for i := 0; i < jumlahPasien; i++ {
		fmt.Printf("Nama: %s, Harga Paket: %d\n", pasienList[i].Nama, pasienList[i].PaketMCU.Harga)
	}
}

func main() {
	var pilihan int
	for {
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
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahPaketMCU()
		case 2:
			editPaketMCU()
		case 3:
			hapusPaketMCU()
		case 4:
			tambahDataPasien()
		case 5:
			editDataPasien()
		case 6:
			hapusDataPasien()
		case 7:
			laporanPemasukan()
		case 8:
			cariPasienBerdasarkanPaket()
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
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
