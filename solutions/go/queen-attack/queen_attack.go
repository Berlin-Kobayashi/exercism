// Package queenattack implements a function for checking if two queens can attack each other on a chess board.
package queenattack

import (
	"errors"
	"regexp"
	"strconv"
)

const testVersion = 1

type chessPosition string

func (chessPos chessPosition) ToCoordinates() (x, y int, err error) {
	chessPositionRegex := regexp.MustCompile("^([a-h])([1-8])$")

	if !chessPositionRegex.Match([]byte(chessPos)) {
		return 0, 0, errors.New("Given chess position is not valid: " + string(chessPos))
	}

	xChessCoordinate := chessPositionRegex.ReplaceAll([]byte(chessPos), []byte("$1"))
	x = int(xChessCoordinate[0] - 'a')
	yChessCoordinate := chessPositionRegex.ReplaceAll([]byte(chessPos), []byte("$2"))
	y, _ = strconv.Atoi(string(yChessCoordinate))
	y--

	return x, y, nil
}

// CanQueenAttack checks if two queens can attack each other on a chess board.
func CanQueenAttack(whitePosition, blackPosition string) (result bool, err error) {
	whiteX, whiteY, whiteErr := chessPosition(whitePosition).ToCoordinates()
	if whiteErr != nil {
		return false, whiteErr
	}

	blackX, blackY, blackErr := chessPosition(blackPosition).ToCoordinates()
	if blackErr != nil {
		return false, blackErr
	}

	if whiteX == blackX && whiteY == blackY {
		return false, errors.New("Given queen positions are equal")
	}

	if whiteX == blackX || whiteY == blackY {
		return true, nil
	}

	if whiteX-blackX == whiteY-blackY || blackX-whiteX == whiteY-blackY {
		return true, nil
	}

	return false, nil
}
