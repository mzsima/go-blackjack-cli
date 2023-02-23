package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Rank  string
	Value int
}

var deck []Card

func main() {
	rand.Seed(time.Now().UnixNano())

	// カードを初期化する
	deck = newDeck()
	shuffleDeck()

	// ゲームを開始する
	fmt.Println("ブラックジャックへようこそ！")
	fmt.Println("ディーラーがカードを配ります。")

	playerCards := []Card{}
	dealerCards := []Card{}

	// プレイヤーがカードを引く
	fmt.Println("あなたのカード:")
	playerCards = append(playerCards, drawCard())
	playerCards = append(playerCards, drawCard())
	showCards(playerCards)

	// ディーラーがカードを引く
	fmt.Println("ディーラーのカード:")
	dealerCards = append(dealerCards, drawCard())
	dealerCards = append(dealerCards, drawCard())
	showCards(dealerCards[:1])
	fmt.Println("[?]")

	// プレイヤーのターン
	for {
		fmt.Println("カードを引きますか？(y/n)")
		var choice string
		fmt.Scan(&choice)
		if choice == "y" {
			playerCards = append(playerCards, drawCard())
			showCards(playerCards)
			if getCardSum(playerCards) > 21 {
				fmt.Println("バストしました。あなたの負けです！")
				return
			}
		} else if choice == "n" {
			break
		} else {
			fmt.Println("yかnを入力してください。")
		}
	}

	// ディーラーのターン
	for getCardSum(dealerCards) < 17 {
		dealerCards = append(dealerCards, drawCard())
		fmt.Println("ディーラーがカードを引きました。")
	}

	// 結果を表示する
	fmt.Println("あなたのカード:")
	showCards(playerCards)
	fmt.Println("ディーラーのカード:")
	showCards(dealerCards)
	playerSum := getCardSum(playerCards)
	dealerSum := getCardSum(dealerCards)
	if playerSum > 21 {
		fmt.Println("バストしました。あなたの負けです！")
	} else if dealerSum > 21 {
		fmt.Println("ディーラーがバストしました。あなたの勝ちです！")
	} else if playerSum > dealerSum {
		fmt.Println("あなたの勝ちです！")
	} else if playerSum == dealerSum {
		fmt.Println("引き分けです！")
	} else {
		fmt.Println("あなたの負けです！")
	}
}

func newDeck() []Card {
	var deck []Card
	suits := []string{"♠️", "♦️", "♣️", "♥️"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	values := []int{11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}
	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(ranks); j++ {
			card := Card{Suit: suits[i], Rank: ranks[j], Value: values[j]}
			deck = append(deck, card)
		}
	}
	return deck
}

func shuffleDeck() {
	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func drawCard() Card {
	card := deck[0]
	deck = deck[1:]
	return card
}

func showCards(cards []Card) {
	for _, card := range cards {
		fmt.Printf("[%s of %s] ", card.Rank, card.Suit)
	}
	fmt.Println("")
}

func getCardSum(cards []Card) int {
	sum := 0
	for _, card := range cards {
		sum += card.Value
	}
	for _, card := range cards {
		if card.Value == 11 && sum > 21 {
			sum -= 10
		}
	}
	return sum
}
