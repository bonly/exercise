package db

import(
"testing"
// log "glog"
"flag"
"github.com/gonum/plot"
"github.com/gonum/plot/plotter"
"github.com/gonum/plot/plotutil"
"github.com/gonum/plot/vg"
)

func init(){
	flag.Parse();
	flag.Set("alsologtostderr", "true");
}

func Test_open(ts *testing.T){
	open();
	create_table();
	select_table();	
}

func Test_view(ts *testing.T){
	open();
	day, _ := select_table();
	// log.Infof("%#v", day);

	p, _ := plot.New();
	p.Title.Text = "test";
	p.X.Label.Text = "X";
	p.X.Label.Text = "Y";

	pts := make(plotter.XYs, len(day));
	for idx, dt := range day{
		pts[idx].X = (float64)(dt.Dt);
		pts[idx].Y =  dt.Open;
	}
	plotutil.AddLinePoints(p,
		"First", pts);

	p.Save(4 * vg.Inch, 4 * vg.Inch, "point.svg");
}