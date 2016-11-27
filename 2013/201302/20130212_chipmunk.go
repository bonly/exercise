package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

func main() {

	var space *chipmunk.Space

	space = chipmunk.NewSpace()
	space.Gravity = vect.Vect{0, -900}

	// Add a static body - lines etc.
	staticBody := chipmunk.NewBodyStatic()
	space.AddBody(staticBody)

	// Create a ball
	ball := chipmunk.NewCircle(vect.Vector_Zero, float32(10))
	ball.SetElasticity(0.95)

	// Create a body for the ball
	body := chipmunk.NewBody(vect.Float(1), ball.Moment(float32(1)))
	body.SetPosition(vect.Vect{vect.Float(1), 600.0})
	body.SetAngle(vect.Float(rand.Float32() * 2 * math.Pi))
	body.AddShape(ball)
	space.AddBody(body)

	// I love tickers for this kind of stuff. We want to calc the state of the
	// space every 60th of a second.
	ticker := time.NewTicker(time.Second / 60)
	for {
		log.Printf("ball %+v", body.Position())
		space.Step(vect.Float(1.0 / 60.0))
		<-ticker.C // wait up to 1/60th of a second
	}
}