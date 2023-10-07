package main

import "fmt"

func main() {
	const premiumPlanName = "Premium Plan"
	// error
	// premiumPlanName = "Basic Plan"

	const basicPlanName = "Basic Plan"

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}
