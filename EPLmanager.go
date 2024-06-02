package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Club merepresentasikan sebuah klub sepakbola
type Club struct {
	Nama         string
	Pertandingan int
	Menang       int
	Seri         int
	Kalah        int
	GolMasuk     int
	GolKemasukan int
	SelisihGol   int
	Point        int
}

// NewClub membuat klub sepakbola baru
func NewClub(nama string) *Club {
	return &Club{
		Nama: nama,
	}
}

// UpdateStats memperbarui statistik klub berdasarkan hasil pertandingan
func (c *Club) UpdateStats(result string, golMasuk, golKemasukan int) {
	c.Pertandingan++
	switch result {
	case "menang":
		c.Menang++
		c.Point += 3
	case "seri":
		c.Seri++
		c.Point++
	case "kalah":
		c.Kalah++
	}
	c.GolMasuk += golMasuk
	c.GolKemasukan += golKemasukan
	c.SelisihGol = c.GolMasuk - c.GolKemasukan
}

// Fixture merepresentasikan jadwal pertandingan sepakbola
type Fixture struct {
	TuanRumah string
	Tamu      string
	Hasil     string
}

// NewFixture membuat jadwal pertandingan baru
func NewFixture(tuanRumah, tamu string) *Fixture {
	return &Fixture{
		TuanRumah: tuanRumah,
		Tamu:      tamu,
	}
}

// UpdateResult memperbarui hasil pertandingan
func (f *Fixture) UpdateResult(result string) {
	f.Hasil = result
}

// EPLManager mengelola Liga Utama Inggris
type EPLManager struct {
	Klub   []*Club
	Jadwal []*Fixture
}

// NewEPLManager membuat manajer Liga Utama Inggris baru
func NewEPLManager() *EPLManager {
	return &EPLManager{
		Klub:   make([]*Club, 0),
		Jadwal: make([]*Fixture, 0),
	}
}

// AddClub menambahkan klub sepakbola ke dalam manajer Liga Utama Inggris
func (epl *EPLManager) AddClub(nama string) {
	epl.Klub = append(epl.Klub, NewClub(nama))
}

// AddFixture menambahkan jadwal pertandingan ke dalam manajer Liga Utama Inggris
func (epl *EPLManager) AddFixture(tuanRumah, tamu string) {
	epl.Jadwal = append(epl.Jadwal, NewFixture(tuanRumah, tamu))
}

// UpdateFixtureResult memperbarui hasil pertandingan dan statistik klub
func (epl *EPLManager) UpdateFixtureResult(tuanRumah, tamu, hasil string, golMasuk, golKemasukan int) {
	for _, f := range epl.Jadwal {
		if f.TuanRumah == tuanRumah && f.Tamu == tamu {
			f.UpdateResult(hasil)
			klubTuanRumah := epl.getClubByName(tuanRumah)
			klubTamu := epl.getClubByName(tamu)
			klubTuanRumah.UpdateStats(hasil, golMasuk, golKemasukan)
			klubTamu.UpdateStats(hasil, golKemasukan, golMasuk)
			break
		}
	}
}

// getClubByName mengembalikan klub dengan nama yang diberikan
func (epl *EPLManager) getClubByName(nama string) *Club {
	for _, c := range epl.Klub {
		if c.Nama == nama {
			return c
		}
	}
	return nil
}

// SelectionSortClubs mengurutkan klub berdasarkan kriteria yang ditentukan dengan menggunakan selection sort
func SelectionSortClubs(klub []*Club, berdasarkan string, urutan string) []*Club {
	klubTerurut := make([]*Club, len(klub))
	copy(klubTerurut, klub)

	switch berdasarkan {
	case "point":
		for i := 0; i < len(klubTerurut)-1; i++ {
			indeksMinimum := i
			for j := i + 1; j < len(klubTerurut); j++ {
				if urutan == "asc" {
					if klubTerurut[j].Point < klubTerurut[indeksMinimum].Point {
						indeksMinimum = j
					}
				} else {
					if klubTerurut[j].Point > klubTerurut[indeksMinimum].Point {
						indeksMinimum = j
					}
				}
			}
			if indeksMinimum != i {
				klubTerurut[i], klubTerurut[indeksMinimum] = klubTerurut[indeksMinimum], klubTerurut[i]
			}
		}
		// Implementasikan pengurutan berdasarkan kriteria lain jika diperlukan
	}

	return klubTerurut
}

