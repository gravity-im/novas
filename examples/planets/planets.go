package main

import (
	"github.com/pebbe/novas"

	"fmt"
)

func main() {

	latitude, longitude := 53.21853, 6.5670 // Groningen, The Netherlands

	// END OF USER SETTINGS

	fmt.Println(novas.EphInfo(), "\n")

	now := novas.Now()
	fmt.Println(now)

	geo := novas.NewPlace(latitude, longitude, 0, 20, 1010)
	fmt.Println("\nLocation:", geo)

	fmt.Println("\n            Distance   Altitude   Azimuth   Disc")
	for _, obj := range []*novas.Body{
		novas.Mercury(),
		novas.Venus(),
		novas.Mars(),
		novas.Jupiter(),
		novas.Saturn(),
		novas.Uranus(),
		novas.Neptune(),
		novas.Pluto(),
	} {
		data := obj.Topo(now, geo, novas.REFR_NONE)
		fmt.Printf("%-8s%12.6f%11.2f%10.2f%7.0f\n", obj.Name(), data.Dis, data.Alt, data.Az, obj.Disc(now))
	}
}
