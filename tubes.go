package main

import (
	"fmt"
	"os"
	"os/exec"
)

const nmax = 1000

type user struct {
	nama string
	skor int
}

type quiz struct {
	pertanyaan string
	opsi1      string
	opsi2      string
	opsi3      string
	opsi4      string
	jawaban    string
}

type bankSoal [nmax]quiz

type pemain [nmax]user

func clearline() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func identity() {
	fmt.Println("*======== TUBES ALGORITMA PEMOGRAMAN ========*")
	fmt.Println("=====  Nazhmi Ahmad Fauzan / 1301223056  =====")
	fmt.Println("=====      Putri Adinda / 1301223073     =====")
	fmt.Println("*============================================*")
}

func header() {
	fmt.Println("*============== SELAMAT DATANG ==============*")
	fmt.Println("=====                 Di                 =====")
	fmt.Println("=====   Who Wants to Be a Millionaire    =====")
	fmt.Println("*--------------------------------------------*")
}

func title() {
	var x string
	fmt.Print("press x to continue...   ")
	fmt.Scan(&x)
	for x != "x" && x != "X" {
		fmt.Print("Wrong input, press x to continue...   ")
		fmt.Scan(&x)
	}
}

func menuRegis(user, pass string) {
	// clearline()
	var choice int
	header()
	fmt.Println("===================")
	fmt.Println("PILIH MENU")
	fmt.Println("-------------------")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("0. Keluar")
	fmt.Println("-------------------")
	fmt.Print("Pilihan : ")
	fmt.Scan(&choice)
	fmt.Println(" ")

	if choice == 1 {
		register(user, pass)
		clearline()
		fmt.Println("Registrasi berhasil!")
		menuRegis(user, pass)
	}
	if choice == 2 {
		login(user, pass)
	}
	if choice == 0 {
		os.Exit(0)
	}
	if choice != 1 && choice != 2 && choice != 3 {
		clearline()
		fmt.Println("Pilihan tidak tersedia!")
		menuRegis(user, pass)
	}
}

func register(user, pass string) {
	clearline()

	var userRegis, userPass string
	header()
	fmt.Println("====================")
	fmt.Println("REGISTRASI")
	fmt.Println("--------------------")
	fmt.Println("Username : ")
	fmt.Scan(&userRegis)

	fmt.Println("Password : ")
	fmt.Scan(&userPass)
	fmt.Println(" ")
	user = userRegis
	pass = userPass
	menuRegis(user, pass)
}

func login(user, pass string) {
	clearline()
	var t bankSoal
	var a pemain
	var nSoal int
	nPemain := 0

	var userLogin, passLogin string
	header()
	fmt.Println("====================")
	fmt.Println("LOGIN")
	fmt.Println("--------------------")
	fmt.Println("Username : ")
	fmt.Scan(&userLogin)
	fmt.Println("Password : ")
	fmt.Scan(&passLogin)
	fmt.Println(" ")

	if userLogin == user && passLogin == pass {
		clearline()
		fmt.Println("Login berhasil!")
		fmt.Println(" ")
	} else {
		clearline()
		fmt.Println("Username atau password salah!")
		fmt.Println(" ")
		menuRegis(user, pass)
	}
	role(t, a, nSoal, nPemain)
}

func role(t bankSoal, a pemain, nSoal, nPemain int) {
	clearline()
	var kode, isAdmin, x int
	header()
	fmt.Println("====================")
	fmt.Println("PILIH PERANGKAT")
	fmt.Println("--------------------")
	fmt.Println("1. Admin")
	fmt.Println("2. Pemain")
	fmt.Println("0. Keluar")
	fmt.Println("--------------------")
	fmt.Print("Pilihan : ")
	fmt.Scan(&x)
	fmt.Println(" ")
	if x == 1 {
		isAdmin = 1
		menu(t, nSoal, nPemain, a, kode, isAdmin)
		role(t, a, nSoal, nPemain)
	}
	if x == 2 {
		menu(t, nSoal, nPemain, a, kode, isAdmin)
	}
	if x == 0 {
		os.Exit(0)
	}
}
func menu(t bankSoal, nSoal, nPemain int, a pemain, kode int, isAdmin int) {
	clearline()
	var x int
	header()
	if isAdmin == 1 {
		fmt.Println("====================")
		fmt.Println("MENU")
		fmt.Println("--------------------")
		fmt.Println("1. Tambah soal")
		fmt.Println("2. Ubah Soal")
		fmt.Println("3. Hapus soal")
		fmt.Println("4. Tampilkan soal")
		fmt.Println("0. Kembali")
		fmt.Println("--------------------")
		fmt.Print("Pilihan : ")
		fmt.Scan(&x)
		if x != 1 && x != 2 && x != 3 && x != 4 && x != 0 {
			clearline()
			fmt.Println("Masukan salah!")
			fmt.Println(" ")
			menu(t, nSoal, nPemain, a, kode, isAdmin)
		} else {
			if x == 1 {
				fmt.Println(" ")
				tambahSoal(&t, &nSoal)
				menu(t, nSoal, nPemain, a, kode, 1)
			}
			if x == 2 {
				clearline()
				ubahSoal(t, nSoal, kode)
				menu(t, nSoal, nPemain, a, kode, 1)

			}
			if x == 3 {
				clearline()
				hapusSoal(t, nSoal, kode)
			}
			if x == 4 {
				clearline()
				printSoal(t, nSoal)
				menu(t, nSoal, nPemain, a, kode, 1)
			}
			if x == 0 {
				role(t, a, nSoal, nPemain)
			}
		}
	} else {
		clearline()
		header()
		var x int
		fmt.Println(nPemain)
		fmt.Println("====================")
		fmt.Println("MENU PEMAIN")
		fmt.Println("--------------------")
		fmt.Println("1. Mulai")
		fmt.Println("2. Leaderboard")
		fmt.Println("0. Kembali")
		fmt.Println("--------------------")
		fmt.Print("Pilihan : ")
		fmt.Scan(&x)
		if x == 1 {
			inputPemain(t, a, nSoal, nPemain)
		}
		if x == 2 {
			sortPemain(t, a, nPemain, nSoal)
		}
		if x == 0 {
			role(t, a, nSoal, nPemain)
		}
	}
}