// InsertionSortClubs mengurutkan klub berdasarkan kriteria yang ditentukan dengan menggunakan insertion sort
func InsertionSortClubs(klub []*Club, berdasarkan string, urutan string) []*Club {
	klubTerurut := make([]*Club, len(klub))
	copy(klubTerurut, klub)

	switch berdasarkan {
	case "pertandingan":
		for i := 1; i < len(klubTerurut); i++ {
			kunci := klubTerurut[i]
			j := i - 1
			if urutan == "asc" {
				for j >= 0 && klubTerurut[j].Pertandingan > kunci.Pertandingan {
					klubTerurut[j+1] = klubTerurut[j]
					j--
				}
			} else {
				for j >= 0 && klubTerurut[j].Pertandingan < kunci.Pertandingan {
					klubTerurut[j+1] = klubTerurut[j]
					j--
				}
			}
			klubTerurut[j+1] = kunci
		}
		// Implementasikan pengurutan berdasarkan kriteria lain jika diperlukan
	}

	return klubTerurut
}

func main() {
	epl := NewEPLManager()
	reader := bufio.NewReader(os.Stdin)

	// Menambahkan klub-klub
	epl.AddClub("MUN")
	epl.AddClub("LIV")
	epl.AddClub("CHE")
	epl.AddClub("ARS")
	epl.AddClub("MCY")

	// Loop menu utama
	for {
		fmt.Println("\n=== Manajer Liga Utama Inggris ===")
		fmt.Println("1. Masukkan Hasil Pertandingan")
		fmt.Println("2. Lihat Klasemen")
		fmt.Println("3. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")

		pilihanStr, _ := reader.ReadString('\n')
		pilihanStr = strings.TrimSpace(pilihanStr)
		pilihan, err := strconv.Atoi(pilihanStr)
		if err != nil {
			fmt.Println("Pilihan tidak valid. Silakan masukkan angka.")
			continue
		}

		switch pilihan {
		case 1:
			fmt.Println("\n=== Masukkan Hasil Pertandingan ===")
			fmt.Print("Masukkan Tim Tuan Rumah: ")
			tuanRumah, _ := reader.ReadString('\n')
			tuanRumah = strings.TrimSpace(tuanRumah)
			fmt.Print("Masukkan Tim Tamu: ")
			tamu, _ := reader.ReadString('\n')
			tamu = strings.TrimSpace(tamu)
			fmt.Print("Masukkan Hasil (menang, seri, kalah): ")
			hasil, _ := reader.ReadString('\n')
			hasil = strings.TrimSpace(hasil)
			fmt.Print("Masukkan Gol Untuk Tim Tuan Rumah: ")
			golMasukStr, _ := reader.ReadString('\n')
			golMasukStr = strings.TrimSpace(golMasukStr)
			golMasuk, err := strconv.Atoi(golMasukStr)
			if err != nil {
				fmt.Println("Input untuk jumlah gol tidak valid.")
				continue
			}
			fmt.Print("Masukkan Gol Untuk Tim Tamu: ")
			golKemasukanStr, _ := reader.ReadString('\n')
			golKemasukanStr = strings.TrimSpace(golKemasukanStr)
			golKemasukan, err := strconv.Atoi(golKemasukanStr)
			if err != nil {
				fmt.Println("Input untuk jumlah gol tidak valid.")
				continue
			}

			epl.UpdateFixtureResult(tuanRumah, tamu, hasil, golMasuk, golKemasukan)
			fmt.Println("Hasil pertandingan berhasil diperbarui.")

		case 2:
			fmt.Println("\n=== Klasemen Liga Utama Inggris ===")
			fmt.Println("1. Urutkan berdasarkan Point (Ascending)")
			fmt.Println("2. Urutkan berdasarkan Point (Descending)")
			fmt.Println("3. Urutkan berdasarkan Pertandingan Dimainkan (Ascending)")
			fmt.Println("4. Urutkan berdasarkan Pertandingan Dimainkan (Descending)")
			fmt.Print("Masukkan pilihan Anda: ")

			pilihanSortirStr, _ := reader.ReadString('\n')
			pilihanSortirStr = strings.TrimSpace(pilihanSortirStr)
			pilihanSortir, err := strconv.Atoi(pilihanSortirStr)
			if err != nil {
				fmt.Println("Pilihan tidak valid. Silakan masukkan angka.")
				continue
			}

			var berdasarkan, urutan string
			switch pilihanSortir {
			case 1:
				berdasarkan, urutan = "point", "asc"
			case 2:
				berdasarkan, urutan = "point", "desc"
			case 3:
				berdasarkan, urutan = "pertandingan", "asc"
			case 4:
				berdasarkan, urutan = "pertandingan", "desc"
			default:
				fmt.Println("Pilihan tidak valid.")
				continue
			}

			klubTerurut := SelectionSortClubs(epl.Klub, berdasarkan, urutan)
			fmt.Println("\nKlasemen Liga Utama Inggris:")
			for i, klub := range klubTerurut {
				fmt.Printf("%d. %s - Point: %d\n", i+1, klub.Nama, klub.Point)
			}

		case 3:
			fmt.Println("Keluar dari program...")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan masukkan pilihan yang benar.")
		}
	}
}
