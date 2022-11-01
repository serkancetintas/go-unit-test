package main

import "fmt"

type Computer struct {
	Pr Printer
}

func (c *Computer) PrintA4(content string) error {
	c.Pr.Print("A4", content)
	return nil
}

type MyPrinter struct{}

func (e *MyPrinter) Print(paperSize string, content string) error {
	fmt.Println("Printing:", content, ", with paper size:", paperSize)
	return nil
}

type Printer interface {
	Print(paperSize string, content string) error
}
