package main 

import (
  "log"
  "github.com/go-gl/gl"
  "github.com/go-gl/glfw"
  "github.com/go-gl/glu"
)

const (
  Title = "Nehe 04";
  Width = 640;
  Height = 480;
)

var (
  trisAngle float32;
  quadAngle float32;
  running bool;
)

func main() {
  var err error;
  if err = glfw.Init(); err != nil { ///初始化环境
    log.Fatalf("%v\n", err);
    return;
  }
  defer glfw.Terminate(); /// 销毁环境
  
  if err = glfw.OpenWindow(Width, Height, 8, 8, 8, 8, 0, 8, glfw.Windowed); err != nil { ///创建窗口
    log.Fatalf("%v\n", err);
    return;
  }
  defer glfw.CloseWindow(); /// 销毁窗口
  
  glfw.SetSwapInterval(1);
  glfw.SetWindowTitle(Title); ///设置标题
  glfw.SetWindowSizeCallback(onResize); /// 回调窗口变化
  glfw.SetKeyCallback(onKey); ///回调按键
  
  initGL();
  
  running = true;
  for running && glfw.WindowParam(glfw.Opened) == 1 {
    drawScene();
  }
}

func onKey(key, state int){
  switch key {
    case glfw.KeyEsc:
      running = false;
  }
}

func initGL() {
  gl.ShadeModel(gl.SMOOTH); /// 设置平滑着色?(是默认的)
  //gl.ShadeModel(gl.FLAT); /// 即单色,使用最后一个点所用的颜色
  gl.ClearColor(0, 0, 0, 0);
  gl.ClearDepth(1);
  gl.Enable(gl.DEPTH_TEST);
  gl.DepthFunc(gl.LEQUAL);
  gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST);
}

func drawScene(){
  gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT); /// 清除屏幕及深度缓存
  
  gl.LoadIdentity();  /// 重置模型观察矩阵 
  gl.Translatef(-1.5, 0, -6);  /// 左移 1.5 单位，并移入屏幕 6.0
  gl.Rotatef (trisAngle, 0, 1, 0); ///以y为轴 参数:角度,X,Y,Z
  
  gl.Begin(gl.TRIANGLES);  /// 绘制三角形
	  gl.Color3f(1, 0, 0);  ///  设置当前色为红色
	  gl.Vertex3f(0, 1, 0); ///上顶点
	  gl.Color3f(0, 1, 0);  /// 设置当前色为绿色
	  gl.Vertex3f(-1, -1, 0); /// 左下
	  gl.Color3f(0, 0, 1); ///设置当前色为蓝色
	  gl.Vertex3f(1, -1, 0); ///右下
  gl.End(); ///三角形绘制结束,三角形将被填充。
  //但是因为每个顶点有不同的颜色，因此看起来颜色从每个角喷出，并刚好在三角形的中心汇合，三种颜色相互混合。这就是平滑着色。

  /// 要让对象绕自身的轴旋转，必须让对象的中心坐标总是(0.0f,0,0f,0,0f),因此这里的四边形是满屏跑的
  gl.LoadIdentity(); ///将当前点移到了屏幕中心，X坐标轴从左至右，Y坐标轴从下至上，Z坐标轴从里至外。
  ///OpenGL屏幕中心的坐标值是X和Y轴上的0.0f点。中心左面的坐标值是负值，右面是正值。移向屏幕顶端是正值，移向屏幕底端是负值。移入屏幕深处是负值，移出屏幕则是正值。
  gl.Rotatef(quadAngle, 1, 0, 0); /// 以x为轴
  gl.Translatef(1.5, 0, -6); ///以当前点为起始点移动?(是的)-6为距离
  //gl.Translatef(3, 0, -6); ///右移3单位,看不见?(因为loadIdentity()置中了)
  gl.Color3f(0.5, 0.5, 1.0); ///一次性将当前色设置为蓝色

  /// @note 顺时针绘制的正方形意味着我们所看见的是四边形的背面
  gl.Begin(gl.QUADS); ///绘制正方形
	  gl.Vertex3f(-1, 1, 0); /// 左上
	  gl.Vertex3f(1, 1, 0); /// 右上
	  gl.Vertex3f(1, -1, 0); /// 右下
	  gl.Vertex3f(-1, -1, 0); /// 左下
  gl.End(); ///正方形绘制结束

  trisAngle += 0.2;
  quadAngle -= 0.15;
  
  glfw.SwapBuffers(); ///必须交换显示区才能展现
}

func onResize(w, h int) {
	if h == 0 {
        h = 1;
	}

	gl.Viewport(0, 0, w, h);
	gl.MatrixMode(gl.PROJECTION);
	gl.LoadIdentity();
	glu.Perspective(45.0, float64(w)/float64(h), 0.1, 100.0);
	gl.MatrixMode(gl.MODELVIEW);
	gl.LoadIdentity();
}
