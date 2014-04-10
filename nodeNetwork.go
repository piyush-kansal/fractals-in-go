// Creates a node network
package main

// Import packages
import ("image"; "image/color"; "image/png";
		"os"
		"log"
		"math/rand"
		"time")
		// "coordinate")

type Node struct {
	Position 	Coordinate
	Ch 			chan *Node
	Peers		[]*Node
	Canvas		*Canvas
	Power		uint8
}

func (n *Node) Listen() {
	// Listen for incoming connection on channel
	for {
		peer := <- n.Ch
		peer.Power -= 5
		n.Power = peer.Power
		n.Canvas.DrawLine(color.RGBA{255, n.Power, 0, 255}, n.Position, peer.Position)
	}

	// Retransmit
	if n.Power > 0 {
		go n.Send()
	}
}

func (n *Node) Send() {
	for _, target := range n.Peers {
		if target.Power == 0 {
			target.Ch <- n
			break
		}
	}
}

func NewNode(peers uint32, canvas Canvas) {
	n := new(Node)
	size := canvas.Bounds().Size()
	n.Position := Coordinate{size.X*rand.Float64(), size.Y*rand.Float64()}
	n.Ch := make(chan *Node)
	n.Peers = make([]*Node, 0, peers)
	n.Canvas := canvas
	n.Power := 0
	go n.Listen()
	return n
}

func initialize(totalNodes uint32, peers uint32, canvas Canvas) {
	nodes = make([]*Node, totalNodes)

	for i:=0; i<totalNodes; i++ {
		nodes[i] = NewNode(peers, canvas)
	}
}

func main() {
	// Declare image size
	width, height := 1024, 1024

	// Create a new image
	img := image.Rect(0, 0, width, height)
	c := NewCanvas(img)
	c.DrawRect(color.RGBA{0, 0, 0, 255}, Coordinate{0, 0}, Coordinate{float64(width), float64(height)})
	rand.Seed(time.Now().UTC().UnixNano())

	initialize(50, 5, c)

	fileName := "nodeNetwork.png"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	png.Encode(file, c)

}