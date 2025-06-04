package main
import "fmt"

const NMAX int = 99
type proyek struct {
	nama string
	target, donasi, donatur int
}

type tabProyek [NMAX]proyek 
var data tabProyek
var nData int

func main() {
	menu()
}

func menu() {
	
	var pilihan int
	
	for pilihan != 11 {
		fmt.Println("\n\n----------------------------CROWDFUNDING APP---------------------------")
		fmt.Println("1.  TAMBAH PROYEK")
		fmt.Println("2.  TAMPILKAN SEMUA PROYEK")
		fmt.Println("3.  TAMBAH DONASI KE PROYEK")
		fmt.Println("4.  CARI PROYEK (SEQUENTIAL SEARCH)")
		fmt.Println("5.  CARI PROYEK (BINARY SEARCH)")
		fmt.Println("6.  URUTKAN BERDASARKAN DANA (SELECTION SORT)")
		fmt.Println("7.  URUTKAN BERDASARKAN DONATUR (INSERTION SORT)")
		fmt.Println("8.  TAMPILKAN PROYEK YANG SUDAH MENCAPAI TARGET")
		fmt.Println("9.  HAPUS PROYEK")
		fmt.Println("10. UBAH PROYEK")
		fmt.Println("11. KELUAR")
		fmt.Print("PILIH MENU: ")
		
		fmt.Scan(&pilihan)
	
		switch pilihan {
			case 1:
				tambahProyek(&data, &nData)
			case 2:
				tampilkanSemuaProyek(data, nData)
			case 3:
				tambahDonasi(&data, nData)
			case 4:
				sequentialSearch(data, nData)
			case 5:
				binarySearch(data, nData)
			case 6:
				selectionSort(&data, nData)
			case 7:
				insertionSort(&data, nData)
			case 8:
				tampilkanProyekYangSudahMencapaiTarget(data, nData)
			case 9:
				hapusProyek(&data, &nData)
			case 10:
				ubahProyek(&data, nData)
			case 11:
				fmt.Println("TERIMA KASIH... ")
			default:
				fmt.Println("MENU TIDAK DITEMUKAN")
		}
	}
}

func tambahProyek(A *tabProyek, n *int) { 
	fmt.Print("NAMA PROYEK: ")
	fmt.Scan(&(*A)[*n].nama)
	fmt.Print("TARGET DANA: ")
	fmt.Scan(&(*A)[*n].target)
	*n++
	fmt.Println("PROYEK BERHASIL DITAMBAHKAN")
}

func tampilkanSemuaProyek(A tabProyek, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s (TARGET: %d, TERKUMPUL: %d, DONATUR; %d)\n", i+1, A[i].nama, A[i].target, A[i].donasi, A[i].donatur)
	}
}

func tambahDonasi(A *tabProyek, n int) {
	var nama string
	var jumlah int
	
	fmt.Print("NAMA PROYEK: ")
	fmt.Scan(&nama)
	fmt.Print("JUMLAH DONASI: ")
	fmt.Scan(&jumlah)
	
	for i := 0; i < n; i++ {
		if nama == (*A)[i].nama {
			(*A)[i].donasi += jumlah
			(*A)[i].donatur += 1
			fmt.Println("DONASI BERHASIL DITAMBAHKAN")
			return
		} 
	}
	fmt.Println("PROYEK TIDAK DITEMUKAN")

}

func tampilkanProyekYangSudahMencapaiTarget(A tabProyek, n int) {
	for i := 0; i < n; i++ {
		if A[i].donasi >= A[i].target {
			fmt.Println(A[i].nama, " SUDAH MENCAPAI TARGET")
		}
	}
}

func binarySearch(A tabProyek, n int) {
    var nama string
	fmt.Print("MASUKKAN NAMA PROYEK: ")
	fmt.Scan(&nama)

	// Sort terlebih dahulu 
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if A[j].nama > A[j+1].nama {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
	}

	// Binary Search
	kiri := 0
	kanan := n - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if A[tengah].nama == nama {
			fmt.Printf("DITEMUKAN: %s (TARGET: %d, DONASI: %d, DONATUR: %d)\n",
				A[tengah].nama, A[tengah].target, A[tengah].donasi, A[tengah].donatur)
			return
		} else if A[tengah].nama < nama {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	fmt.Println("PROYEK TIDAK DITEMUKAN.")
}

func sequentialSearch(A tabProyek, n int) {
	var nama string
	var i int
	
	fmt.Print("MASUKKAN NAMA PROYEK: ")
	fmt.Scan(&nama)
	for i = 0; i < n; i++ {
		if A[i].nama == nama {
			fmt.Printf("DITEMUKAN: %s (TARGET: %d, DONASI: %d, DONATUR: %d)\n", A[i].nama, A[i].target, A[i].donasi, A[i].donatur)
			return
		}
	}
	fmt.Print("PROYEK TIDAK DITEMUKAN")
}

func selectionSort(A *tabProyek, n int) {
	var pass, idx, i int
	var temp proyek
	
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].donasi > A[idx].donasi {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass+1
	}
	tampilkanSemuaProyek(data, nData)
}

func insertionSort(A *tabProyek, n int) {
	var pass, i int
	var temp proyek
	
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.donatur > A[i-1].donatur {
			A[i] = A[i-1]
			i = i-1
		}
		A[i] = temp
		pass = pass + 1
	}
	tampilkanSemuaProyek(data, nData)

}

func hapusProyek(A *tabProyek, n *int) {
	var nama string
	fmt.Print("NAMA PROYEK YANG INGIN DIHAPUS: ")
	fmt.Scan(&nama)
	
	for i := 0; i < *n; i++ {
		if A[i].nama == nama {
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			fmt.Println("PROYEK BERHASIL DIHAPUS")
			return
		}
	}
}

func ubahProyek(A *tabProyek, n int) {
	var namaYangInginDiubah string
	var i int
	var ditemukan bool = false
	
	fmt.Print("NAMA PROYEK YANG INGIN DIUBAH: ")
	fmt.Scan(&namaYangInginDiubah)
	
	for i < n && !ditemukan {
		if A[i].nama == namaYangInginDiubah {
			fmt.Print("NAMA BARU: ")
			fmt.Scan(&(*A)[i].nama)
			fmt.Print("TARGET BARU: ")
			fmt.Scan(&(*A)[i].target)
			fmt.Println("PROYEK BERHASIL DIUBAH")
			ditemukan = true
		} else {
			i++
		}
	}
	if !ditemukan {
		fmt.Println("PROYEK TIDAK DITEMUKAN")
	}
}


