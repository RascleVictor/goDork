package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <dork_option> <search_query> [additional_operators]")
		fmt.Println("Dork options:")
		fmt.Println("  1. Search in titles (intitle)")
		fmt.Println("  2. Search in URLs (inurl)")
		fmt.Println("  3. Search in a specific site (site)")
		fmt.Println("  4. Search for a specific file type (filetype)")
		fmt.Println("  5. Search for specific text in the page (intext)")
		fmt.Println("  6. Combine multiple options")
		fmt.Println("  7. Search for known vulnerabilities")
		fmt.Println("  8. Custom targeted search (site:[TARGET] inurl:_cpanel/forgotpwd, etc.)")
		fmt.Println("  9. Advanced Google Dorking (inurl:\"/admin/login\" intitle:\"login\")")
		fmt.Println(" 10. Search for Config files")
		fmt.Println(" 11. Search for Database files")
		fmt.Println(" 12. Search for Backup files")
		fmt.Println(" 13. Search for .git folder")
		fmt.Println(" 14. Search for Exposed documents")
		fmt.Println(" 15. Search for SQL errors")
		fmt.Println(" 16. Search for PHP errors")
		fmt.Println(" 17. Search for Login pages")
		fmt.Println(" 18. Search for Open redirects")
		fmt.Println(" 19. Search for Apache Struts RCE")
		fmt.Println(" 20. Search for Wordpress files")
		fmt.Println(" 21. Search for Other files")
		fmt.Println(" 22. Search for Linkedin employees")
		fmt.Println(" 23. Search for AWS S3 Buckets")
		fmt.Println(" 24. Search for Azure")
		fmt.Println(" 25. Search for Google Cloud")
		fmt.Println("Additional operators (optional):")
		fmt.Println("  - Wildcard Operator (*): go run main.go 1 'search query' *")
		fmt.Println("  - Filetype Operator: go run main.go 4 'filetype:pdf search'")
		os.Exit(1)
	}

	var query string

	switch os.Args[1] {
	case "1":
		query = fmt.Sprintf("intitle:%s", strings.Join(os.Args[2:], " "))
	case "2":
		query = fmt.Sprintf("inurl:%s", strings.Join(os.Args[2:], " "))
	case "3":
		query = fmt.Sprintf("site:%s", strings.Join(os.Args[2:], " "))
	case "4":
		query = fmt.Sprintf("filetype:%s", strings.Join(os.Args[2:], " "))
	case "5":
		query = fmt.Sprintf("intext:%s", strings.Join(os.Args[2:], " "))
	case "6":
		query = strings.Join(os.Args[2:], " ")
	case "7":
		query = "known vulnerabilities"
	case "8":
		// Custom targeted search
		if len(os.Args) < 5 {
			fmt.Println("Custom targeted search requires a target and at least one search term.")
			os.Exit(1)
		}
		query = fmt.Sprintf("site:%s %s", os.Args[3], strings.Join(os.Args[4:], " "))
	case "9":
		// Advanced Google Dorking
		if len(os.Args) < 4 {
			fmt.Println("Advanced Google Dorking requires at least one search term.")
			os.Exit(1)
		}
		query = strings.Join(os.Args[2:], " ")
	case "10":
		query = "filetype:ini OR filetype:env OR filetype:config"
	case "11":
		query = "filetype:sql OR filetype:db"
	case "12":
		query = "filetype:bkf OR filetype:bkp OR filetype:bak OR filetype:old OR filetype:backup"
	case "13":
		query = "inurl:.git"
	case "14":
		query = "filetype:doc OR filetype:docx OR filetype:ppt OR filetype:pptx OR filetype:xls OR filetype:xlsx OR filetype:pdf"
	case "15":
		query = "intext:\"SQL syntax error\" OR intext:\"Warning: mysql_connect()\" OR intext:\"Warning: mysqli_connect()\""
	case "16":
		query = "intext:\"Parse error\" OR intext:\"Fatal error\" OR intext:\"Warning: include\" OR intext:\"Warning: require\""
	case "17":
		query = "inurl:login OR inurl:signin OR intitle:login"
	case "18":
		query = "inurl:redir OR inurl:redirect"
	case "19":
		query = "intitle:\"Apache Struts Default User Interface\""
	case "20":
		query = "filetype:wpd OR filetype:wps OR filetype:wpa OR filetype:wpb OR filetype:wpf OR filetype:wpg OR filetype:wpp OR filetype:wpt OR filetype:wpw"
	case "21":
		query = "filetype:log OR filetype:lst OR filetype:pwd OR filetype:sql OR filetype:conf OR filetype:inc OR filetype:ini OR filetype:bkf OR filetype:bkp OR filetype:bak OR filetype:old OR filetype:backup"
	case "22":
		query = "site:linkedin.com employee"
	case "23":
		query = "site:s3.amazonaws.com"
	case "24":
		query = "site:azure.com"
	case "25":
		query = "site:cloud.google.com"
	default:
		fmt.Println("Invalid dork option. Please choose a valid option.")
		os.Exit(1)
	}

	// Check if there are additional operators
	if len(os.Args) > 3 {
		query += " " + strings.Join(os.Args[3:], " ")
	}

	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(query))

	doc, err := goquery.NewDocument(searchURL)
	if err != nil {
		fmt.Println("Error fetching search results:", err)
		os.Exit(1)
	}

	fmt.Println("Google Dorking Results for:", query)

	file, err := os.Create("results.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists && strings.HasPrefix(link, "/url?q=") {
			link = strings.TrimPrefix(link, "/url?q=")
			result := fmt.Sprintf("%d. %s\n", i+1, link)
			fmt.Print(result)
			file.WriteString(result)
		}
	})

	fmt.Println("Results saved to results.txt")
}
