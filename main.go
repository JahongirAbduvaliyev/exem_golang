package main

import (
	"errors"
	"fmt"
)


type Vehicle interface {
	Rent() error
	Return() error
	Details() error
}


type Car struct {
	ID     int
	Brand  string
	Rented bool
}


func (c *Car) Rent() error {
	if c.Rented {
		return errors.New("car ijaraga olingan")
	}
	c.Rented = true
	return nil
}

func (c *Car) Return() error {
	if c.Rented {
		return errors.New("car ijaraga olinmagan")
	}
	c.Rented = false
	return nil
}

func (c *Car) Details() error {
	fmt.Printf("Car - ID: %d, Brand: %s, Rented: %t\n", c.ID, c.Brand, c.Rented)
	return nil
}


type Bike struct {
	ID     int
	Brand  string
	Rented bool
}

func (b *Bike) Rent() error {
	if !b.Rented {
		return errors.New("bike ijaraga olingan")
	}
	b.Rented = true
	return nil
}

func (b *Bike) Return() error {
	if !b.Rented {
		return errors.New("bike ijaraga olinmagan")
	}
	b.Rented = false
	return nil
}

func (b *Bike) Details() error {
	fmt.Printf("Bike - ID: %d, Brand: %s, Rented: %t\n", b.ID, b.Brand, b.Rented)
	return nil
}


type Skate struct {
	ID     int
	Brand  string
	Rented bool
}

func (s *Skate) Rent() error {
	if s.Rented {
		return errors.New("skate ijaraga olingan")
	}
	s.Rented = true
	return nil
}

func (s *Skate) Return() error {
	if !s.Rented {
		return errors.New("skate ijaraga olinmagan")
	}
	s.Rented = false
	return nil
}

func (s *Skate) Details() error {
	fmt.Printf("Skate - ID: %d, Brand: %s, Rented: %t\n", s.ID, s.Brand, s.Rented)
	return nil
}

func main() {
	var vehicles []Vehicle

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add a Vehicle")
		fmt.Println("2. Rent a Vehicle")
		fmt.Println("3. Return a Vehicle")
		fmt.Println("4. Show Vehicle Details")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice:")
		fmt.Scanln(&choice)

		switch choice {
		case 1: 
			var vehicleType string
			fmt.Print("Enter vehicle type (car/bike/skate): ")
			fmt.Scanln(&vehicleType)

			var id int
			var brand string
			fmt.Print("Enter ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Brand: ")
			fmt.Scanln(&brand)

			switch vehicleType {
			case "car":
				vehicles = append(vehicles, &Car{ID: id, Brand: brand})
				fmt.Println("Car added ")
			case "bike":
				vehicles = append(vehicles, &Bike{ID: id, Brand: brand})
				fmt.Println("Bike added ")
			case "skate":
				vehicles = append(vehicles, &Skate{ID: id, Brand: brand})
				fmt.Println("Skate added ")
			default:
				fmt.Println("no vehicle")
			}

		case 2: 
			var id int
			fmt.Print("Enter the ID of the vehicle to rent: ")
			fmt.Scanln(&id)

			found := false
			for _, v := range vehicles {
				switch t := v.(type) {
				case *Car:
					if t.ID == id {
						found = true
						if err := t.Rent(); err != nil {
							fmt.Println(err)
						} else {
							fmt.Println("Car rented successfully!")
						}
					}
				case *Bike:
					if t.ID == id {
						found = true
						if err := t.Rent(); err != nil {
							fmt.Println(err)
						} else {
							fmt.Println("Bike rented successfully!")
						}
					}
				case *Skate:
					if t.ID == id {
						found = true
						if err := t.Rent(); err != nil {
							fmt.Println(err)
						} else {
							fmt.Println("Skate rented successfully!")
						}
					}
				}
			}
			if !found {
				fmt.Println("Vehicle not found!")
			}

		case 3: 
			var id int
			fmt.Print("Enter the ID of the vehicle to return: ")
			fmt.Scanln(&id)

			found := false
			for _, v := range vehicles {
				switch t := v.(type) {
				case *Car:
					if t.ID == id {
						found = true
						if err := t.Return(); err != nil {
							fmt.Println(err)
						} else {
							fmt.Println("Car returned successfully!")
						}
					}
				case *Bike:
					if t.ID == id {
						found = true
						if err := t.Return(); err != nil {
							fmt.Println(err)
						} else {
							fmt.Println("Bike returned successfully!")
						}
					}
				case *Skate:
					if t.ID == id {
						found = true
						if err := t.Return(); err != nil {
							fmt.Println(err)
						} else {
							fmt.Println("Skate returned successfully!")
						}
					}
				}
			}
			if !found {
				fmt.Println("Vehicle not found!")
			}

		case 4: 
			fmt.Println("\nVehicle Details:")
			for _, v := range vehicles {
				v.Details()
			}

		case 5: 
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}
