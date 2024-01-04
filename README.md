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

### Simulate Games: `POST /v1/games/simulation`

- Description: Initiates a simulation of all games scheduled for the given week, resulting in real-time score updates and changes to the game, team, and player statistics.

- Payload Parameters:
  - `week`: The week number to simulate games for.

Initiates the simulation process for all games in the specified week. This endpoint creates or modifies resources representing the games' outcomes and should be used to trigger a new simulation.

---

### 2. Retrieve Schedule: `GET /v1/games`
Query Parameters:
- `week`: The week number to retrieve the schedule for.

Description:
Retrieves the game schedule for the specified week.

---

### 3. Team Standings: `GET /v1/teams`
Query Parameters:
- `sort`: Send spesific sort parameter for a lose and win field example: lose.asc
Description:
Retrieves the current standings of all teams, sorted by wins and other criteria.

---

### 4. Player Statistics: `GET /v1/players`

Description:
Provides detailed statistics for each player, including points, attempts, and other relevant metrics.

