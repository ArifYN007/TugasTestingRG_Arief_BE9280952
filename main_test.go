package main_test

import (
	//PERLU DINGAT KETIKA MENGIMPORT FUNGSI DARI MAIN.GO FUNCTION TERSEBUT DIAWAL HARUS DIAWALI KAPITAL!!
	. "TugasTestingRG"
	"fmt"
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
	Describe("multiplication", func() {
		When("one or both operands are empty", func() {
			It("should return an error", func() {
				_, err := Multiplication("", "") //EMANGGIL FUNGSI DENGAN OOPERAND KOSONG
				Expect(err).To(HaveOccurred())   //MENGGUNAKAN GOMEGA KETIKA ERROR= ERROR MAKA PENGUJIAN BERHASIL
			})
		})
		When("one or both operands are not valid numbers", func() {
			It("should return an error", func() {
				_, err := Multiplication("abc", "def")
				Expect(err).To(HaveOccurred())
			})
		})
		When("both operands valid numbers", func() {
			It("should return the correct result", func() {
				result, err := Multiplication("2", "3")
				Expect(err).ToNot(HaveOccurred())               //memerika jika error tidak berisi error maka pengujian berhasil
				Expect(result).To(Equal("Hasil adalah : 6.00")) //memeriska hasil result apakah sesuai dengan harapan orang tua
			})
		})
	})
	Describe("division", func() {
		When("one or both operands are empty", func() {
			It("should return an error", func() {
				_, err := Division("", "")
				Expect(err).To(HaveOccurred())
			})
		})
		When("one or both operands are not valid numbers", func() {
			It("should return an error", func() {
				_, err := Division("", "def")
				Expect(err).To(HaveOccurred())
			})
		})
		When("both operands valid numbers", func() {
			It("should return the correct result", func() {
				result, err := Division("2", "3")
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal("Hasil adalah : 0.67"))
			})
			It("should return the correct result", func() {
				result, err := Division("3", "3")
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal("Hasil adalah : 1.00"))
			})
		})
	})
	Describe("exponent", func() {
		When("one or both operands are empty", func() {
			It("should return an error", func() {
				_, err := Exponen("", "")
				Expect(err).To(HaveOccurred())
			})
		})
		When("one or both operands are not valid numbers", func() {
			It("should return an error", func() {
				_, err := Exponen("adfd", "")
				Expect(err).To(HaveOccurred())
			})
		})
		When("both operands valid numbers", func() {
			It("should return the correct result", func() {
				result, err := Exponen("10", "11")
				Expect(err).ToNot(HaveOccurred())
				expected := math.Pow(10, 11) //hitung hasil
				Expect(result).To(Equal(fmt.Sprintf("Hasil adalah : %.2f", expected)))
			})
		})
	})
	Describe("sqrt", func() {
		When("one or both operands are empty", func() {
			It("should return an error", func() {
				_, err := Sqrt("")
				Expect(err).To(HaveOccurred())
			})
		})
		When("one or both operands are not valid numbers", func() {
			It("should return an error", func() {
				_, err := Sqrt("wkwk")
				Expect(err).To(HaveOccurred())
			})
		})
		When("both operands valid numbers", func() {
			It("should return the correct result", func() {
				result, err := Sqrt("10")
				Expect(err).ToNot(HaveOccurred())
				expected := math.Sqrt(10) //hitung hasil
				Expect(result).To(Equal(fmt.Sprintf("Hasil adalah : %.2f", expected)))
			})
		})
	})
})
