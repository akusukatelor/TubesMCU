package data

const MAXDATA = 100

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
	PasienList   [MAXDATA]Pasien
	PaketMCUList [MAXDATA]PaketMCU
	ListDokter   [MAXDATA]Dokter
	JumlahPasien int
	JumlahPaket  int
)
