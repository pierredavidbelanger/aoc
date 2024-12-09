package day5

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	X, Y int
}

type Update struct {
	Pages []int
}

func parse(file string) ([]Rule, []Update) {
	rules := make([]Rule, 0)
	updates := make([]Update, 0)
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(bytes.NewReader(data))
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		lineParts := strings.Split(line, "|")
		rule := Rule{}
		rule.X, err = strconv.Atoi(lineParts[0])
		if err != nil {
			log.Fatal(err)
		}
		rule.Y, err = strconv.Atoi(lineParts[1])
		if err != nil {
			log.Fatal(err)
		}
		rules = append(rules, rule)
	}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		lineParts := strings.Split(line, ",")
		update := Update{}
		update.Pages = make([]int, 0)
		for _, linePart := range lineParts {
			page, err := strconv.Atoi(linePart)
			if err != nil {
				log.Fatal(err)
			}
			update.Pages = append(update.Pages, page)
		}
		updates = append(updates, update)
	}
	return rules, updates
}

func findCorrectlyOrderedUpdates(rules []Rule, updates []Update) ([]Update, []Update) {
	correctUpdates, incorrectUpdates := make([]Update, 0), make([]Update, 0)
	for _, update := range updates {
		if isUpdateCorrectlyOrdered(rules, update) {
			correctUpdates = append(correctUpdates, update)
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}
	return correctUpdates, incorrectUpdates
}

func isUpdateCorrectlyOrdered(rules []Rule, update Update) bool {
	for pageIndex, page := range update.Pages {
		rulesForPage := findRulesForPage(rules, page)
		if !isPageCorrectlyOrdered(rulesForPage, update, pageIndex) {
			return false
		}
	}
	return true
}

func isPageCorrectlyOrdered(rulesForPage []Rule, update Update, pageIndex int) bool {
	for i := 0; i < pageIndex; i++ {
		for _, rule := range rulesForPage {
			if update.Pages[i] == rule.Y {
				return false
			}
		}
	}
	return true
}

func findRulesForPage(rules []Rule, page int) []Rule {
	rulesForPage := make([]Rule, 0)
	for _, rule := range rules {
		if rule.X == page {
			rulesForPage = append(rulesForPage, rule)
		}
	}
	return rulesForPage
}

func sumMiddlePages(updates []Update) int {
	sum := 0
	for _, update := range updates {
		pageIndex := int(math.Floor(float64(len(update.Pages)) / 2))
		sum += update.Pages[pageIndex]
	}
	return sum
}

func fixUpdates(rules []Rule, updates *[]Update) {
	for updateIndex, _ := range *updates {
		fixUpdate(rules, &(*updates)[updateIndex])
	}
}

func fixUpdate(rules []Rule, update *Update) {
	for {
		for pageIndex, page := range update.Pages {
			rulesForPage := findRulesForPage(rules, page)
			if !isPageCorrectlyOrdered(rulesForPage, *update, pageIndex) {
				fixPage(rulesForPage, update, pageIndex)
			}
		}
		if isUpdateCorrectlyOrdered(rules, *update) {
			break
		}
	}
}

func fixPage(rulesForPage []Rule, update *Update, pageIndex int) {
	for _, rule := range rulesForPage {
		yPageIndex := -1
		for i := 0; i < pageIndex; i++ {
			if update.Pages[i] == rule.Y {
				yPageIndex = i
				break
			}
		}
		if yPageIndex > -1 {
			for i := pageIndex; i > yPageIndex; i-- {
				update.Pages[i-1], update.Pages[i] = update.Pages[i], update.Pages[i-1]
			}
			pageIndex = yPageIndex
		}
	}
}
