package main

import "fmt"

func main() {
	fmt.Println("masukan nama anda : ")
	var (
		nama   string
		umur   int
		agama  string
		alamat string
	)
	fmt.Scanln(&nama)
	fmt.Println("masukan umur anda : ")
	fmt.Scanln(&umur)
	fmt.Println("masukan agama anda : ")
	fmt.Scanln(&agama)
	fmt.Println("masukan alamat anda : ")
	fmt.Scanln(&alamat)
	fmt.Println("nama anda adalah : ", nama)
	fmt.Println("umur anda adalah : ", umur)
	switch agama {
	case "islam":
		fmt.Println("anda beragama : ", agama)
	case "non islam":
		fmt.Println("rajin ke beribadah ketempat ibadah sesuai kepercayaan anda ya! yaitu: : ", agama)
	default:
		fmt.Println("nggak ada agamanya itu")
	}
	fmt.Println("alamat anda  : ", alamat)
	// fmt.Println("assalamualaikum")
}
