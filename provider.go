package main

import (
	"fmt"
	"log"
	"regexp"
	"encoding/json"

	curl "github.com/andelf/go-curl"
)

func GetCurrencies() (map[string]string, error) {
	result, err := LoadDataFromCache("currencies")
	if err != nil {
		return nil, err
	}

	if result != nil {
		return result, nil
	}

	urlPage := "https://bonbast.com/"
	urlJson := "https://bonbast.com/json"

	html := loadHtmlPage(urlPage)
	key, err := getPostJsonKeyFromHtml(html)
	if err != nil {
		log.Panicln(err)
	}

	jsonData := loadJsonData(urlJson, key)

	_, err = CacheData("currencies", jsonData, 180)
	if err != nil {
		return result, err
	}

	json.Unmarshal([]byte(jsonData), &result)
	return result, nil
}

func loadHtmlPage(url string) string {
	var result string
	ret := curl.EasyInit()
	defer ret.Cleanup()

	fn := func(buf []byte, userdata any) bool {
		result = string(buf)
		return true
	}

	ret.Setopt(curl.OPT_URL, url)
	ret.Setopt(curl.OPT_TRANSFERTEXT, 1)
	ret.Setopt(curl.OPT_SSL_VERIFYPEER, false)
	ret.Setopt(curl.OPT_SSL_VERIFYHOST, false)
	ret.Setopt(curl.OPT_FOLLOWLOCATION, true)
	ret.Setopt(curl.OPT_USERAGENT, "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	ret.Setopt(curl.OPT_COOKIEJAR, "cookie.txt")
	ret.Setopt(curl.OPT_COOKIEFILE, "cookie.txt")
	ret.Setopt(curl.OPT_REFERER, "https://bonbast.com")
	ret.Setopt(curl.OPT_ENCODING, "gzip, deflate, br")
	ret.Setopt(curl.OPT_HTTPHEADER, []string{
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Accept-Encoding: gzip, deflate, br",
		"Accept-Language: en-US,en;q=0.9",
		"Cache-Control: max-age=0",
		"Connection: keep-alive",
		"Host: bonbast.com",
		"Upgrade-Insecure-Requests: 1",
		"User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
	})
	ret.Setopt(curl.OPT_WRITEFUNCTION, fn)

	ret.Perform()

	return result
}

func loadJsonData(url string, key string) string {
	var result string
	ret := curl.EasyInit()
	defer ret.Cleanup()

	fn := func(buf []byte, userdata any) bool {
		result = string(buf)
		return true
	}

	ret.Setopt(curl.OPT_URL, url)
	ret.Setopt(curl.OPT_TRANSFERTEXT, 1)
	ret.Setopt(curl.OPT_SSL_VERIFYPEER, false)
	ret.Setopt(curl.OPT_SSL_VERIFYHOST, false)
	ret.Setopt(curl.OPT_FOLLOWLOCATION, true)
	ret.Setopt(curl.OPT_USERAGENT, "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	ret.Setopt(curl.OPT_COOKIEJAR, "cookie.txt")
	ret.Setopt(curl.OPT_COOKIEFILE, "cookie.txt")
	ret.Setopt(curl.OPT_REFERER, "https://bonbast.com")
	ret.Setopt(curl.OPT_ENCODING, "gzip, deflate, br")
	ret.Setopt(curl.OPT_HTTPHEADER, []string{"Accept: application/json, text/javascript, */*; q=0.01",
		"Accept-Encoding: gzip, deflate, br",
		"Accept-Language: en-US,en;q=0.9,fa;q=0.8,it;q=0.7",
		"Cache-Control: max-age=0",
		"Connection: keep-alive",
		"Content-Type: application/x-www-form-urlencoded; charset=UTF-8",
		"Host: bonbast.com",
		"Origin: https://bonbast.com",
		"Referer: https://bonbast.com/",
		"Sec-Ch-Ua: \"Chromium\";v=\"104\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"104\"",
		"Sec-Ch-Ua-Mobile: ?0",
		"Sec-Ch-Ua-Platform: \"Windows\"",
		"Sec-Fetch-Dest: empty",
		"Sec-Fetch-Mode: cors",
		"Sec-Fetch-Site: same-origin",
		"User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		"X-Requested-With: XMLHttpRequest",
	})
	ret.Setopt(curl.OPT_POST, 1)
	ret.Setopt(curl.OPT_POSTFIELDS, fmt.Sprintf("data=%s&webdriver=false", key))
	ret.Setopt(curl.OPT_WRITEFUNCTION, fn)

	ret.Perform()

	return result
}

func getPostJsonKeyFromHtml(jd string) (string, error) {
	r, err := regexp.Compile(`\$.post\(\'\/json\',{ data:\"(?P<data>.*)\"`)
	if err != nil {
		return "", err
	}

	key := r.FindStringSubmatch(jd)
	result := key[1]

	return result, nil
}
