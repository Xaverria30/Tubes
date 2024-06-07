package main

import (
	"bufio"
	"fmt"
	"os"
)

type Sparepart struct {
	ID      int
	Nama    string
	Harga   int
	Terjual int
}

type Pelanggan struct {
	ID   int
	Nama string
}

type Transaksi struct {
	IDPelanggan int
	IDSparepart int
	Jumlah      int
	Harga       int
	Pajak       int
	TotalHarga  int
}

type spareparts [100]Sparepart
type pelanggan [100]Pelanggan
type transaksi [100]Transaksi

func main() {
	var S spareparts
	var P pelanggan
	var T transaksi
	var sData, pData, tData int
	S[0] = Sparepart{ID: 1, Nama: "Ban Motor", Harga: 150000, Terjual: 5}
	S[1] = Sparepart{ID: 2, Nama: "Kampas Rem", Harga: 75000, Terjual: 7}
	S[2] = Sparepart{ID: 3, Nama: "Busi", Harga: 30000, Terjual: 10}
	S[3] = Sparepart{ID: 4, Nama: "Oli Mesin", Harga: 50000, Terjual: 0}
	sData = 4

	P[0] = Pelanggan{ID: 1, Nama: "Andi"}
	P[1] = Pelanggan{ID: 2, Nama: "Budi"}
	P[2] = Pelanggan{ID: 3, Nama: "Caca"}
	pData = 3

	T[0] = Transaksi{IDPelanggan: 1, IDSparepart: 1, Jumlah: 5, Harga: 150000, Pajak: 37500, TotalHarga: 787500}
	T[1] = Transaksi{IDPelanggan: 2, IDSparepart: 2, Jumlah: 7, Harga: 75000, Pajak: 26250, TotalHarga: 551250}
	T[2] = Transaksi{IDPelanggan: 3, IDSparepart: 3, Jumlah: 10, Harga: 30000, Pajak: 15000, TotalHarga: 315000}
	tData = 3
	selectionSort(S[:sData])

	menu(&S, &P, &T, &sData, &pData, &tData)
}

func menu(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	fmt.Println("================================================")
	fmt.Println("   SELAMAT DATANG DI APLIKASI SPAREPART MOTOR")
	fmt.Println("================================================")
	fmt.Println("                 Created by: ")
	fmt.Println("         Rara Nur Annisa (103032300077)")
	fmt.Println("       Neng Intan Nuraeini (103032330031)")
	fmt.Println("------------------------------------------------")
	fmt.Println()
	fmt.Println("                1. Login Admin")
	fmt.Println("                2. Login Pembeli")
	fmt.Println("                3. Keluar")
	fmt.Println()
	fmt.Println("================================================")
	fmt.Print("Pilih: ")
	main_menu(S, P, T, sData, pData, tData)
}

func main_menu(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	var pilih int
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		login_admin(S, P, T, sData, pData, tData)
	case 2:
		login_pembeli(S, P, T, sData, pData, tData)
	case 3:
		return
	default:
		menu(S, P, T, sData, pData, tData)
	}
}

func login_admin(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	fmt.Println("==============================================")
	fmt.Println("        SELAMAT DATANG DI SISTEM ADMIN      ")
	fmt.Println("==============================================")
	fmt.Println("                 Pilihan Menu                  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Lihat Transaksi")
	fmt.Println("5. Lihat daftar pelanggan ")
	fmt.Println("   yang membeli sparepart tertentu")
	fmt.Println("6. Kembali")
	fmt.Println("==============================================")
	fmt.Print("Pilih: ")
	var pilih int
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		tambah_data_admin(S, P, T, sData, pData, tData)
	case 2:
		edit_data_admin(S, P, T, sData, pData, tData)
	case 3:
		hapus_data_admin(S, P, T, sData, pData, tData)
	case 4:
		viewTransactions(S, P, T, sData, pData, tData)
	case 5:
		daftar_pelanggan_sparepart(S, P, T, sData, pData, tData)
	case 6:
		menu(S, P, T, sData, pData, tData)
	default:
		login_admin(S, P, T, sData, pData, tData)
	}
}

