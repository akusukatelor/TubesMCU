package dummy

import (
	"Tubes-MCU/data"
	"fmt"
)

func InitDummyData() {
	data.ListDokter[0].ID = 001
	data.ListDokter[0].Nama = "Dr.Alfin God"
	data.ListDokter[0].Spesialis = "Umum"
	data.ListDokter[0].Jadwal = "Senin-Jumat"
	data.ListDokter[1].ID = 002
	data.ListDokter[1].Nama = "Dr.Raqa Anshor"
	data.ListDokter[1].Spesialis = "THT"
	data.ListDokter[1].Jadwal = "Sabtu"
	data.ListDokter[2].ID = 003
	data.ListDokter[2].Nama = "Dr.Asbi hihi"
	data.ListDokter[2].Spesialis = "Mata"
	data.ListDokter[2].Jadwal = "Minggu"
	data.PaketMCUList[0] = data.PaketMCU{NamaPaket: "Basic", Harga: 500}
	data.PaketMCUList[1] = data.PaketMCU{NamaPaket: "Advanced", Harga: 100}
	data.PaketMCUList[2] = data.PaketMCU{NamaPaket: "Premium", Harga: 200}
	data.JumlahPaket = 3
	data.PasienList[0] = data.Pasien{Nama: "Udin", JenisKelamin: "P", Umur: 25, PaketMCU: data.PaketMCUList[0], WaktuMCU: 15, HasilRekap: "Sehat"}
	data.PasienList[1] = data.Pasien{Nama: "Agus", JenisKelamin: "L", Umur: 30, PaketMCU: data.PaketMCUList[1], WaktuMCU: 16, HasilRekap: "Sehat"}
	data.PasienList[2] = data.Pasien{Nama: "Santoso", JenisKelamin: "L", Umur: 40, PaketMCU: data.PaketMCUList[2], WaktuMCU: 17, HasilRekap: "Hipertensi"}
	data.JumlahPasien = 3
}
func LihatJadwalDokter() {

	for i := 0; i < 3; i++ {
		fmt.Println("ID: ", data.ListDokter[i].ID, ", Nama: ", data.ListDokter[i].Nama, ", Spesialis: ", data.ListDokter[i].Spesialis, ", Jadwal: ", data.ListDokter[i].Jadwal)
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
				fmt.Println(data.ListDokter[0].Nama, ": Untuk solusi gejala awal: \nIstirahat yang cukup: Tubuh membutuhkan energi untuk melawan infeksi.\nMinum banyak cairan: Air, teh hangat, atau sup kaldu membantu mencegah dehidrasi.\nKompres hangat: Tempelkan kain basah hangat di dahi untuk meredakan ketidaknyamanan.\nGunakan pakaian tipis: Hindari pakaian yang terlalu tebal agar panas tubuh dapat keluar.\nObat penurun demam: Paracetamol atau ibuprofen dapat digunakan sesuai dosis untuk menurunkan suhu tubuh.")
			} else {
				fmt.Println(data.ListDokter[0].Nama, ": Jika demam anda sudah lebih dari 3 hari maka disarankan untuk pergi ke rumah sakit langsung, akan tetapi untuk pertolongan awal anda bisa meminum obat seperti paracetamol atau ibuprofen dan disertai istirahat total.")
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
				fmt.Println(data.ListDokter[0].Nama, ": Jika sakit yang anda rasakan kurang dari 4 hari,mungkin yang anda bisa lakukan:\n1.Minum Air Putih,Dehidrasi sering menjadi penyebab utama pusing. Minumlah segelas air putih dan pastikan hidrasi kamu cukup sepanjang hari.\n2. Istirahat Sejenak, Matikan layar, pejamkan mata, dan istirahat selama 10–15 menit di tempat tenang. Hindari cahaya terang dan suara bising.\n3.  Makan atau Minum yang Manis, Pusing bisa disebabkan oleh gula darah rendah. Konsumsi makanan ringan seperti roti, buah, atau minuman hangat (teh manis/jahe).")
			} else {
				fmt.Println(data.ListDokter[0].Nama, ": Jika lebih dari 3 hari anda sudah melakukan penanganan awal seperti minum obat,istirahat,dll masih belum sembuh. Dan kika pusing disertai gejala seperti mual parah, muntah, lemas, pandangan kabur, disarankan untuk segera pergi ke rumah sakit langsung.")
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
				fmt.Println(data.ListDokter[2].Nama, ": Untuk gejala awal mungkin anda bisa lakukan: \n1.  Istirahatkan Mata,Hindari menatap layar komputer, ponsel, atau TV terlalu lama. Gunakan aturan 20-20-20: setiap 20 menit, alihkan pandangan ke sesuatu sejauh 20 kaki (6 meter) selama 20 detik.\n2. Kompres Hangat atau Dingin, kompres hangat untuk mata lelah atau bengkak, kompres dingin untuk mata merah,gatal atau meradang.\n3. Gunakan Obat Tetes Mata, Mata Kering gunakan artificial tears (air mata buatan), Mata Merah, Pilih tetes mata dengan decongestant (misal: Visine).")
			} else {
				fmt.Println(data.ListDokter[2].Nama, ": Jika mata kamu lebih dari 3 hari masih sakit setelah sudah diberi penanganan awal dan terasa tanda berikut: Nyeri parah atau bengkak berat, Penglihatan buram atau kabur, Keluar cairan abnormal seperti nanah. Disarankan langsung segera ke rumah sakit untuk penanganan lebih lanjut.")
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
				fmt.Println(data.ListDokter[0].Nama, ": Diare yang berlangsung kurang dari 4 hari umumnya dianggap sebagai diare akut. Berikut adalah beberapa solusi yang bisa dilakukan:\nMinum cairan lebih banyak: Untuk menggantikan cairan yang hilang, minum air putih, oralit (larutan rehidrasi oral), atau cairan elektrolit lain.\nHindari dehidrasi: Gejala dehidrasi meliputi mulut kering, rasa haus, lemas, atau warna urin yang lebih gelap.")
			} else {
				fmt.Println(data.ListDokter[0].Nama, ": Segera pergi ke dokter jika:\nDiare berlangsung lebih dari 4 hari.\nAda tanda dehidrasi berat (seperti mata cekung, tidak buang air kecil).\nAda darah dalam tinja atau tinja berwarna hitam.Disertai demam tinggi atau nyeri perut hebat.")
			}
		case 5:
			fmt.Print("Sudah berapa lama Anda mengalami gejala ini? (hari): ")
			var hari int
			fmt.Scan(&hari)

			if hari < 4 {
				fmt.Println(data.ListDokter[1].Nama, ": Untuk penangan awal:\nKompres hangat: Tempelkan kain hangat (tidak terlalu panas) pada telinga yang sakit selama 10-15 menit untuk mengurangi nyeri.\nKompres dingin: Jika nyeri disertai pembengkakan, gunakan kain yang dibungkus es batu sebagai kompres.")
			} else {
				fmt.Println(data.ListDokter[1].Nama, ": Segera pergi ke dokter jika:\nNyeri telinga berlangsung lebih dari 4 hari atau semakin parah.\nAda cairan keluar dari telinga (seperti nanah atau darah).Disertai demam tinggi, kehilangan pendengaran, atau pusing.Anak kecil sangat rewel, sering menarik-narik telinga, atau tidak mau makan.")
			}
		}
	}
}

func PenggunaPasien() {
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
			LihatJadwalDokter()
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
