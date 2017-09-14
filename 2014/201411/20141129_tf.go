package main 
import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
	"fmt"
)

func multiplie(){
	sc     := op.NewScope();
	input  := op.Placeholder(sc, tf.Float);
	output := op.MatMul(sc,
		op.Const(sc, [][]float32{{10}, {20}}),
		input,
		op.MatMulTransposeB(true));
	if sc.Err() != nil{
		panic(sc.Err());
	} 

	shape, _ := output.Shape();
	fmt.Println(shape);
}

func main(){
	multiplie();

}