func tambah_data_admin(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	fmt.Println("==============================================")
	fmt.Println("                   TAMBAH DATA                ")
	fmt.Println("==============================================")
	fmt.Println("  1. Sparepart               ")
	fmt.Println("  2. Pelanggan               ")
	fmt.Println("  3. Transaksi               ")
	fmt.Println("  4. Kembali                 ")
	fmt.Println("==============================================")
	fmt.Print("Pilih: ")
	var pilih int
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		tambah_sparepart(S, sData, P, T, pData, tData)
	case 2:
		tambah_pelanggan(S, P, T, sData, pData, tData)
	case 3:
		tambah_transaksi(S, P, T, sData, pData, tData)
	case 4:
		login_admin(S, P, T, sData, pData, tData)
	default:
		tambah_data_admin(S, P, T, sData, pData, tData)
	}
}

func tambah_sparepart(S *spareparts, sData *int, P *pelanggan, T *transaksi, pData, tData *int) {
	var sparepart Sparepart
	fmt.Print("Masukan ID Sparepart: ")
	fmt.Scanln(&sparepart.ID)
	fmt.Print("Masukan Nama Sparepart: ")
	sparepart.Nama = scan_word()
	fmt.Print("Masukan Harga Sparepart: ")
	fmt.Scanln(&sparepart.Harga)
	S[*sData] = sparepart
	(*sData)++
	fmt.Println("Data Sparepart berhasil ditambahkan!")
	login_admin(S, P, T, sData, pData, tData)
}

func tambah_pelanggan(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	var pelanggan Pelanggan
	fmt.Print("Masukan ID Pelanggan: ")
	fmt.Scanln(&pelanggan.ID)
	fmt.Print("Masukan Nama Pelanggan: ")
	pelanggan.Nama = scan_word()
	P[*pData] = pelanggan
	(*pData)++
	fmt.Println("Data Pelanggan berhasil ditambahkan!")
	login_admin(S, P, T, sData, pData, tData)
}

func tambah_transaksi(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	var transaksi Transaksi
	fmt.Print("Masukan ID Pelanggan:")
	fmt.Scanln(&transaksi.IDPelanggan)
	if !checkPelangganExist(P, transaksi.IDPelanggan, *pData) {
		fmt.Println("ID Pelanggan tidak ditemukan.")
		login_admin(S, P, T, sData, pData, tData)
		return
	}
	fmt.Print("Masukan ID Sparepart:")
	fmt.Scanln(&transaksi.IDSparepart)
	if !checkSparepartExist(S, transaksi.IDSparepart, *sData) {
		fmt.Println("ID Sparepart tidak ditemukan.")
		login_admin(S, P, T, sData, pData, tData)
		return
	}
	fmt.Print("Masukan Jumlah:")
	fmt.Scanln(&transaksi.Jumlah)
	transaksi.Harga = transaksi.Jumlah * getHargaSparepart(S, transaksi.IDSparepart, *sData)
	transaksi.Pajak = transaksi.Harga * 5 / 100
	transaksi.TotalHarga = transaksi.Harga + transaksi.Pajak
	T[*tData] = transaksi
	(*tData)++

	for i := 0; i < *sData; i++ {
		if S[i].ID == transaksi.IDSparepart {
			S[i].Terjual += transaksi.Jumlah
			i = *sData
		}
	}

	fmt.Println("Data Transaksi berhasil ditambahkan!")
	login_admin(S, P, T, sData, pData, tData)
}

