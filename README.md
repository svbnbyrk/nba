# Basketball Game Simulation Service

This service simulates basketball games with real-time score updates and provides RESTful APIs to retrieve game simulations, schedules, team standings, and player statistics.

## Features

- **Game Simulation:** Simulate basketball games in a fast-forwarded real-time manner, reflecting each game's progress and final results.
- **Schedule Retrieval:** Retrieve game schedules for specific weeks.
- **Team Standings:** Access current team standings.
- **Player Statistics:** Get detailed statistics for each player.

## Getting Started

### Prerequisites

- Go (Version 1.20 or later)
- PostgreSQL- for storing game, team, and player data

### Installation

Clone the repository:
   ```bash
   git clone
   cd nba
   make build
   make run
   ```
   or
   ```bash
   docker compose up
   ```

## API Endpoints

### 1. Simulate Games: `GET /v1/games/simulation`
Query Parameters:
- `week`: The week number to simulate games for.

Description:
Simulates all games scheduled for the given week and provides real-time score updates.

---

### 2. Retrieve Schedule: `GET /v1/games/schedule`
Query Parameters:
- `week`: The week number to retrieve the schedule for.

Description:
Retrieves the game schedule for the specified week.

---

### 3. Team Standings: `GET /v1/teams/standings`

Description:
Retrieves the current standings of all teams, sorted by wins and other criteria.

---

### 4. Player Statistics: `GET /v1/players/stats`

Description:
Provides detailed statistics for each player, including points, attempts, and other relevant metrics.

