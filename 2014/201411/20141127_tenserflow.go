package main 
import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"fmt"
)

func Const(g *tf.Graph, name string, value interface{}) (tf.Output, error) {
	t, ok := value.(*tf.Tensor)
	if !ok {
		fmt.Println("not tensor var");
		var err error
		if t, err = tf.NewTensor(value); err != nil {
			fmt.Println("new tensor failed");
			return tf.Output{}, err
		}
	}
	op, err := g.AddOperation(tf.OpSpec{
		Type: "Const",
		Name: name,
		Attrs: map[string]interface{}{
			"dtype": t.DataType(),
			"value": t,
		},
	})
	fmt.Println("add op exec");
	return op.Output(0), err
}

func main(){
	graph := tf.NewGraph();
	c1, err := Const(graph, "c1", int32(1));
	if err != nil{
		fmt.Println(err);
		return;
	}
	c2, err := Const(graph, "c2", int32(3));
	if err != nil{
		fmt.Println(err);
		return;
	}

	op, err := graph.AddOperation(tf.OpSpec{
		Type: "Concat",
		Input: []tf.Input{c1, c2},
	});
	if err != nil{
		fmt.Println(err);
		return;
	}

	session, err := tf.NewSession(graph,  &tf.SessionOptions{});
	if err != nil{
		fmt.Println(err);
		return;
	}
	defer session.Close();

	output, err := session.Run(nil, []tf.Output{op.Output(0)}, nil);
	if err != nil{
		fmt.Println(err);
		return;
	}
	fmt.Printf("%#v\n", output);
}

