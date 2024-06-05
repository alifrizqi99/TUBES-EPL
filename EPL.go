package main

import "fmt"

type Club struct {
	Name                                                                string
	Matches, Wins, Loses, Draws, Goals, GoalsAgainst, GoalsDiff, Points int
}

var (
	CLUBS_DATA []Club
	total_data = 0
)

func main() {
	var option string

	for {
		fmt.Println("\n============== MASUKAN CLUB PESERTA ==============")
		fmt.Println("Pilihan: ")
		fmt.Println("1. Tambah Club")
		fmt.Println("2. Tambah Hasil Pertandingan")
		fmt.Println("3. Hapus Club")
		fmt.Println("==================================================")
		printClubs()

		fmt.Println("\nPilihan: ")
		fmt.Scanf("%s\n", &option)

		if option == "1" {
			addClub()
		} else if option == "2" {
			addMatch()
		} else if option == "3" {
			deleteClub()
		} else {
			fmt.Println("pilihan tidak valid, tolong coba lagi.")
		}
	}
}

func addClub() {
	var name string
	fmt.Println("Masukan nama klub: ")
	fmt.Scanf("%s\n", &name)

	newClub := Club{
		Name: name,
	}

	CLUBS_DATA = append(CLUBS_DATA, newClub)
	total_data += 1
	fmt.Println("Klub berhasil ditambahkan!")
}

func sortHighestPoint() {
	for i := 0; i < total_data-1; i++ {
		minIndex := i
		for j := i + 1; j < total_data; j++ {
			if CLUBS_DATA[j].Points > CLUBS_DATA[minIndex].Points {
				minIndex = j
			}
		}
		temp := CLUBS_DATA[minIndex]
		CLUBS_DATA[minIndex] = CLUBS_DATA[i]
		CLUBS_DATA[i] = temp
	}
}

func printClubs() {
	if total_data == 0 {
		fmt.Println("Tidak ada data")
		return
	}

	fmt.Printf("\n%-20s %-20s %-8s %-8s %-8s %-8s %-15s %-15s %-8s\n", "Nama Club", "Total Pertandingan", "Menang", "Kalah", "Seri", "Gol", "Gol Kemasukan", "Selisih Goal", "Point")
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------------")

	sortHighestPoint()

	for i := 0; i < total_data; i++ {
		fmt.Printf("%-20s %-20d %-8d %-8d %-8d %-8d %-15d %-15d %-8d\n", CLUBS_DATA[i].Name, CLUBS_DATA[i].Matches, CLUBS_DATA[i].Wins, CLUBS_DATA[i].Loses, CLUBS_DATA[i].Draws, CLUBS_DATA[i].Goals, CLUBS_DATA[i].GoalsAgainst, CLUBS_DATA[i].GoalsDiff, CLUBS_DATA[i].Points)
	}
}

func addMatch() {
	var clubName string

	fmt.Println("Masukan Nama klub: ")
	fmt.Scanf("%s\n", &clubName)

	index := getClubIndex(clubName)

	if index == -1 {
		fmt.Println("Nama klub tidak valid")
		addMatch()
		return
	} else {
		var hasil string
		fmt.Println("Masukan hasil pertandingan (menang, kalah, seri): ")
		fmt.Scanf("%s\n", &hasil)

		if hasil == "menang" {
			CLUBS_DATA[index].Wins += 1
			CLUBS_DATA[index].Points += 3
		} else if hasil == "kalah" {
			CLUBS_DATA[index].Loses += 1
			CLUBS_DATA[index].Points += 0
		} else if hasil == "seri" {
			CLUBS_DATA[index].Draws += 1
			CLUBS_DATA[index].Points += 1
		} else {
			fmt.Println("Hasil tidak valid")
			addMatch()
			return
		}

		CLUBS_DATA[index].Matches++
		addGoal(index)
	}
}

func addGoal(index int) {
	for {
		CLUBS_DATA[index].GoalsDiff = CLUBS_DATA[index].Goals - CLUBS_DATA[index].GoalsAgainst

		fmt.Println("Pilih salah satu:")
		fmt.Println("1. Tambah Jumlah Gol")
		fmt.Println("2. Tambah Kemasukan Gol")
		fmt.Println("3. Selesai")

		var option string
		fmt.Println("Pilihan: ")
		fmt.Scanf("%s\n", &option)

		if option == "1" {
			var goal int

			fmt.Println("Jumlah gol: ")
			fmt.Scanf("%d\n", &goal)

			CLUBS_DATA[index].Goals += goal

			fmt.Println("Gol berhasil ditambahkan!\n")
			continue
		} else if option == "2" {
			var goal int

			fmt.Println("Jumlah gol kemasukan: ")
			fmt.Scanf("%d\n", &goal)

			CLUBS_DATA[index].GoalsAgainst += goal

			fmt.Println("Gol kemasukan berhasil ditambahkan!")
			continue
		} else if option == "3" {
			return
		} else {
			fmt.Println("Pilihan tidak valid, tolong coba lagi.")
		}
	}
}

func deleteClub() {
	var clubName string

	fmt.Println("Masukan Nama klub: ")
	fmt.Scanf("%s\n", &clubName)

	index := getClubIndex(clubName)

	if index == -1 {
		fmt.Println("Nama klub tidak valid")
		deleteClub()
		return
	} else {
		var choose string
		fmt.Println("Apakah anda yakin ingin menghapusnya? (y/n)")
		fmt.Scanf("%s\n", &choose)
		if choose == "y" {
			CLUBS_DATA = append(CLUBS_DATA[:index], CLUBS_DATA[index+1:]...)
			total_data -= 1
			fmt.Println("Klub berhasil dihapus!")
		} else {
			fmt.Println("Penghapusan digagalkan pengguna")
		}
	}
}

func getClubIndex(name string) int {
	for i := 0; i < total_data; i++ {
		if CLUBS_DATA[i].Name == name {
			return i
		}
	}
	return -1
}
