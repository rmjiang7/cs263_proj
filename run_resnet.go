package main

import (
    tf "github.com/tensorflow/tensorflow/tensorflow/go"
	// "github.com/tensorflow/tensorflow/tensorflow/go/op"
	"bufio"
	"fmt"
	"os"
)

func main() {
	imagePath := "./test_data/cat.jpg"
	file, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err =  buffer.Read(bytes)

	savedModel, err := tf.LoadSavedModel("./trained_models/resnet_v2_fp32_savedmodel_NCHW_jpg/1538687370", []string{"serve"}, nil)
	if err != nil {
		panic(err)
	}

	feedsOutput := tf.Output{
		Op: savedModel.Graph.Operation("input_tensor"),
		Index: 0,
	}

	feedsTensor, err := tf.NewTensor(string(bytes))
	if err != nil {
		panic(err)
	}

	feeds := map[tf.Output]*tf.Tensor{feedsOutput: feedsTensor}

	fetches := []tf.Output{
		{
			Op: savedModel.Graph.Operation("ArgMax"),
			Index: 0,
		},
		{
			Op: savedModel.Graph.Operation("softmax_tensor"),
			Index: 0,
		},
	}

	prediction, err := savedModel.Session.Run(feeds, fetches, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(prediction[0].Value())



	// // Construct a graph with an operation that produces a string constant.
    // s := op.NewScope()
    // c := op.Const(s, "Hello from TensorFlow version " + tf.Version())
    // graph, err := s.Finalize()
    // if err != nil {
    //     panic(err)
    // }

    // // Execute the graph in a session.
    // sess, err := tf.NewSession(graph, nil)
    // if err != nil {
    //     panic(err)
    // }
    // output, err := sess.Run(nil, []tf.Output{c}, nil)
    // if err != nil {
    //     panic(err)
    // }
    // fmt.Println(output[0].Value())
}