func viewTransactions(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	if *tData == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}

	for i := 0; i < *tData; i++ {
		transaction := T[i]
		var customerName, sparePartName string
		foundCustomer := false
		foundSparePart := false

		for j := 0; j < *pData && !foundCustomer; j++ {
			if P[j].ID == transaction.IDPelanggan {
				customerName = P[j].Nama
				foundCustomer = true
			}
		}

		for k := 0; k < *sData && !foundSparePart; k++ {
			if S[k].ID == transaction.IDSparepart {
				sparePartName = S[k].Nama
				foundSparePart = true
			}
		}
		HargaAsli := getHargaSparepart(S, transaction.IDSparepart, *sData)
		transaction.Harga = transaction.Jumlah * getHargaSparepart(S, transaction.IDSparepart, *sData)
		transaction.Pajak = transaction.Harga * 5 / 100
		transaction.TotalHarga = transaction.Harga + transaction.Pajak

		fmt.Printf("Transaksi ke-%d\n", i+1)
		if foundCustomer {
			fmt.Printf("  Nama: %s\n", customerName)
		} else {
			fmt.Printf("  Nama: Tidak Ditemukan\n")
		}
		fmt.Printf("  ID: %d\n", transaction.IDPelanggan)
		fmt.Println("  --------------------------------------")
		fmt.Printf("  Nama Spare-part       Harga       Jumlah\n")
		if foundSparePart {
			fmt.Printf("  %-18s Rp.%-11d %-5d\n", sparePartName, HargaAsli, transaction.Jumlah)
		} else {
			fmt.Printf("  Nama Spare-part Tidak Ditemukan Rp.%-10d %-5d\n", transaction.Harga, transaction.Jumlah)
		}
		fmt.Println("  --------------------------------------")
		fmt.Printf("  Pajak: Rp.%d\n", transaction.Pajak)
		fmt.Printf("  Total Harga: Rp.%d\n", transaction.TotalHarga)
		fmt.Println()
	}

	login_admin(S, P, T, sData, pData, tData)
}

func daftar_pelanggan_sparepart(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	var id int
	fmt.Print("Masukan ID Sparepart yang ingin dicari pelanggannya: ")
	fmt.Scanln(&id)
	fmt.Println("Daftar Pelanggan yang membeli Sparepart ID:", id)
	found := false
	for i := 0; i < *tData; i++ {
		if T[i].IDSparepart == id {
			for j := 0; j < *pData; j++ {
				if P[j].ID == T[i].IDPelanggan {
					fmt.Println("------------------------------------------------")
					fmt.Println("ID Pelanggan:", P[j].ID)
					fmt.Println("Nama Pelanggan:", P[j].Nama)
					fmt.Println("Total Pembelian:", T[i].Jumlah)
					fmt.Println("Harga:", T[i].Harga)
					fmt.Println("Pajak:", T[i].Pajak)
					fmt.Println("Total Harga:", T[i].TotalHarga)
					found = true
				}
			}
		}
	}
	if !found {
		fmt.Println("Tidak ada pelanggan yang membeli sparepart dengan ID tersebut.")
	}
	login_admin(S, P, T, sData, pData, tData)
}

func getHargaSparepart(S *spareparts, id, sData int) int {
	for i := 0; i < sData; i++ {
		if S[i].ID == id {
			return S[i].Harga
		}
	}
	return 0
}

func checkPelangganExist(P *pelanggan, id, pData int) bool {
	for i := 0; i < pData; i++ {
		if P[i].ID == id {
			return true
		}
	}
	return false
}

func checkSparepartExist(S *spareparts, id, sData int) bool {
	kiri := 0
	kanan := sData - 1
	tengah := (kiri + kanan) / 2
	for kiri <= kanan {
		if S[tengah].ID == id {
			return true
		} else if S[tengah].ID < id {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
		tengah = (kiri + kanan) / 2
	}
	return false
}

func edit_data_admin(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	fmt.Println("==============================================")
	fmt.Println("                   EDIT DATA               ")
	fmt.Println("==============================================")
	fmt.Println("  1. Sparepart               ")
	fmt.Println("  2. Pelanggan               ")
	fmt.Println("  3. Transaksi               ")
	fmt.Println("  4. Kembali                               ")
	fmt.Println("==============================================")
	fmt.Print("Pilih: ")
	var pilih int
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		edit_sparepart(S, sData, P, T, pData, tData)
	case 2:
		edit_pelanggan(S, P, T, sData, pData, tData)
	case 3:
		edit_transaksi(S, P, T, sData, pData, tData)
	case 4:
		login_admin(S, P, T, sData, pData, tData)
	default:
		edit_data_admin(S, P, T, sData, pData, tData)
	}
}

