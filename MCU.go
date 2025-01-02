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
type Dokter struct {
	ID        int
	Nama      string
	Spesialis string
	Jadwal    string
}

var (
	paketMCUList [maxData]PaketMCU
	listDokter   [maxData]Dokter
	pasienList   [maxData]Pasien
	jumlahPaket  int
	jumlahPasien int
)

func penggunaRumahSakit() {
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
			var tanggal int
			fmt.Print("Masukkan Tanggal (1-31) untuk Laporan Pemasukan: ")
			fmt.Scan(&tanggal)
			fmt.Print("Pemasukan pada tanggal ", tanggal, ": ")
			fmt.Print(laporanPemasukan(tanggal))
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
			fmt.Println("Kembali ke menu utama.")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
func tanyaPenyakit() {
	for {
		var pilihan int
		fmt.Println("Tanyakan Penyakit:")
		fmt.Println("1. Demam")
		fmt.Println("2. Pusing")
		fmt.Println("3. Sakit Mata")
		fmt.Println("4. Diare")
		fmt.Println("5. Sakit Telinga")
		fmt.Print("Pilihan Anda: ")

		fmt.Scan(&pilihan)
		var ada string

		switch pilihan {
		case 1:
			fmt.Print("Sudah berapa lama Anda mengalami gejala ini? (hari): ")
			var hari int
			fmt.Scan(&hari)

			if hari < 4 {
				fmt.Println(listDokter[0].Nama, ": Untuk solusi gejala awal: \nIstirahat yang cukup: Tubuh membutuhkan energi untuk melawan infeksi.\nMinum banyak cairan: Air, teh hangat, atau sup kaldu membantu mencegah dehidrasi.\nKompres hangat: Tempelkan kain basah hangat di dahi untuk meredakan ketidaknyamanan.\nGunakan pakaian tipis: Hindari pakaian yang terlalu tebal agar panas tubuh dapat keluar.\nObat penurun demam: Paracetamol atau ibuprofen dapat digunakan sesuai dosis untuk menurunkan suhu tubuh.")
			} else {
				fmt.Println(listDokter[0].Nama, ": Jika demam anda sudah lebih dari 3 hari maka disarankan untuk pergi ke rumah sakit langsung, akan tetapi untuk pertolongan awal anda bisa meminum obat seperti paracetamol atau ibuprofen dan disertai istirahat total.")
			}

			fmt.Print("Apakah ada keluhan lainnya? (y/n):")
			fmt.Scan(&ada)
			if ada == "y" {
				continue
			} else {
				fmt.Println("Terima kasih, semoga lekas sembuh :)")
				return
			}
		case 2:
			fmt.Print("Sudah berapa lama Anda mengalami gejala ini? (hari): ")
			var hari int
			fmt.Scan(&hari)

			if hari < 4 {
				fmt.Println(listDokter[0].Nama, ": Jika sakit yang anda rasakan kurang dari 4 hari,mungkin yang anda bisa lakukan:\n1.Minum Air Putih,Dehidrasi sering menjadi penyebab utama pusing. Minumlah segelas air putih dan pastikan hidrasi kamu cukup sepanjang hari.\n2. Istirahat Sejenak, Matikan layar, pejamkan mata, dan istirahat selama 10–15 menit di tempat tenang. Hindari cahaya terang dan suara bising.\n3.  Makan atau Minum yang Manis, Pusing bisa disebabkan oleh gula darah rendah. Konsumsi makanan ringan seperti roti, buah, atau minuman hangat (teh manis/jahe).")
			} else {
				fmt.Println(listDokter[0].Nama, ": Jika lebih dari 3 hari anda sudah melakukan penanganan awal seperti minum obat,istirahat,dll masih belum sembuh. Dan kika pusing disertai gejala seperti mual parah, muntah, lemas, pandangan kabur, disarankan untuk segera pergi ke rumah sakit langsung.")
			}

			fmt.Print("Apakah ada keluhan lainnya? (y/n):")
			fmt.Scan(&ada)
			if ada == "y" {
				continue
			} else {
				fmt.Println("Terima kasih, semoga lekas sembuh :)")
				return
			}
		case 3:
			fmt.Print("Sudah berapa lama Anda mengalami gejala ini? (hari): ")
			var hari int
			fmt.Scan(&hari)

			if hari < 4 {
				fmt.Println(listDokter[2].Nama, ": Untuk gejala awal mungkin anda bisa lakukan: \n1.  Istirahatkan Mata,Hindari menatap layar komputer, ponsel, atau TV terlalu lama. Gunakan aturan 20-20-20: setiap 20 menit, alihkan pandangan ke sesuatu sejauh 20 kaki (6 meter) selama 20 detik.\n2. Kompres Hangat atau Dingin, kompres hangat untuk mata lelah atau bengkak, kompres dingin untuk mata merah,gatal atau meradang.\n3. Gunakan Obat Tetes Mata, Mata Kering gunakan artificial tears (air mata buatan), Mata Merah, Pilih tetes mata dengan decongestant (misal: Visine).")
			} else {
				fmt.Println(listDokter[2].Nama, ": Jika mata kamu lebih dari 3 hari masih sakit setelah sudah diberi penanganan awal dan terasa tanda berikut: Nyeri parah atau bengkak berat, Penglihatan buram atau kabur, Keluar cairan abnormal seperti nanah. Disarankan langsung segera ke rumah sakit untuk penanganan lebih lanjut.")
			}
			fmt.Print("Apakah ada keluhan lainnya? (y/n):")
			fmt.Scan(&ada)
			if ada == "y" {
				continue
			} else {
				fmt.Println("Terima kasih, semoga lekas sembuh :)")
				return
			}
		case 4:
			fmt.Print("Sudah berapa lama Anda mengalami gejala ini? (hari): ")
			var hari int
			fmt.Scan(&hari)

			if hari < 4 {
				fmt.Println(listDokter[0].Nama, ": Diare yang berlangsung kurang dari 4 hari umumnya dianggap sebagai diare akut. Berikut adalah beberapa solusi yang bisa dilakukan:\nMinum cairan lebih banyak: Untuk menggantikan cairan yang hilang, minum air putih, oralit (larutan rehidrasi oral), atau cairan elektrolit lain.\nHindari dehidrasi: Gejala dehidrasi meliputi mulut kering, rasa haus, lemas, atau warna urin yang lebih gelap.")
			} else {
				fmt.Println(listDokter[0].Nama, ": Segera pergi ke dokter jika:\nDiare berlangsung lebih dari 4 hari.\nAda tanda dehidrasi berat (seperti mata cekung, tidak buang air kecil).\nAda darah dalam tinja atau tinja berwarna hitam.Disertai demam tinggi atau nyeri perut hebat.")
			}
		case 5:
			fmt.Print("Sudah berapa lama Anda mengalami gejala ini? (hari): ")
			var hari int
			fmt.Scan(&hari)

			if hari < 4 {
				fmt.Println(listDokter[1].Nama, ": Untuk penangan awal:\nKompres hangat: Tempelkan kain hangat (tidak terlalu panas) pada telinga yang sakit selama 10-15 menit untuk mengurangi nyeri.\nKompres dingin: Jika nyeri disertai pembengkakan, gunakan kain yang dibungkus es batu sebagai kompres.")
			} else {
				fmt.Println(listDokter[1].Nama, ": Segera pergi ke dokter jika:\nNyeri telinga berlangsung lebih dari 4 hari atau semakin parah.\nAda cairan keluar dari telinga (seperti nanah atau darah).Disertai demam tinggi, kehilangan pendengaran, atau pusing.Anak kecil sangat rewel, sering menarik-narik telinga, atau tidak mau makan.")
			}
		}
	}
}

func penggunaPasien() {
	for {
		fmt.Println("\nMenu Pengguna Terdaftar:")
		fmt.Println("1. Tanyakan Penyakit")
		fmt.Println("2. Lihat Jadwal Dokter")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilihan Anda: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tanyaPenyakit()
		case 2:
			lihatJadwalDokter()
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
func tambahPaketMCU() {
	if jumlahPaket >= maxData {
		fmt.Println("Data paket MCU penuh!")
		return
	}
	var paket PaketMCU
	fmt.Print("Masukkan Nama Paket MCU: ")
	fmt.Scan(&paket.NamaPaket)
	fmt.Print("Masukkan Harga Paket MCU: (K)")
	fmt.Scan(&paket.Harga)
	paketMCUList[jumlahPaket] = paket
	jumlahPaket++
	fmt.Println("Paket MCU berhasil ditambahkan!")
}
func initDummyData() {
	paketMCUList[0] = PaketMCU{"Basic", 500}
	paketMCUList[1] = PaketMCU{"Advanced", 100}
	paketMCUList[2] = PaketMCU{"Premium", 200}
	jumlahPaket = 3
	pasienList[0] = Pasien{"Alice", "P", 25, paketMCUList[0], 15, "Sehat"}
	pasienList[1] = Pasien{"Bob", "L", 30, paketMCUList[1], 16, "Sehat"}
	pasienList[2] = Pasien{"Charlie", "L", 40, paketMCUList[2], 17, "Hipertensi"}
	jumlahPasien = 3
}
func lihatJadwalDokter() {
	listDokter[0].ID = 001
	listDokter[0].Nama = "Dr.Alfin God"
	listDokter[0].Spesialis = "Umum"
	listDokter[0].Jadwal = "Senin-Jumat"
	listDokter[1].ID = 002
	listDokter[1].Nama = "Dr.Raqa Anshor"
	listDokter[1].Spesialis = "THT"
	listDokter[1].Jadwal = "Sabtu"
	listDokter[2].ID = 003
	listDokter[2].Nama = "Dr.Asbi hihi"
	listDokter[2].Spesialis = "Mata"
	listDokter[2].Jadwal = "Minggu"

	for i := 0; i < 3; i++ {
		fmt.Println("ID: ", listDokter[i].ID, ", Nama: ", listDokter[i].Nama, ", Spesialis: ", listDokter[i].Spesialis, ", Jadwal: ", listDokter[i].Jadwal)
	}
}

func editPaketMCU() {
	var namaPaket string
	fmt.Print("Masukkan Nama Paket MCU yang Ingin Diedit: ")
	fmt.Scan(&namaPaket)
	for i := 0; i < jumlahPaket; i++ {
		if paketMCUList[i].NamaPaket == namaPaket {
			fmt.Print("Masukkan Nama Paket MCU Baru: ")
			fmt.Scan(&paketMCUList[i].NamaPaket)
			fmt.Print("Masukkan Harga Paket MCU Baru: (K)")
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
func laporanPemasukan(tanggal int) int {
	var total int

	total = 0
	for i := 0; i < jumlahPasien; i++ {
		if pasienList[i].WaktuMCU == tanggal {
			total += pasienList[i].PaketMCU.Harga
		}
	}
	return total
}

func cariPasienBerdasarkanPaket() {
	var namaPaket string
	fmt.Print("Masukkan Nama Paket MCU untuk Mencari Pasien: ")
	fmt.Scan(&namaPaket)
	found := false

	for i := 0; i < jumlahPasien; i++ {
		if pasienList[i].PaketMCU.NamaPaket == namaPaket {
			fmt.Printf(
				"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
				pasienList[i].Nama,
				pasienList[i].JenisKelamin,
				pasienList[i].Umur,
				pasienList[i].PaketMCU.NamaPaket,
				float64(pasienList[i].PaketMCU.Harga),
				pasienList[i].WaktuMCU,
				pasienList[i].HasilRekap,
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

	for i := 0; i < jumlahPasien; i++ {
		if pasienList[i].WaktuMCU == tanggal {
			fmt.Printf(
				"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
				pasienList[i].Nama,
				pasienList[i].JenisKelamin,
				pasienList[i].Umur,
				pasienList[i].PaketMCU.NamaPaket,
				float64(pasienList[i].PaketMCU.Harga),
				pasienList[i].WaktuMCU,
				pasienList[i].HasilRekap,
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
	if pasienList[mid].Nama == nama {
		fmt.Printf(
			"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
			pasienList[mid].Nama,
			pasienList[mid].JenisKelamin,
			pasienList[mid].Umur,
			pasienList[mid].PaketMCU.NamaPaket,
			float64(pasienList[mid].PaketMCU.Harga),
			pasienList[mid].WaktuMCU,
			pasienList[mid].HasilRekap,
		)
		return
	} else if pasienList[mid].Nama < nama {
		cariPasienBerdasarkanNamaRekursif(nama, mid+1, high)
	} else {
		cariPasienBerdasarkanNamaRekursif(nama, low, mid-1)
	}
}

func cariPasienBerdasarkanNama() {
	selectionSortNama()
	var nama string
	fmt.Print("Masukkan Nama Pasien: ")
	fmt.Scan(&nama)
	cariPasienBerdasarkanNamaRekursif(nama, 0, jumlahPasien-1)
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
		fmt.Printf(
			"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
			pasienList[i].Nama,
			pasienList[i].JenisKelamin,
			pasienList[i].Umur,
			pasienList[i].PaketMCU.NamaPaket,
			float64(pasienList[i].PaketMCU.Harga),
			pasienList[i].WaktuMCU,
			pasienList[i].HasilRekap,
		)
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
		fmt.Printf(
			"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
			pasienList[i].Nama,
			pasienList[i].JenisKelamin,
			pasienList[i].Umur,
			pasienList[i].PaketMCU.NamaPaket,
			float64(pasienList[i].PaketMCU.Harga),
			pasienList[i].WaktuMCU,
			pasienList[i].HasilRekap,
		)
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
		fmt.Printf(
			"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
			pasienList[i].Nama,
			pasienList[i].JenisKelamin,
			pasienList[i].Umur,
			pasienList[i].PaketMCU.NamaPaket,
			float64(pasienList[i].PaketMCU.Harga),
			pasienList[i].WaktuMCU,
			pasienList[i].HasilRekap,
		)
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
		fmt.Printf(
			"Nama: %s, Jenis Kelamin: %s, Umur: %d, Nama Paket: %s, Harga Paket: Rp. %.3f,  Tanggal MCU: %d, Hasil Rekap: %s\n",
			pasienList[i].Nama,
			pasienList[i].JenisKelamin,
			pasienList[i].Umur,
			pasienList[i].PaketMCU.NamaPaket,
			float64(pasienList[i].PaketMCU.Harga),
			pasienList[i].WaktuMCU,
			pasienList[i].HasilRekap,
		)
	}
}

func main() {
	initDummyData()
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
			penggunaRumahSakit()
		case 2:
			penggunaPasien()
		case 3:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