func tambahSoal(t *bankSoal, nSoal *int) {
	clearline()
	// space := bufio.NewReader(os.Stdin)
	*nSoal = 0
	fmt.Println("Masukan 5 soal! ")
	for *nSoal < 5 {
		fmt.Println("Soal ke ", (*nSoal + 1))
		fmt.Println("Pertanyaan: (Gunakan underscore untuk spasi)")
		fmt.Scan(&t[*nSoal].pertanyaan)
		fmt.Print("Opsi 1 :")
		fmt.Scan(&t[*nSoal].opsi1)
		fmt.Print("Opsi 2 :")
		fmt.Scan(&t[*nSoal].opsi2)
		fmt.Print("Opsi 3 :")
		fmt.Scan(&t[*nSoal].opsi3)
		fmt.Print("Opsi 4 :")
		fmt.Scan(&t[*nSoal].opsi4)
		fmt.Print("Kunci jawaban(A/B/C/D): ")
		fmt.Scan(&t[*nSoal].jawaban)
		fmt.Println("-------------------------------")
		fmt.Println(" ")
		*nSoal++
	}
}
func cariSoal(t bankSoal, kode int) int {
	i := 0
	for i <= 5 {
		if i == kode-1 {
			return i
		}
		i++
	}
	return -1
}

func printSoal(t bankSoal, nSoal int) {
	var x string
	i := 0
	for i < nSoal {
		fmt.Println("Soal ke ", i+1, ":")
		fmt.Println(t[i].pertanyaan)
		fmt.Println("Opsi :")
		fmt.Println("A. ", t[i].opsi1)
		fmt.Println("B. ", t[i].opsi2)
		fmt.Println("C. ", t[i].opsi3)
		fmt.Println("D. ", t[i].opsi4)
		fmt.Println(t[i].jawaban)
		fmt.Println("=======================")
		i++
	}
	fmt.Print("Tekan x untuk lanjut: ")
	for x != "x" && x != "X" {
		fmt.Print("Input Salah, tekan x untuk lanjut...   ")
		fmt.Scan(&x)
	}
}

