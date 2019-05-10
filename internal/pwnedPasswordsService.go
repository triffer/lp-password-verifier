package internal

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var cachedPwnedRequests  = make(map[string][]string)

func FindPwnedPasswordsByHash(passwordHash string) []string {
	passwordPrefix := string(passwordHash[0:5])

	pwnedPasswords, ok := cachedPwnedRequests[passwordPrefix]

	if !ok {
		pwnedPasswords = getPwnedPasswordsFromApi(passwordPrefix)
	}

	return pwnedPasswords
}

func getPwnedPasswordsFromApi(passwordPrefix string) []string {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get("https://api.pwnedpasswords.com/range/" + passwordPrefix)
	CheckError(err)

	defer func() {
		err := resp.Body.Close()
		CheckError(err)
	}()

	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)

	result := string(body)

	var pwnedPasswords []string

	scanner := bufio.NewScanner(strings.NewReader(result))

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		pwnedPasswordPostfix := s[0]
		pwnedPasswords = append(pwnedPasswords, passwordPrefix+pwnedPasswordPostfix)
	}

	cachedPwnedRequests[passwordPrefix] = pwnedPasswords

	return pwnedPasswords
}