// Package redovalnica implements a simple grade book system for managing students and their grades.
//
// It allows adding students, recording their grades, calculating averages, and printing reports.
//
// Example usage:
//
//	redovalnica := redovalnica.NewRedovalnica(0, 10, 1)
//	redovalnica.DodajStudenta("63210001", "Janez", "Novak")
//	redovalnica.DodajOceno("63210001", 10)
//	redovalnica.DodajOceno("63210001", 9)
//	redovalnica.IzpisVsehOcen()
//	redovalnica.IzpisiKoncniUspeh()
//
// The grade book supports the following features:
//   - Adding students with unique ID numbers
//   - Recording grades within a specified range
//   - Calculating average grades for each student
//   - Printing all grades and final performance reports
package redovalnica

import "fmt"

// Student represents a student with a name, surname, and a list of grades.
type Student struct {
	ime     string
	priimek string
	ocene   []int
}

// Redovalnica represents a grade book managing multiple students and their grades.
type Redovalnica struct {
	studenti map[string]Student
	minOcena int
	maxOcena int
	stOcen   int
}

// NewRedovalnica creates a new grade book with specified minimum and maximum grades and required number of grades for average calculation.
func NewRedovalnica(minOcena int, maxOcena int, stOcen int) *Redovalnica {
	return &Redovalnica{
		studenti: make(map[string]Student),
		minOcena: minOcena,
		maxOcena: maxOcena,
		stOcen:   stOcen,
	}
}

// DodajStudenta adds a new student to the grade book.
func (r *Redovalnica) DodajStudenta(vpisnaStevilka string, ime string, priimek string) {
	_, ok := r.studenti[vpisnaStevilka]
	if ok {
		fmt.Println("Študent s to vpisno številko že obstaja")
		return
	}
	r.studenti[vpisnaStevilka] = Student{
		ime:     ime,
		priimek: priimek,
		ocene:   []int{},
	}
}

// DodajOceno adds a grade to the specified student in the grade book.
func (r *Redovalnica) DodajOceno(vpisnaStevilka string, ocena int) {
	if ocena < r.minOcena || ocena > r.maxOcena {
		fmt.Printf("Ocena mora biti med %d in %d\n", r.minOcena, r.maxOcena)
		return
	}
	student, ok := r.studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študent s to vpisno številko ne obstaja")
		return
	}
	student.ocene = append(student.ocene, ocena)
	r.studenti[vpisnaStevilka] = student
}

// povprecje calculates the average grade for the specified student.
func (r *Redovalnica) povprecje(vpisnaStevilka string) float64 {
	student, ok := r.studenti[vpisnaStevilka]
	if !ok {
		fmt.Println("Študent s to vpisno številko ne obstaja")
		return -1.0
	}
	sum := 0
	for _, ocena := range student.ocene {
		sum += ocena
	}
	if len(student.ocene) < r.stOcen {
		return 0.0
	}
	avg := float64(sum) / float64(len(student.ocene))
	return avg
}

// IzpisVsehOcen prints all grades for each student in the grade book.
func (r *Redovalnica) IzpisVsehOcen() {
	fmt.Println("REDOVALNICA:")
	for vpisnaStevilka, student := range r.studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisnaStevilka, student.ime, student.priimek, student.ocene)
	}
}

// IzpisiKoncniUspeh prints the final performance report for each student based on their average grade.
func (r *Redovalnica) IzpisiKoncniUspeh() {
	fmt.Println("KONČNI USPEH:")
	for vpisnaStevilka, student := range r.studenti {
		avg := r.povprecje(vpisnaStevilka)
		if avg >= (0.9 * (float64(r.maxOcena) - float64(r.minOcena))) {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Odličen študent!\n", student.ime, student.priimek, avg)
		} else if avg >= (0.6 * (float64(r.maxOcena) - float64(r.minOcena))) {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Povprečen študent\n", student.ime, student.priimek, avg)
		} else {
			fmt.Printf("%s %s: povprečna ocena %.1f -> Neuspešen študent\n", student.ime, student.priimek, avg)
		}
	}
}
