package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/erko-perko/redovalnica-PS/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Upravljanje z redovalnico študentov",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Najmanjše število ocen potrebnih za pozitivno oceno",
				Value: 6,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Najmanjša možna ocena",
				Value: 0,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Največja možna ocena",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")
			return runRedovalnica(stOcen, minOcena, maxOcena)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func runRedovalnica(stOcen int, minOcena int, maxOcena int) error {
	redovalnica := redovalnica.NewRedovalnica(minOcena, maxOcena, stOcen)

	redovalnica.DodajStudenta("63210001", "Ana", "Novak")
	redovalnica.DodajStudenta("63210002", "Boris", "Kralj")
	redovalnica.DodajStudenta("63210003", "Janez", "Novak")
	redovalnica.DodajStudenta("63210004", "Maja", "Horvat")
	redovalnica.DodajStudenta("63210005", "Luka", "Zupan")
	redovalnica.DodajOceno("63210001", 10)
	redovalnica.DodajOceno("63210001", 9)
	redovalnica.DodajOceno("63210001", 8)
	redovalnica.DodajOceno("63210002", 6)
	redovalnica.DodajOceno("63210002", 7)
	redovalnica.DodajOceno("63210002", 5)
	redovalnica.DodajOceno("63210002", 8)
	redovalnica.DodajOceno("63210003", 4)
	redovalnica.DodajOceno("63210003", 5)
	redovalnica.DodajOceno("63210003", 3)
	redovalnica.DodajOceno("63210003", 5)
	redovalnica.DodajOceno("63210004", 9)
	redovalnica.DodajOceno("63210004", 10)
	redovalnica.DodajOceno("63210004", 10)
	redovalnica.DodajOceno("63210004", 9)
	redovalnica.DodajOceno("63210004", 8)
	redovalnica.DodajOceno("63210004", 10)
	redovalnica.DodajOceno("63210005", 7)
	redovalnica.DodajOceno("63210005", 6)
	redovalnica.DodajOceno("63210005", 8)
	redovalnica.DodajOceno("63210005", 7)
	redovalnica.DodajOceno("63210005", 6)
	redovalnica.DodajOceno("63210005", 5)
	redovalnica.DodajOceno("63210005", 7)

	fmt.Println("Originalni slovar:")
	redovalnica.IzpisVsehOcen()

	fmt.Println()
	fmt.Println("Običajno dodajanje:")
	redovalnica.DodajOceno("63210002", maxOcena-1)
	redovalnica.IzpisVsehOcen()

	fmt.Println()
	fmt.Println("Neveljavna ocena:")
	redovalnica.DodajOceno("63210002", maxOcena+1) // neveljavna ocena
	redovalnica.IzpisVsehOcen()

	fmt.Println()
	fmt.Println("Neobstoječ študent:")
	redovalnica.DodajOceno("63210000", maxOcena-1) // neobstojec student

	fmt.Println()
	redovalnica.IzpisVsehOcen()

	fmt.Println()
	redovalnica.IzpisiKoncniUspeh()

	return nil
}
