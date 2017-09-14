package main  

import (
"github.com/gizak/termui"
)

func main(){
	err := termui.Init();
	if err != nil{
		panic(err);
	}
	defer termui.Close();

	lc := termui.NewLineChart();
	lc.BorderLabel = "data"
	lc.Mode = "dot";
	// lc.DotStyle = '+';
	lc.Data = []float64{1.1, 2.2, 3.4, 4.1, 5.5, 6.1, 7.0};
	lc.DataLabels = []string{"A", "B", "C", "D","E","F"};
	lc.Width = 77;
	lc.Height = 16;
	lc.X = 0;
	lc.Y = 12
	lc.AxesColor = termui.ColorWhite;
	lc.LineColor = termui.ColorCyan | termui.AttrBold;

	termui.Render(lc);
	termui.Handle("/sys/kbd/q", func(termui.Event){
		termui.StopLoop();
	});
	termui.Loop();
}