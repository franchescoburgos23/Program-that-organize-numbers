package main

import "fmt"

func regards() {
	fmt.Println("  *********************************************************  ")
	fmt.Println("Hi Im a Program that organizes the numbers from less to higher")
	fmt.Println("  *********************************************************  ")
}

func goodbye() {
	fmt.Println("  *********************************************************  ")
	fmt.Println("               Thank you for used me goodbye :D  ")
	fmt.Println("  *********************************************************  ")
}

func main() {

	regards()

	arr := [4]int{}

	fmt.Println("  Please enter your 4 Numbers ")

	fmt.Scan(&arr[0])
	fmt.Scan(&arr[1])
	fmt.Scan(&arr[2])
	fmt.Scan(&arr[3])

	slic := arr[0:3]

	fmt.Println("  *********************************************************  ")

	fmt.Println("Do you want Print 3 or 4 Numbers")

	var answer int

	fmt.Scan(&answer)
	if answer == 3 {

		if slic[0] < slic[2] && slic[2] < slic[1] {

			fmt.Println("*********************")

			fmt.Println(slic[0], slic[2], slic[1])

		} else if slic[2] < slic[0] && slic[0] < slic[1] {

			fmt.Println("*********************")

			fmt.Println(slic[2], slic[0], slic[1])

		} else if slic[1] < slic[2] && slic[2] < slic[0] {

			fmt.Println("*********************")

			fmt.Println(slic[1], slic[2], slic[0])

		} else if slic[2] < slic[1] && slic[1] < slic[0] {

			fmt.Println("*********************")

			fmt.Println(slic[2], slic[1], slic[0])

		} else if slic[0] < slic[1] && slic[1] < slic[2] {

			fmt.Println("*********************")

			fmt.Println(slic[0], slic[1], slic[2])

		} else if slic[1] < slic[0] && slic[0] < slic[2] {

			fmt.Println("*********************")

			fmt.Println(slic[1], slic[0], slic[2])

		}

	} else if answer == 4 {

		if arr[0] < arr[2] && arr[2] < arr[3] && arr[3] < arr[1] {

			fmt.Println("*********************")

			fmt.Println(arr[0], arr[2], arr[3], arr[1])

		} else if arr[0] < arr[3] && arr[3] < arr[2] && arr[2] < arr[1] {

			fmt.Println("*********************")

			fmt.Println(arr[0], arr[3], arr[2], arr[1])

		} else if arr[2] < arr[0] && arr[0] < arr[3] && arr[3] < arr[1] {

			fmt.Println("*********************")

			fmt.Println(arr[2], arr[0], arr[3], arr[1])

		} else if arr[2] < arr[3] && arr[3] < arr[0] && arr[3] < arr[1] {

			fmt.Println("*********************")

			fmt.Println(arr[2], arr[0], arr[3], arr[1])

		} else if arr[3] < arr[2] && arr[2] < arr[0] && arr[0] < arr[1] {

			fmt.Println("*********************")

			fmt.Println(arr[3], arr[2], arr[0], arr[1])

		} else if arr[3] < arr[0] && arr[0] < arr[2] && arr[2] < arr[1] {

			fmt.Println("*********************")

			fmt.Println(arr[3], arr[0], arr[2], arr[1])

		}

		if arr[1] < arr[2] && arr[2] < arr[3] && arr[3] < arr[0] {

			fmt.Println("*********************")

			fmt.Println(arr[1], arr[2], arr[3], arr[0])

		} else if arr[1] < arr[3] && arr[3] < arr[2] && arr[2] < arr[0] {

			fmt.Println("*********************")

			fmt.Println(arr[1], arr[3], arr[2], arr[0])

		} else if arr[2] < arr[1] && arr[1] < arr[3] && arr[3] < arr[0] {

			fmt.Println("*********************")

			fmt.Println(arr[2], arr[1], arr[3], arr[0])

		} else if arr[2] < arr[3] && arr[3] < arr[1] && arr[3] < arr[0] {

			fmt.Println("*********************")

			fmt.Println(arr[2], arr[1], arr[3], arr[0])

		} else if arr[3] < arr[2] && arr[2] < arr[1] && arr[1] < arr[0] {

			fmt.Println("*********************")

			fmt.Println(arr[3], arr[2], arr[1], arr[0])

		} else if arr[3] < arr[1] && arr[1] < arr[2] && arr[2] < arr[0] {

			fmt.Println("*********************")

			fmt.Println(arr[3], arr[1], arr[2], arr[0])
		}

		if arr[0] < arr[1] && arr[1] < arr[3] && arr[3] < arr[2] {

			fmt.Println("*********************")

			fmt.Println(arr[0], arr[1], arr[3], arr[2])

		} else if arr[0] < arr[3] && arr[3] < arr[1] && arr[1] < arr[2] {

			fmt.Println("*********************")

			fmt.Println(arr[0], arr[3], arr[1], arr[2])

		} else if arr[1] < arr[0] && arr[0] < arr[3] && arr[3] < arr[2] {

			fmt.Println("*********************")

			fmt.Println(arr[1], arr[0], arr[3], arr[2])

		} else if arr[1] < arr[3] && arr[3] < arr[0] && arr[0] < arr[2] {

			fmt.Println("*********************")

			fmt.Println(arr[1], arr[3], arr[0], arr[2])

		} else if arr[3] < arr[0] && arr[0] < arr[1] && arr[1] < arr[2] {

			fmt.Println("*********************")

			fmt.Println(arr[3], arr[0], arr[1], arr[2])

		} else if arr[3] < arr[1] && arr[1] < arr[0] && arr[0] < arr[2] {

			fmt.Println("*********************")

			fmt.Println(arr[3], arr[1], arr[0], arr[2])
		}

		if arr[0] < arr[1] && arr[1] < arr[2] && arr[2] < arr[3] {

			fmt.Println("*********************")

			fmt.Println(arr[0], arr[1], arr[2], arr[3])

		} else if arr[0] < arr[2] && arr[2] < arr[1] && arr[1] < arr[3] {

			fmt.Println("*********************")

			fmt.Println(arr[0], arr[2], arr[1], arr[3])

		} else if arr[1] < arr[0] && arr[0] < arr[2] && arr[2] < arr[3] {

			fmt.Println("*********************")

			fmt.Println(arr[1], arr[0], arr[2], arr[3])

		} else if arr[1] < arr[2] && arr[2] < arr[0] && arr[0] < arr[3] {

			fmt.Println("*********************")

			fmt.Println(arr[1], arr[2], arr[0], arr[3])

		} else if arr[2] < arr[0] && arr[0] < arr[1] && arr[1] < arr[3] {

			fmt.Println("*********************")

			fmt.Println(arr[2], arr[0], arr[1], arr[3])

		} else if arr[2] < arr[1] && arr[1] < arr[0] && arr[0] < arr[3] {

			fmt.Println("*********************")

			fmt.Println(arr[2], arr[1], arr[0], arr[3])
		}
	} else {
		fmt.Println("Error please select a number between 3 and 4")
	}

	goodbye()

}
