package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "") // Create a new PDF document
	pdf.AddPage()                          // Add a new page

	// Set font
	pdf.SetFont("Arial", "B", 16)

	// Set the invoice header
	pdf.Cell(190, 10, "Invoice")
	pdf.Ln(20)

	// Set font for the rest of the content
	pdf.SetFont("Arial", "", 12)

	// Add invoice details
	invoiceNumber := "INV-2023-001"
	invoiceDate := "August 6, 2023"
	clientName := "Shani Kumar"

	pdf.Cell(40, 10, fmt.Sprintf("Invoice Number: %s", invoiceNumber))
	pdf.Ln(7)
	pdf.Cell(40, 10, fmt.Sprintf("Invoice Date: %s", invoiceDate))
	pdf.Ln(7)
	pdf.Cell(40, 10, fmt.Sprintf("Client Name: %s", clientName))
	pdf.Ln(20)

	// Add table headers
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Item")
	pdf.Cell(40, 10, "Quantity")
	pdf.Cell(40, 10, "Price")
	pdf.Cell(40, 10, "Total")
	pdf.Ln(10)

	// Sample invoice items
	items := []struct {
		item     string
		quantity int
		price    float64
	}{
		{"Item 1", 2, 10.0},
		{"Item 2", 3, 15.0},
	}

	// Add invoice items to the table
	pdf.SetFont("Arial", "", 12)
	totalAmount := 0.0
	for _, item := range items {
		total := float64(item.quantity) * item.price
		pdf.Cell(40, 10, item.item)
		pdf.Cell(40, 10, fmt.Sprintf("%d", item.quantity))
		pdf.Cell(40, 10, fmt.Sprintf("%.2f", item.price))
		pdf.Cell(40, 10, fmt.Sprintf("%.2f", total))
		pdf.Ln(7)
		totalAmount += total
	}

	// Add total amount
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(120, 10, "Total Amount")
	pdf.Cell(40, 10, fmt.Sprintf("%.2f", totalAmount))
	pdf.Ln(20)

	// Output the PDF to a file
	err := pdf.OutputFileAndClose("invoice.pdf")
	if err != nil {
		fmt.Println("Error while generating PDF:", err)
		return
	}

	fmt.Println("PDF invoice generated successfully.")
}
