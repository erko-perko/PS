package main

import "fmt"

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func main() {
	var studenti map[string]Student
	studenti = make(map[string]Student)

	studenti = map[string]Student{
		"63210001": {"Ana", "Novak", []int{10, 9, 8}},
		"63210002": {"Boris", "Kralj", []int{6, 7, 5, 8}},
		"63210003": {"Janez", "Novak", []int{4, 5, 3, 5}},
		"63210004": {"Maja", "Horvat", []int{9, 10, 10, 9, 8, 10}},
		"63210005": {"Luka", "Zupan", []int{7, 6, 8, 7, 6, 5, 7}},
	}

	fmt.Println("Originalni slovar:")
	fmt.Println(studenti)

	fmt.Println()
	fmt.Println("Običajno dodajanje:")
	dodajOceno(studenti, "63210002", 9)
	fmt.Println(studenti)

	fmt.Println()
	fmt.Println("Neveljavna ocena:")
	dodajOceno(studenti, "63210002", 11) // neveljavna ocena
	fmt.Println(studenti)

	fmt.Println()
	fmt.Println("Neobstoječ študent:")
	dodajOceno(studenti, "63210000", 9) // neobstojec student
	fmt.Println(studenti)

	fmt.Println()
	fmt.Println("Običajno računanje povprečja:")
	vpisnaStevilka := "63210004"
	student, _ := studenti[vpisnaStevilka]
	avg := povprecje(studenti, vpisnaStevilka)
	fmt.Printf("Povprečje študenta %s %s je %.2f\n", student.ime, student.priimek, avg)

	fmt.Println()
	fmt.Println("Premalo ocen:")
	vpisnaStevilka = "63210001" // premalo ocen
	student, _ = studenti[vpisnaStevilka]
	avg = povprecje(studenti, vpisnaStevilka)
	fmt.Printf("Povprečje študenta %s %s je %.2f\n", student.ime, student.priimek, avg)

	fmt.Println()
	fmt.Println("Neobstoječ študent:")
	vpisnaStevilka = "63210000" // neobstojec student
	student, _ = studenti[vpisnaStevilka]
	avg = povprecje(studenti, vpisnaStevilka)
	fmt.Printf("Vrnjeno povprečje je %.2f\n", avg)

	fmt.Println()
	izpisRedovalnice(studenti)

	fmt.Println()
	izpisiKoncniUspeh(studenti)
}
func dodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena < 0 || ocena > 10 {
		fmt.Println("Ocena mora biti med 0 in 10")
		return
	}
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študent s to vpisno številko ne obstaja")
		return
	}
	student.ocene = append(student.ocene, ocena)
	studenti[vpisnaStevilka] = student
}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študent s to vpisno številko ne obstaja")
		return -1.0
	}
	sum := 0
	for _, ocena := range student.ocene {
		sum += ocena
	}
	if len(student.ocene) < 6 {
		return 0.0
	}
	avg := float64(sum) / float64(len(student.ocene))
	return avg
}

func izpisRedovalnice(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisnaStevilka, student := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisnaStevilka, student.ime, student.priimek, student.ocene)
	}
}

func izpisiKoncniUspeh(studenti map[string]Student) {
	fmt.Println("KONČNI USPEH:")
	for vpisnaStevilka, student := range studenti {
		avg := povprecje(studenti, vpisnaStevilka)
		if avg >= 9.0 {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Odličen študent!\n", student.ime, student.priimek, avg)
		} else if avg >= 6.0 {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Povprečen študent\n", student.ime, student.priimek, avg)
		} else {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Neuspešen študent\n", student.ime, student.priimek, avg)
		}
	}
}
