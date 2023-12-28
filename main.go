package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Player struct {
	Name               string
	TeamID             int
	TwoPointAttempts   int
	TwoPointMade       int
	ThreePointAttempts int
	ThreePointMade     int
}

type Game struct {
	Team1  []Player
	Team2  []Player
	Score1 int
	Score2 int
}

type TeamStats struct {
	Team    string
	Wins    int
	Losses  int
	Players []Player
}

var teamStats = make(map[string]*TeamStats)
var mutex = &sync.Mutex{}

func simulateMatch(matchID int, match Game, team1Name string, team2Name string, updates chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for minute := 0; minute < 48; minute++ {
		attackCountTeam1 := rand.Intn(2) + 1 // 0, 1 veya 2 atak
		attackCountTeam2 := rand.Intn(2) + 1 // 0, 1 veya 2 atak

		score1, updatedTeam1 := simulateTeamAttack(match.Team1, attackCountTeam1)
		score2, updatedTeam2 := simulateTeamAttack(match.Team2, attackCountTeam2)

		// Skorları ve takımları güncelle
		mutex.Lock()
		teamStats[team1Name].Players = updatedTeam1
		teamStats[team2Name].Players = updatedTeam2
		match.Score1 += score1
		match.Score2 += score2
		mutex.Unlock()

		// Skoru güncelle
		fmt.Printf("Match %d: %s vs %s | Score: %d - %d\n", matchID, team1Name, team2Name, match.Score1, match.Score2)

		time.Sleep(1 * time.Second)
	}

	// İstatistikleri güncelle
	updateStats(team1Name, match.Team1, match.Team2, match)
	updateStats(team2Name, match.Team2, match.Team1, match)

	// Maç sonucunu channel üzerinden gönder
	finalUpdate := fmt.Sprintf("Match %d FINAL: %s vs %s | Final Score: %d - %d", matchID, team1Name, team2Name, match.Score1, match.Score2)
	updates <- finalUpdate
}

// Bir takımın ataklarını simüle etme fonksiyonu
func simulateTeamAttack(team []Player, attackCount int) (int, []Player) {
	score := 0
	for i := 0; i < attackCount; i++ {
		playerIndex := rand.Intn(len(team)) // Rastgele oyuncu seç
		attackResult, player := simulateAttack(team[playerIndex])

		// Oyuncu ve skor istatistiklerini güncelle
		team[playerIndex] = player
		score += attackResult
	}
	return score, team
}

// Her atak için skor sonucunu simüle eden fonksiyon
func simulateAttack(player Player) (int, Player) {
	attackType := rand.Intn(2) // 0: 2 puan, 1: 3 puan

	if attackType == 0 {
		player.TwoPointAttempts++
		if rand.Float32() < 0.65 { // Basit bir başarı olasılığı
			player.TwoPointMade++
			return 2, player
		} else {
			return 0, player
		}
	} else if attackType == 1 {
		player.ThreePointAttempts++
		if rand.Float32() < 0.40 { // Basit bir başarı olasılığı
			player.ThreePointMade++
			return 3, player
		} else {
			return 0, player
		}
	}
	return 0, player
}

// Takım istatistiklerini güncelleme fonksiyonu
func updateStats(teamName string, team1 []Player, team2 []Player, match Game) {
	mutex.Lock()
	defer mutex.Unlock()

	stats, exists := teamStats[teamName]
	if !exists {
		stats = &TeamStats{Team: teamName, Players: team1}
		teamStats[teamName] = stats
	}

	if match.Score1 > match.Score2 {
		stats.Wins++
	} else {
		stats.Losses++
	}
}

// Takımları kazanma sayılarına göre sırala ve yazdır
func printSortedStats() {
	mutex.Lock()
	defer mutex.Unlock()

	var statsSlice []*TeamStats
	for _, stats := range teamStats {
		statsSlice = append(statsSlice, stats)
	}

	sort.Slice(statsSlice, func(i, j int) bool {
		return statsSlice[i].Wins > statsSlice[j].Wins
	})

	fmt.Println("\nTeam Rankings:")
	for i, stats := range statsSlice {
		fmt.Printf("%d. %s Wins: %d, Losses: %d\n", i+1, stats.Team, stats.Wins, stats.Losses)
		for _, player := range stats.Players {
			fmt.Printf("Player %s: 2P Made/Attempts: %d/%d, 3P Made/Attempts: %d/%d\n", player.Name, player.TwoPointMade, player.TwoPointAttempts, player.ThreePointMade, player.ThreePointAttempts)
		}
	}
}

func main() {
	// GetGames by week
	team1 := []Player{{Name: "Alperen Sengun"}, {Name: "Lebron James"}, {Name: "Luka Doncic"}, {Name: "Damien  Lillard"}, {Name: "Antony Edward"}}
	team2 := []Player{{Name: "Joel Embid"}, {Name: "Jokic"}, {Name: "Jalen Bronsun"}, {Name: "Giannis Antetekunto"}, {Name: "Tyres Maxey"}}
	// Diğer takımlar ve oyuncular...
	teamStats["Team East"] = &TeamStats{Team: "Team East", Players: team1}
	teamStats["Team West"] = &TeamStats{Team: "Team West", Players: team2}

	matches := []Game{
		{Team1: team1, Team2: team2},
		//{Team1: "Team C", Team2: "Team D"},
		// Diğer maçlar...
	}
	var wg sync.WaitGroup
	updates := make(chan string, len(matches))

	// create go func every game
	for i, match := range matches {
		wg.Add(1)
		go simulateMatch(i+1, match, "Team East", "Team West", updates, &wg)
	}

	wg.Wait()
	close(updates)

	printSortedStats()
}
