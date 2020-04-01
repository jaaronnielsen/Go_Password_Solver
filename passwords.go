/*I learned how to write to a file on https://golangcode.com/writing-to-file/
I leanred how to read from a file, and how to do SHA1 hashing from https://gobyexample.com/
I learned how to change an array of ints to an array of strings from https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string/37533144*/

package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	for i := 1; i > 0; {
		newPassword := "hello"
		j := 0
		fmt.Println("Enter 1 to enter a new password or 0 to quit")
		fmt.Scan(&j)

		if j == 1 {
			fmt.Println("Enter a password that contains no spaces and I will see if I can guess it. ")
			fmt.Scan(&newPassword)

			checkRegex(newPassword)

			brute(newPassword)

			checkFile(newPassword)

		} else if j == 0 {
			os.Exit(3)
		}

	}

}

func hashToFile(newPassword string, filename string) {
	h := sha1.New()
	h.Write([]byte(newPassword))
	bs := h.Sum(nil)
	hashed := fmt.Sprintf("%x\n", bs)
	err := WriteToFile(filename, hashed, newPassword)
	check(err)

}

func hashWord(newPassword string) (hashed string) {
	h := sha1.New()
	h.Write([]byte(newPassword))
	bs := h.Sum(nil)
	hashed = fmt.Sprintf("%x\n", bs)
	return hashed
}

func checkRegex(newPassword string) {
	r, _ := regexp.Compile("^(?=.*0-9)(?=.*[a-z])(?=.*[A-Z]).{4,8}$")
	if r.MatchString(newPassword) {
		fmt.Println("A strong password was found in the file ")
		err := WriteToFile("regex.txt", newPassword, " ")
		check(err)
	} else {
		fmt.Println("no strong passwords were found in the file")
	}

}

func checkFile(newPassword string) {
	dat, err := ioutil.ReadFile("passwords.txt")
	check(err)
	word := string(dat)
	hashed := hashWord(newPassword)

	if strings.Contains(word, hashed) == false {
		fmt.Println("The password was not found. ")
		hashToFile(newPassword, "passwords.txt")

	} else {
		fmt.Println("The password was found using stored passwords:", newPassword)

	}

}

func brute(newPassword string) {
	pass := []int{0, 0, 0, 0}

	for i := 0; i < 10; i += 1 {
		pass[0] = i
		for j := 0; j < 10; j += 1 {
			pass[1] = j
			for k := 0; k < 10; k += 1 {
				pass[2] = k
				for l := 0; l < 10; l += 1 {
					pass[3] = l
					nPass := arrayToString(pass, "")
					if nPass == newPassword {
						fmt.Println("The password was found using brute force: ", nPass)
						main()
					}

				}
			}
		}
	}
}

func arrayToString(a []int, delim string) string {
	//trim takes off leading or trailing characters
	//replaces the strings you don't want with strings that you do want
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
func WriteToFile(filename string, data string, password string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()

	_, err = io.WriteString(file, password+": ")
	check(err)

	_, err = io.WriteString(file, data)
	check(err)
	return file.Sync()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
