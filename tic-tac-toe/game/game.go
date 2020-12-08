package game

import (
	"tic-tac-toe/analyzer"
	"tic-tac-toe/board"
	gamestatus "tic-tac-toe/gameStatus"
	"tic-tac-toe/player"
)

// Game is to create game type
type Game struct {
	players       []*player.Player
	board         *board.Board
	analyzer      *analyzer.Analyzer
	gameStatus    string
	addToken      bool
	currentPlayer *player.Player
	nextPlayer    *player.Player
}

// NewPlayer is to create a new struct of type game
func NewPlayer(playersList []*player.Player, size int) *Game {
	tempBoard := board.NewBoard(size)
	tenmpAnalyzer := analyzer.NewAnalyzer(tempBoard)
	tempGame := &Game{
		players:       playersList,
		board:         tempBoard,
		analyzer:      tenmpAnalyzer,
		gameStatus:    gamestatus.INPROGRESS,
		addToken:      true,
		currentPlayer: playersList[0],
		nextPlayer:    playersList[1],
	}
	return tempGame
}

// IsAddToken to chcek if cell is marked
func (g *Game) IsAddToken() bool {
	return g.addToken
}

// GetBoard to return the board
func (g *Game) GetBoard() *board.Board {
	return g.GetBoard()
}

// GetPlayers to return the players
func (g *Game) GetPlayers() []*player.Player {
	return g.GetPlayers()
}

// GetAnalyzer ti return the analyzer
func (g *Game) GetAnalyzer() *analyzer.Analyzer {
	return g.GetAnalyzer()
}

// GetGameStatus to return the current game status
func (g *Game) GetGameStatus() string {
	return g.GetGameStatus()
}

// GetCurrentPlayer to return the current player
func (g *Game) GetCurrentPlayer() *player.Player {
	return g.currentPlayer
}

func (g *Game) play(boardNumber int) {
	var tempPlayer player.Player

	addToken = addXO(boardNumber, currentPlayer.getMark())
	if addToken == false {
		return
	}

	if iResultAnalyzable.checkStatus(currentPlayer.getMark()) {
		gameStatus = GameStatus.WIN
		return
	}
	if iBoardable.checkIfBoardFull() {
		gameStatus = GameStatus.DRAW
		return
	}
	tempPlayer = currentPlayer
	currentPlayer = nextPlayer
	nextPlayer = tempPlayer
}
