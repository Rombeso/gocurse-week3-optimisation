package main

import (
	"bufio"
	"fmt"
	"hw3/data"
	"io"
	"os"
	"regexp"
	"strings"
)

var r = regexp.MustCompile("@")

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	/*
		!!! !!! !!!
		обратите внимание - в задании обязательно нужен отчет
		делать его лучше в самом начале, когда вы видите уже узкие места, но еще не оптимизировалм их
		так же обратите внимание на команду в параметром -http
		перечитайте еще раз задание
		!!! !!! !!!
	*/
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	seenBrowsers := []string{}

	uniqueBrowsers := 0
	//var foundUsers bytes.Buffer
	users := make([]data.User, len(lines))

	for j, line := range lines {
		var u = data.User{}
		// fmt.Printf("%v %v\n", err, line)
		err := u.UnmarshalJSON([]byte(line))
		if err != nil {
			panic(err)
		}
		users[j] = u
	}
	fmt.Fprintf(out, "found users:\n")
	for i, user := range users {

		isAndroid := false
		isMSIE := false

		if len(user.Browsers) == 0 {
			//log.Println("cant cast browsers")
			continue
		}

		for _, browserRaw := range user.Browsers {
			if browserRaw == "" {
				//log.Println("cant cast browser to string")
				continue
			}
			if ok := strings.Contains(browserRaw, "Android"); ok {
				//log.Println("browser patAndroid: ", browserRaw)
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browserRaw {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					//log.Printf("SLOW New browser: %s, first seen: %s", browserRaw, user.Name)
					seenBrowsers = append(seenBrowsers, browserRaw)
					uniqueBrowsers++

				}
			}
			if ok := strings.Contains(browserRaw, "MSIE"); ok {
				//log.Println("browser patMSIE: ", browser)
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browserRaw {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					//log.Printf("SLOW New browser: %s, first seen: %s", browserRaw, user.Name)
					seenBrowsers = append(seenBrowsers, browserRaw)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		//log.Println("Android and MSIE user:", user.Name, user.Email)

		email := r.ReplaceAllString(user.Email, " [at] ")
		fmt.Fprintf(out, "[%d] %s <%s>\n", i, user.Name, email)
		//foundUsers.WriteString(fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email))

	}
	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
	err = file.Close()
	if err != nil {
		panic(err)
	}
	//SlowSearch(out)
}
