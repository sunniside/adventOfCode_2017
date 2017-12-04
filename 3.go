package main

import "fmt"

func main() {
  partOne();
  fmt.Println("");
}

func partOne() {
  buildSpiralGrid();
}


func buildSpiralGrid() ([500000]Square) {
  currentPoint := Point{0,0};
  xCount := 1;
  yCount := 1;
  squareIndex := 1;
  var squares [500000]Square;
  squares[0] = Square{0, currentPoint, 1};
  for squareIndex <= 361526 {
    if(squareIndex == 0) {
      continue;
    }
    horizontal := xCount == yCount;

    squaresToPlace := 0;
    if(horizontal) {
      squaresToPlace = xCount;
    } else {
      squaresToPlace = yCount;
    }

    for j := 0; j < squaresToPlace; j++ {
      currentPoint = calculatePoint(squaresToPlace, currentPoint, horizontal);
      square := Square{squareIndex, currentPoint, 0}
      square.value = calculateValue(squares, square);
      fmt.Println(square.value);
      if(square.value > 361527) {
        return squares;
      }
      squares[squareIndex] = square;
      squareIndex++;
    }
    if(horizontal) {
      xCount++;
    } else {
      yCount++;
    }
  }
  theSquare := squares[361526];

  fmt.Println(theSquare.location);
  fmt.Println(theSquare.location.x + theSquare.location.y);

  return squares;
}

func calculateValue(squares [500000]Square, square Square) (int) {
   neighbors := findAdjacentSquares(squares, square);
   for i := 0; i < len(neighbors); i++ {
      square.value += neighbors[i].value;
   }
   return square.value;
}

func findAdjacentSquares(squares [500000]Square, square Square) ([]Square) {
  var neighbors []Square;
  currentPoint := square.location;
  possiblePoints := getPossibleAdjacentPoints(currentPoint);
  for i := 0; i < square.index; i++ {
      candidate := squares[i];
      candidatePoint := candidate.location;
      for j := 0; j < len(possiblePoints); j++ {
        if(candidatePoint == possiblePoints[j]) {
          neighbors = append(neighbors, candidate);
        }
      }
  }
  return neighbors;
}

func getPossibleAdjacentPoints(currentPoint Point) ([]Point){
    p1 := Point{currentPoint.x - 1, currentPoint.y};
    p2 := Point{currentPoint.x - 1, currentPoint.y - 1};
    p3 := Point{currentPoint.x - 1, currentPoint.y + 1};
    p4 := Point{currentPoint.x, currentPoint.y + 1};
    p5 := Point{currentPoint.x, currentPoint.y - 1 };
    p6 := Point{currentPoint.x + 1, currentPoint.y - 1 };
    p7 := Point{currentPoint.x + 1, currentPoint.y + 1 };
    p8 := Point{currentPoint.x + 1, currentPoint.y};

    return []Point{p1, p2, p3, p4, p5, p6, p7, p8};
}

func calculatePoint(squaresToPlace int, point Point, horizontal bool) (Point) {
  if(squaresToPlace % 2 == 0) {
    if(horizontal) {
      point.x -= 1;
    } else {
      point.y -= 1
    }
  } else {
    if(horizontal) {
      point.x += 1;
    } else {
      point.y += 1
    }
  }
  return point;
}


type Point struct {
  x, y int;
}

type Square struct {
  index int;
  location Point;
  value int;
}
