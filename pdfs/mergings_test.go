package pdfs

import "testing"

func TestServiceMergePdf_Run(t *testing.T) {
	fileList := []string{"tax.pdf", "account-confirmation.pdf", "bank-statement.pdf", "proff-address.pdf"}
	outputFile := "out.pdf"
	hub := ServiceMergePdf{
		FileList:   fileList,
		OutputFile: outputFile,
	}
	hub.Run()
}
