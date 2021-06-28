package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type TrieNode struct {
	Children map[rune]*TrieNode
	value    string
}

func InsertTrieNode(node *TrieNode, value string) {
	for _, c := range value {
		if _, ok := node.Children[c]; !ok {
			node.Children[c] = &TrieNode{
				Children: map[rune]*TrieNode{},
			}
		}
		node = node.Children[c]
	}
	node.value = value
}
func FindTrieNode(node *TrieNode, key string) string {
	for _, c := range key {
		if _, ok := node.Children[c]; ok {
			node = node.Children[c]
		} else {
			return ""
		}
	}
	return node.value
}

func GetTrieNode(node *TrieNode, key string) *TrieNode {
	for _, c := range key {
		if _, ok := node.Children[c]; ok {
			node = node.Children[c]
		} else {
			return nil
		}
	}
	return node
}

func FindForPrefix(node *TrieNode, prefix string) []string {
	results := make([]string, 0)
	startNode := GetTrieNode(node, prefix)
	Collect(startNode, &results)
	return results
}

func Collect(node *TrieNode, results *[]string) {
	if node == nil {
		return
	}
	if node.value != "" {
		*results = append(*results, node.value)
	}
	for c := range node.Children {
		Collect(node.Children[c], results)
	}

}

func main() {
	node := &TrieNode{
		Children: map[rune]*TrieNode{},
		value:    "",
	}

	f, err := os.Open("words_alpha.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		InsertTrieNode(node, scanner.Text())
		i += 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d words available \n", i)

	fmt.Println("enter a seach query for autocomplete")
	reader := bufio.NewReader(os.Stdin)
	query, _ := reader.ReadString('\n')
	query = query[:len(query)-2]
	results := FindForPrefix(node, query)
	fmt.Println("here are the results!")
	fmt.Print(results)
}