func edit_sparepart(S *spareparts, sData *int, P *pelanggan, T *transaksi, pData, tData *int) {
	var id int
	fmt.Print("Masukan ID Sparepart yang akan diubah: ")
	fmt.Scanln(&id)
	for i := 0; i < *sData; i++ {
		if S[i].ID == id {
			fmt.Print("Masukan Nama Sparepart baru: ")
			S[i].Nama = scan_word()
			fmt.Print("Masukan Harga Sparepart baru: ")
			fmt.Scanln(&S[i].Harga)
			fmt.Println("Data Sparepart berhasil diubah!")
			login_admin(S, P, T, sData, pData, tData)
			return
		}
	}
	fmt.Println("Data Sparepart tidak ditemukan!")
	login_admin(S, P, T, sData, pData, tData)
}

func edit_pelanggan(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	var id int
	fmt.Print("Masukan ID Pelanggan yang akan diubah: ")
	fmt.Scanln(&id)
	for i := 0; i < *pData; i++ {
		if P[i].ID == id {
			fmt.Print("Masukan Nama Pelanggan baru: ")
			P[i].Nama = scan_word()
			fmt.Println("Data Pelanggan berhasil diubah!")
			login_admin(S, P, T, sData, pData, tData)
			return
		}
	}
	fmt.Println("Data Pelanggan tidak ditemukan!")
	login_admin(S, P, T, sData, pData, tData)
}

func edit_transaksi(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	var id int
	fmt.Print("Masukan ID Transaksi yang akan diubah: ")
	fmt.Scanln(&id)
	for i := 0; i < *tData; i++ {
		if T[i].IDPelanggan == id {
			fmt.Print("Masukan ID Sparepart baru: ")
			fmt.Scanln(&T[i].IDSparepart)
			fmt.Print("Masukan Jumlah baru: ")
			fmt.Scanln(&T[i].Jumlah)
			T[i].Harga = T[i].Jumlah * getHargaSparepart(S, T[i].IDSparepart, *sData)
			T[i].Pajak = T[i].Harga * 5 / 100
			T[i].TotalHarga = T[i].Harga + T[i].Pajak
			fmt.Println("Data Transaksi berhasil diubah!")
			login_admin(S, P, T, sData, pData, tData)
			return
		}
	}
	fmt.Println("Data Transaksi tidak ditemukan!")
	login_admin(S, P, T, sData, pData, tData)
}

func hapus_data_admin(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	fmt.Println("==============================================")
	fmt.Println("                   HAPUS DATA               ")
	fmt.Println("==============================================")
	fmt.Println("  1. Sparepart               ")
	fmt.Println("  2. Pelanggan               ")
	fmt.Println("  3. Transaksi               ")
	fmt.Println("  4. Kembali                               ")
	fmt.Println("==============================================")
	fmt.Print("Pilih: ")
	var pilih int
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		hapus_sparepart(S, sData, P, T, pData, tData)
	case 2:
		hapus_pelanggan(S, P, T, sData, pData, tData)
	case 3:
		hapus_transaksi(S, P, T, sData, pData, tData)
	case 4:
		login_admin(S, P, T, sData, pData, tData)
	default:
		hapus_data_admin(S, P, T, sData, pData, tData)
	}
}