func ubahSoal(t bankSoal, nSoal int, kode int) {
	var a pemain
	var nPemain int
	fmt.Print("Pilih nomor soal untuk di ubah : ")
	fmt.Scan(&kode)
	found := cariSoal(t, kode)
	if found == -1 {
		clearline()
		fmt.Println("Soal tidak ditemukan!")
	} else {
		fmt.Println("Soal :", t[found].pertanyaan)
		fmt.Println("---------------------------------------")
		fmt.Println("A. ", t[found].opsi1, "    ", "B. ", t[found].opsi2)
		fmt.Println("C. ", t[found].opsi3, "    ", "D. ", t[found].opsi4)
		fmt.Println("Ubah soal: Gunakan underscore untuk spasi")
		fmt.Scan(&t[found].pertanyaan)
		fmt.Print("Opsi 1 :")
		fmt.Scan(&t[found].opsi1)
		fmt.Print("Opsi 2 :")
		fmt.Scan(&t[found].opsi2)
		fmt.Print("Opsi 3 :")
		fmt.Scan(&t[found].opsi3)
		fmt.Print("Opsi 4 :")
		fmt.Scan(&t[found].opsi4)
		fmt.Print("Kunci jawaban(A/B/C/D): ")
		fmt.Scan(&t[found].jawaban)
	}
	menu(t, nSoal, nPemain, a, kode, 1)
}
func hapusSoal(t bankSoal, nSoal int, kode int) {
	var x string
	var a pemain
	var nPemain int
	fmt.Print("Pilih nomor soal untuk di hapus : ")
	fmt.Scan(&kode)
	found := cariSoal(t, kode)
	if found == -1 {
		clearline()
		fmt.Println("Soal tidak ditemukan!")
	} else {
		fmt.Println("Soal :", t[found].pertanyaan)
		fmt.Println("---------------------------------------")
		fmt.Println("A. ", t[found].opsi1, "    ", "B. ", t[found].opsi2)
		fmt.Println("C. ", t[found].opsi3, "    ", "D. ", t[found].opsi4)
		fmt.Println(" ")
		fmt.Println("Anda yakin ingin menghapus soal ini? (y/n)")
		fmt.Scan(&x)
		if x == "y" || x == "Y" { // menggeser array ke kiri
			for i := found; i < nSoal; i++ {
				t[i] = t[i+1]
			}
			nSoal--
			fmt.Println("Data berhasil dihapus")
			fmt.Println(" ")
			fmt.Print("Tekan x untuk lanjut: ")
			fmt.Scan(&x)
			for x != "x" && x != "X" {
				fmt.Print("Input Salah, tekan x untuk lanjut...   ")
				fmt.Scan(&x)
			}
			menu(t, nSoal, nPemain, a, kode, 1)
		} else if x == "n" && x == "N" {
			menu(t, nSoal, nPemain, a, kode, 1)
		}
	}
}

func inputPemain(t bankSoal, a pemain, nSoal, nPemain int) {
	for nPemain < nmax {
		fmt.Println("Masukan nama pemain :")
		fmt.Scan(&a[nPemain].nama)
		mengerjakanSoal(t, a, nSoal, nPemain)
	}
}

func mengerjakanSoal(t bankSoal, a pemain, nSoal, nPemain int) {
	var jawab string
	var k int
	var startQuiz string
	var isAdmin int
	i := 0
	a[nPemain].skor = 0
	for i < nSoal {
		fmt.Println("Soal", i+1, t[i].pertanyaan)
		fmt.Println("---------------------------------------")
		fmt.Println("A. ", t[i].opsi1, "    ", "B. ", t[i].opsi2)
		fmt.Println("C. ", t[i].opsi3, "    ", "D. ", t[i].opsi4)
		fmt.Print("Jawaban : ")
		fmt.Scan(&jawab)
		fmt.Println(" ")
		if jawab == t[i].jawaban {
			a[nPemain].skor += 20
			fmt.Println("Jawaban benar")
		} else {
			fmt.Println("Jawaban salah!")
			fmt.Print(" ")
			fmt.Print("Jawaban yang benar : ")
			fmt.Println(t[i].jawaban)
		}
		i++
	}
	clearline()
	fmt.Println("nSoal = ", nSoal)
	fmt.Println("Nilai anda : ", a[nPemain].skor)
	nPemain++

	fmt.Println("Apakah ingin mengulang? (y/n)")
	fmt.Scan(&startQuiz)
	for startQuiz != "y" && startQuiz != "Y" && startQuiz != "n" && startQuiz != "N" {
		fmt.Println("Error")
		fmt.Println("Apakah ingin mengulang? (y/n)")
		fmt.Scan(&startQuiz)
	}
	if startQuiz == "Y" || startQuiz == "y" {
		mengerjakanSoal(t, a, nSoal, nPemain)
	} else if startQuiz == "N" || startQuiz == "n" {
		menu(t, nSoal, nPemain, a, k, isAdmin)
	}
}

func sortPemain(t bankSoal, a pemain, n, nSoal int) { // mengurutkan pemain berdasarkan skor terbesar menggunakan selection sort

	var temp user
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i].skor < a[j].skor {
				temp = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	printScoreBoard(t, a, n, nSoal)
}

func printScoreBoard(t bankSoal, a pemain, nPemain, nSoal int) {
	clearline()
	header()
	var kembali string
	var k, isAdmin int
	i := 0
	fmt.Println("*========LEADERBOARD========*")
	fmt.Println("-----------------------------")
	for i < nPemain {
		fmt.Println(i+1, ". Nama : ", a[i].nama, ", Skor :", a[i].skor)
		i++
	}
	fmt.Print("Tekan x untuk kembali ke menu... ")
	fmt.Scan(&kembali)
	if kembali != "x" && kembali != "X" {
		for kembali != "x" && kembali != "X" {
			fmt.Print("Input salah, tekan x untuk lanjut...   ")
			fmt.Scan(&kembali)
		}
	} else {
		menu(t, nSoal, nPemain, a, k, isAdmin)
	}
}

func main() {
	var user, pass string
	// var t bankSoal
	// var a pemain
	// var n int
	identity()
	title()
	menuRegis(user, pass)
	// role(t, a, n)
}
