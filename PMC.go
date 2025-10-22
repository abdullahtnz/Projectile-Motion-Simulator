package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func PMCalculator(height, angle, velocity, mass int, ar bool) {
	const g float64 = 9.81
	const dt float64 = 0.01

	h := float64(height)
	a := float64(angle)
	v := float64(velocity)
	m := float64(mass)
	radian := a * math.Pi / 180

	const rho float64 = 1.2
	const Cd float64 = 0.47
	const A float64 = 0.05
	k := 0.5 * rho * Cd * A

	var velocity_x0, velocity_y0 float64
	var velocity_xA, velocity_yA float64
	var time_passed, displacement_x, max_height float64

	if ar {
		velocity_x := v * math.Cos(radian)
		velocity_y := v * math.Sin(radian)
		x := 0.0
		y := h
		t := 0.0

		max_height = h

		for y >= 0 {
			v_current := math.Sqrt(velocity_x*velocity_x + velocity_y*velocity_y)

			acceleration_x := -(k / m) * v_current * velocity_x
			acceleration_y := -g - (k/m)*v_current*velocity_y

			velocity_x += acceleration_x * dt
			velocity_y += acceleration_y * dt

			x += velocity_x * dt
			y += velocity_y * dt
			t += dt

			if y > max_height {
				max_height = y
			}
		}

		velocity_xA = velocity_x
		velocity_yA = velocity_y
		time_passed = t
		displacement_x = x

	} else {
		velocity_x0 = v * math.Cos(radian)
		velocity_y0 = v * math.Sin(radian)

		velocity_xA = velocity_x0
		velocity_yA = -math.Sqrt(math.Pow(velocity_y0, 2) + 2*g*h)

		time_passed = (velocity_y0 + math.Sqrt(math.Pow(velocity_y0, 2)+2*g*h)) / g

		displacement_x = time_passed * velocity_x0

		max_height = h + math.Pow(velocity_y0, 2)/(2*g)
	}

	fmt.Println(" ")
	fmt.Println("______________________________________________________________________________ ")
	fmt.Println(" ")

	fmt.Printf("Initial velocity at X direction is: %.2f m/s\n", velocity_x0)
	fmt.Printf("Initial velocity at Y direction is: %.2f m/s\n", velocity_y0)
	fmt.Println("______________________________________________________________________________ ")
	fmt.Println(" ")

	fmt.Printf("Final velocity at X direction is: %.2f m/s\n", velocity_xA)
	fmt.Printf("Final velocity at Y direction is: %.2f m/s ±0.04 percentage error \n", velocity_yA)
	fmt.Println("______________________________________________________________________________ ")
	fmt.Println(" ")

	fmt.Printf("The time passed while this motion is: %.2f s\n", time_passed)
	fmt.Println(" ")

	fmt.Printf("The Range(distance passed in x direction) is: %.2f m ±0.04 percentage error\n", displacement_x)
	fmt.Println("______________________________________________________________________________ ")
	fmt.Println(" ")

	fmt.Printf("Max height in this motion is: %.2f m\n", max_height)
	fmt.Println("______________________________________________________________________________ ")
	fmt.Println(" ")

	fmt.Print("!!! All the calculations and results are in fundamental physical units (meter, second, degree) \n")
	fmt.Println("______________________________________________________________________________ ")
	fmt.Println(" ")
	//-------------------------------------------Visualization-------------------------------------------------

	var pts plotter.XYs
	for t := 0.0; t <= time_passed; t += 0.05 {
		x := velocity_x0 * t
		y := h + velocity_y0*t - 0.5*g*t*t
		if y < 0 {
			break
		}
		pts = append(pts, plotter.XY{X: x, Y: y})
	}

	p := plot.New()
	p.Title.Text = "Projectile Motion"
	p.X.Label.Text = "Horizontal Distance (m)"
	p.Y.Label.Text = "Height (m)"

	extra := 2.0
	p.X.Min = 0
	p.X.Max = displacement_x + extra
	p.Y.Min = 0
	p.Y.Max = max_height + 2

	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}

	p.Add(line)
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "golang/PMvisualization.png"); err != nil {
		panic(err)
	}

	fmt.Print("--- Note: A graphic file is created for this motion visualization. \n")
	fmt.Println("______________________________________________________________________________ ")
	fmt.Println(" ")
}

func main() {

	var ar bool
	var height int
	var angle int
	var velocity int
	var mass int

	fmt.Print("Do you want to consider air resistance? ( true or false ): ")
	fmt.Scan(&ar)

	if ar {
		fmt.Print("Enter mass in kg: ")
		fmt.Scan(&mass)

	}

	fmt.Print("Enter velocity in m/s: ")
	fmt.Scan(&velocity)

	fmt.Print("Enter height in m: ")
	fmt.Scan(&height)

	fmt.Print("Enter angle in degrees: ")
	fmt.Scan(&angle)

	PMCalculator(height, angle, velocity, mass, ar)

}