func hapus_sparepart(S *spareparts, sData *int, P *pelanggan, T *transaksi, pData, tData *int) {
	var id int
	fmt.Print("Masukan ID Sparepart yang akan dihapus: ")
	fmt.Scanln(&id)
	for i := 0; i < *sData; i++ {
		if S[i].ID == id {
			for j := i; j < *sData-1; j++ {
				S[j] = S[j+1]
			}
			(*sData)--
			fmt.Println("Data Sparepart berhasil dihapus!")
			login_admin(S, P, T, sData, pData, tData)
			return
		}
	}
	fmt.Println("Data Sparepart tidak ditemukan!")
	login_admin(S, P, T, sData, pData, tData)
}

func hapus_pelanggan(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	var id int
	fmt.Print("Masukan ID Pelanggan yang akan dihapus: ")
	fmt.Scanln(&id)
	for i := 0; i < *pData; i++ {
		if P[i].ID == id {
			for j := i; j < *pData-1; j++ {
				P[j] = P[j+1]
			}
			(*pData)--
			fmt.Println("Data Pelanggan berhasil dihapus!")
			login_admin(S, P, T, sData, pData, tData)
			return
		}
	}
	fmt.Println("Data Pelanggan tidak ditemukan!")
	login_admin(S, P, T, sData, pData, tData)
}

func hapus_transaksi(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	var id int
	fmt.Print("Masukan ID Pelanggan yang ingin Transaksinya dihapus: ")
	fmt.Scanln(&id)
	for i := 0; i < *tData; i++ {
		if T[i].IDPelanggan == id {
			for j := i; j < *tData-1; j++ {
				T[j] = T[j+1]
			}
			(*tData)--
			fmt.Println("Data Transaksi berhasil dihapus!")
			login_admin(S, P, T, sData, pData, tData)
			return
		}
	}
	fmt.Println("Data Transaksi tidak ditemukan!")
	login_admin(S, P, T, sData, pData, tData)
}

func login_pembeli(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	fmt.Println("==============================================")
	fmt.Println("        SELAMAT DATANG DI SISTEM PEMBELI     ")
	fmt.Println("==============================================")
	fmt.Println("                 Pilihan Menu                  ")
	fmt.Println("----------------------------------------------")
	fmt.Println("1. Beli Sparepart motor")
	fmt.Println("2. Cari Sparepart motor")
	fmt.Println("3. Lihat daftar harga Sparepart motor")
	fmt.Println("4. Kembali")
	fmt.Println("==============================================")
	fmt.Print("Pilih: ")
	var pilih int
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		beli_sparepart(S, P, T, sData, pData, tData)
	case 2:
		cari_sparepart(S, sData, P, T, pData, tData)
	case 3:
		daftar_sparepart(S, sData, P, T, pData, tData)
	case 4:
		menu(S, P, T, sData, pData, tData)
	default:
		login_pembeli(S, P, T, sData, pData, tData)
	}
}

func beli_sparepart(S *spareparts, P *pelanggan, T *transaksi, sData, pData, tData *int) {
	var transaksi Transaksi
	fmt.Print("Masukan ID Pelanggan: ")
	fmt.Scanln(&transaksi.IDPelanggan)
	if !checkPelangganExist(P, transaksi.IDPelanggan, *pData) {
		fmt.Println("ID Pelanggan tidak ditemukan.")
		login_pembeli(S, P, T, sData, pData, tData)
		return
	}
	fmt.Print("Masukan ID Sparepart: ")
	fmt.Scanln(&transaksi.IDSparepart)
	if !checkSparepartExist(S, transaksi.IDSparepart, *sData) {
		fmt.Println("ID Sparepart tidak ditemukan.")
		login_pembeli(S, P, T, sData, pData, tData)
		return
	}
	fmt.Print("Masukan Jumlah: ")
	fmt.Scanln(&transaksi.Jumlah)
	transaksi.Harga = transaksi.Jumlah * getHargaSparepart(S, transaksi.IDSparepart, *sData)
	transaksi.Pajak = transaksi.Harga * 5 / 100
	transaksi.TotalHarga = transaksi.Harga + transaksi.Pajak
	T[*tData] = transaksi
	(*tData)++

	for i := 0; i < *sData; i++ {
		if S[i].ID == transaksi.IDSparepart {
			S[i].Terjual += transaksi.Jumlah
			i = *sData
		}
	}
	fmt.Println("Pembelian Berhasil!     ")
	fmt.Println()
	fmt.Println("===========================")
	fmt.Println("        STRUK BELANJA       ")
	fmt.Println("===========================")
	fmt.Printf("Harga        : Rp %10d\n", transaksi.Harga)
	fmt.Printf("Pajak        : Rp %10d\n", transaksi.Pajak)
	fmt.Printf("---------------------------\n")
	fmt.Printf("Total Harga  : Rp %10d\n", transaksi.TotalHarga)
	fmt.Println("===========================")
	fmt.Println("       Terima Kasih!")
	fmt.Println("===========================")
	fmt.Println()
	login_pembeli(S, P, T, sData, pData, tData)
}

