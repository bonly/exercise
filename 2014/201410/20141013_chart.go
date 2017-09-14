package main  

import(
"github.com/vdobler/chart"
"github.com/vdobler/chart/txtg"
"github.com/ajstarks/svgo"
"github.com/vdobler/chart/svgg"
"image/color"
"math/rand"
"os"
)

func main(){
	box := chart.BoxChart{Title: "Influence of dos`es on effect"}
	box.XRange.Label, box.YRange.Label = "Number of unit doses applied", "Effect [a.u.]"
	box.Key.Pos = "orc"
	box.NextDataSet("Male",
		chart.Style{Symbol: '#', LineColor: color.NRGBA{0x00, 0x00, 0xcc, 0xff}, LineWidth: 1, LineStyle: chart.SolidLine})
	for x := 10; x < 50; x += 5 {
		points := make([]float64, 70)
		a := rand.Float64() * 10
		v := rand.Float64()*5 + 2
		for i := 0; i < len(points); i++ {
			x := rand.NormFloat64()*v + a
			points[i] = x
		}
		box.AddSet(float64(x), points, true)
	}
	box.NextDataSet("Female",
		chart.Style{Symbol: '%', LineColor: color.NRGBA{0xcc, 0x00, 0x00, 0xff}, LineWidth: 1, LineStyle: chart.SolidLine})
	for x := 12; x <= 50; x += 10 {
		points := make([]float64, 60)
		a := rand.Float64()*15 + 30
		v := rand.Float64()*5 + 2
		for i := 0; i < len(points); i++ {
			x := rand.NormFloat64()*v + a
			points[i] = x
		}
		box.AddSet(float64(x), points, true)
	}

	txtFile, _ := os.Create("abc.txt")
	
	tgr := txtg.New(100, 30)
	box.Plot(tgr);

	txtFile.Write([]byte(tgr.String() + "\n\n\n"));
	txtFile.Close();


	svgFile, _ := os.Create("abc.svg")
	sg := svg.New(svgFile)
	sg.Start(500, 500);
	sg.Title("bonly test")
	sg.Rect(0, 0, 500, 500, "fill: #ffffff")	
	sgr := svgg.AddTo(sg, 10, 10, 500, 500, "", 12, color.RGBA{0xff, 0xff, 0xff, 0xff})
	box.Plot(sgr);
	sg.End();
	svgFile.Close();
}

/*
// Fancy tics
	trigc := chart.ScatterChart{Title: ""}
	trigc.XRange.Fixed(0, 4*math.Pi, math.Pi)
	trigc.YRange.Fixed(-1.25, 1.25, 0.5)
	trigc.XRange.TicSetting.Format = func(f float64) string {
		w := int(180*f/math.Pi + 0.5)
		return fmt.Sprintf("%d°", w)
	}
	trigc.AddFunc("Sin(x)", func(x float64) float64 { return math.Sin(x) }, chart.PlotStyleLines,
		chart.Style{Symbol: '@', LineWidth: 2, LineColor: color.NRGBA{0x00, 0x00, 0xcc, 0xff}, LineStyle: 0})
	trigc.AddFunc("Cos(x)", func(x float64) float64 { return math.Cos(x) }, chart.PlotStyleLines,
		chart.Style{Symbol: '%', LineWidth: 2, LineColor: color.NRGBA{0x00, 0xcc, 0x00, 0xff}, LineStyle: 0})
	trigc.XRange.TicSetting.Tics, trigc.YRange.TicSetting.Tics = 1, 1
	trigc.XRange.TicSetting.Mirror, trigc.YRange.TicSetting.Mirror = 2, 2
	trigc.XRange.TicSetting.Grid, trigc.YRange.TicSetting.Grid = 2, 1
	trigc.XRange.ShowZero = true
*/