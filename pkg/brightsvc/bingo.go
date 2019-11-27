package brightsvc

import (
	"github.com/shihtzu-systems/bright/pkg/bingo"
	"log"
	"math/rand"
	"time"
)

var christmasMovie = bingo.Boxes{
	"Main Character Returns to Small Town",
	"Storm",
	"Winter Carnival/Festival",
	"Concert",
	"Wise Old Women/Man/Couple",
	"Single Parent",
	"Sob Story",
	"Christmas Theme Name for Character",
	"Going out of Business",
	"Christmas Play",
	"Town with Christmas-themed Name",
	"Hunky Santa",
	"Fake Engagement/Marriage",
	"Travel Setbacks",
	"Dead Parent/Spouse",
	"Main Character Dislikes Holidays",
	"Odd Couple Share a Bed",
	"Odd Couple Teamup",
	"Celebrity Cameo",
	"Real Santa",
	"Busy Career Woman",
	"Movie Title Pun",
	"Decorating Montage",
	"Disapproving Parent",
	"Magical Item",
	"Highschool Sweethearts with Bad Breakup",
	"Sick/Dying Relative",
}

func NewBingo(theme string) bingo.Board {

	var boxes bingo.Boxes
	switch theme {
	case "christmas":
		boxes = christmasMovie
	default:
		log.Fatalf("unknown theme: [%s]", theme)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(boxes), func(i, j int) { boxes[i], boxes[j] = boxes[j], boxes[i] })

	return bingo.Board{
		B: boxes[:5],
		I: boxes[5:10],
		N: boxes[10:15],
		G: boxes[15:20],
		O: boxes[20:25],
	}
}
