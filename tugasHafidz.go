package main

import (
	"encoding/json"
	"fmt"
)

type propertyListing struct {
	Nama      string `json:"Nama_go"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"Pekerjaan"`
	Umur      int    `json:"Umur"`
	Deskripsi string `json:"Deskripsi"`
}

func CekData(Nama_go string) (propertyListing, error) {

	var dProp []propertyListing
	data := `[
		{
			"Nama_go":"M.Hafidz Dasuqi",
			"alamat":"Jl. Hangtuah",
			"Pekerjaan":"Mahasiswa",
			"Umur": 21,
			"Deskripsi":"manambah ilmu dan ingin memperdalam bahasa pemograman golang"

		},
		{
			"Nama_go":"Agus Arya",
			"alamat":"Jl . boyolali",
			"Pekerjaan":"Mahasiswa",
			"Umur": 20,
			"Deskripsi":"mendalamin ilmu golang"
		},
		{
			"Nama_go":"Devo Mahaswara",
			"alamat":"jogja",
			"Pekerjaan":"Mahasiswa && junior developer",
			"Umur": 23,
			"Deskripsi":"Ingin mencari sesuatu yang baru"
		},
		{
			"Nama_go":"Fajrian Nugraha",
			"alamat":"jombang",
			"Pekerjaan":"Mahasiswa && junior developer",
			"Umur": 22,
			"Deskripsi":"Ingin mencari sesuatu yang baru"
		},
		{
			"Nama_go":"Moohammad HAnif",
			"alamat":"Semarang",
			"Pekerjaan":"Mahasiswa && junior developer",
			"Umur": 22,
			"Deskripsi":"Ingin mencari sesuatu yang baru"
		},
		{
			"Nama_go":"Hanif Fadilah",
			"alamat":"purwokerto",
			"Pekerjaan":"Mahasiswa && junior developer",
			"Umur": 21,
			"Deskripsi":"Ingin mencari sesuatu yang baru"
		},
		{
			"Nama_go":"Wahyu Prihatono",
			"alamat":"jogja",
			"Pekerjaan":"Mahasiswa && junior developer",
			"Umur": 20,
			"Deskripsi":"Ingin mencari sesuatu yang baru"
		},
		{
			"Nama_go":"Maulita",
			"alamat":"Kediri",
			"Pekerjaan":"Mahasiswa && junior developer",
			"Umur": 20,
			"Deskripsi":"Ingin mencari sesuatu yang baru"
		},
		{
			"Nama_go":"Gidan",
			"alamat":"cirebon",
			"Pekerjaan":"Mahasiswa && junior developer",
			"Umur": 22,
			"Deskripsi":"Ingin mencari sesuatu yang baru"
		},
		{
			"Nama_go":"Tasya Gracinia",
			"alamat":"Bali",
			"Pekerjaan":"Mahasiswa && junior developer",
			"Umur": 21,
			"Deskripsi":"Ingin mencari sesuatu yang baru"
		},
		{
			"Nama_go":"Wandhie",
			"alamat":"jakarta",
			"Pekerjaan":"Mahasiswa && junior developer",
			"Umur": 21,
			"Deskripsi":"Ingin mencari sesuatu yang baru"
		},
		{
			"Nama_go":"Putry Adetya Azhari",
			"alamat":"jakarta",
			"Pekerjaan":"Mahasiswa && junior developer",
			"Umur": 21,
			"Deskripsi":"Ingin mencari sesuatu yang baru"
		}
	]`

	textBytes := []byte(data)
	err := json.Unmarshal(textBytes, &dProp)
	if err != nil {
		return propertyListing{}, err
	}

	for i := 0; i < len(dProp); i++ {
		if dProp[i].Nama == Nama_go {
			return dProp[i], nil
		}
	}
	return propertyListing{}, err
}

func main() {
	d, err := CekData("Gidan") //input
	if err != nil {
		fmt.Println("data error")
	}
	fmt.Println("Nama          :", d.Nama)
	fmt.Println("alamat        :", d.Alamat)
	fmt.Println("Pekerjaan     :", d.Pekerjaan)
	fmt.Println("Usia          :", d.Umur)
	fmt.Println("Alasan Masuk  :", d.Deskripsi)
}