func cari_sparepart(S *spareparts, sData *int, P *pelanggan, T *transaksi, pData, tData *int) {
	var id int
	fmt.Print("Masukan ID Sparepart yang dicari: ")
	fmt.Scanln(&id)
	kiri := 0
	kanan := *sData - 1
	tengah := (kiri + kanan) / 2
	for kiri <= kanan {
		if S[tengah].ID == id {
			fmt.Println("Sparepart ditemukan!")
			fmt.Println("ID:", S[tengah].ID)
			fmt.Println("Nama:", S[tengah].Nama)
			fmt.Println("Harga:", S[tengah].Harga)
			fmt.Println("Terjual:", S[tengah].Terjual)
			login_pembeli(S, P, T, sData, pData, tData)
			return
		} else if S[tengah].ID < id {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
		tengah = (kiri + kanan) / 2
	}
	fmt.Println("Sparepart tidak ditemukan!")
	login_pembeli(S, P, T, sData, pData, tData)
}

func daftar_sparepart(S *spareparts, sData *int, P *pelanggan, T *transaksi, pData, tData *int) {
	var pilih int
	fmt.Println("Daftar Harga Sparepart: ")
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("| %-3s | %-20s | %-10s | %-7s |\n", "ID", "Nama Sparepart", "Harga", "Terjual")
	fmt.Println("-----------------------------------------------------")
	for i := 0; i < *sData; i++ {
		fmt.Printf("| %-3d | %-20s | %-10d | %-7d |\n", S[i].ID, S[i].Nama, S[i].Harga, S[i].Terjual)
	}
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1. Urutkan sparepart dari Pembeli Terbanyak")
	fmt.Println("2. Urutkan sparepart dari Pembeli Terdikit")
	fmt.Println("3. Kembali")
	fmt.Println("------------------------------------------------")
	fmt.Print("Pilih: ")
	fmt.Scanln(&pilih)
	switch pilih {
	case 1:
		descending_sparepart(S, sData)
		daftar_sparepart(S, sData, P, T, pData, tData)
	case 2:
		ascending_sparepart(S, sData)
		daftar_sparepart(S, sData, P, T, pData, tData)
	case 3:
		login_pembeli(S, P, T, sData, pData, tData)
	default:
		daftar_sparepart(S, sData, P, T, pData, tData)
	}
}

func descending_sparepart(S *spareparts, sData *int) {
	for i := 1; i < *sData; i++ {
		key := S[i]
		j := i - 1
		for j >= 0 && S[j].Terjual < key.Terjual {
			S[j+1] = S[j]
			j = j - 1
		}
		S[j+1] = key
	}
}

func ascending_sparepart(S *spareparts, sData *int) {
	for i := 0; i < *sData-1; i++ {
		minIdx := i
		for j := i + 1; j < *sData; j++ {
			if S[j].Terjual < S[minIdx].Terjual {
				minIdx = j
			}
		}
		S[i], S[minIdx] = S[minIdx], S[i]
	}
}

func scan_word() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word := scanner.Text()
	return word
}

func selectionSort(arr []Sparepart) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].ID < arr[minIdx].ID {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
