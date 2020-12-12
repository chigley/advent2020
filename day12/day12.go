package day12

import (
	"fmt"
	"strconv"

	"github.com/chigley/advent2020"
)

type Ship struct {
	pos advent2020.XY

	// Only used in part 1
	dir advent2020.CompassDirection

	// Only used in part 2
	waypoint advent2020.XY
}

func (s *Ship) Move(dir advent2020.CompassDirection, steps int) {
	s.pos = s.pos.Add(advent2020.Compass[dir].Multiply(steps))
}

func (s *Ship) MoveWaypoint(dir advent2020.CompassDirection, steps int) {
	s.waypoint = s.waypoint.Add(advent2020.Compass[dir].Multiply(steps))
}

func Part1(in []string) (int, error) {
	ship := Ship{
		dir: advent2020.East,
	}

	for _, l := range in {
		n, err := strconv.Atoi(l[1:])
		if err != nil {
			return 0, fmt.Errorf("day12: atoi: %w", err)
		}

		dir := l[0]
		switch dir {
		case 'N':
			ship.Move(advent2020.North, n)
		case 'S':
			ship.Move(advent2020.South, n)
		case 'E':
			ship.Move(advent2020.East, n)
		case 'W':
			ship.Move(advent2020.West, n)
		case 'L':
			ship.dir = ship.dir.RotateAnticlockwise(n / 90)
		case 'R':
			ship.dir = ship.dir.RotateClockwise(n / 90)
		case 'F':
			ship.Move(ship.dir, n)
		default:
			return 0, fmt.Errorf("day12: unknown action: %q", dir)
		}
	}

	return ship.pos.Manhattan(), nil
}

func Part2(in []string) (int, error) {
	ship := Ship{
		waypoint: advent2020.XY{X: 10, Y: 1},
	}

	for _, l := range in {
		n, err := strconv.Atoi(l[1:])
		if err != nil {
			return 0, fmt.Errorf("day12: atoi: %w", err)
		}

		dir := l[0]
		switch dir {
		case 'N':
			ship.MoveWaypoint(advent2020.North, n)
		case 'S':
			ship.MoveWaypoint(advent2020.South, n)
		case 'E':
			ship.MoveWaypoint(advent2020.East, n)
		case 'W':
			ship.MoveWaypoint(advent2020.West, n)
		case 'L':
			ship.waypoint = ship.waypoint.RotateAnticlockwise(n / 90)
		case 'R':
			ship.waypoint = ship.waypoint.RotateClockwise(n / 90)
		case 'F':
			ship.pos = ship.pos.Add(ship.waypoint.Multiply(n))
		default:
			return 0, fmt.Errorf("day12: unknown action: %q", dir)
		}
	}

	return ship.pos.Manhattan(), nil
}
