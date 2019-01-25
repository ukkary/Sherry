package service

import "fmt"

func getTemplate(userName string, templateName string) string {
	fmt.Println(userName, templateName)
	return "1"
}
