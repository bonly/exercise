package main 

import (
  "errors"
  "log"
  "github.com/go-gl/gl/v4.5-core/gl"
  "github.com/go-gl/glfw/v3.1/glfw"
  "github.com/go-gl/glu"
)

const (
  Title = "Nehe 03";
  Width = 640;
  Height = 480;
)

var (
  running bool;
  rotation [3]float32;
  textures []gl.Texture;
  texturefiles [1]string;
)

func init(){
  texturefiles[0] = "/home/opt/Downloads/NeHe.tga";
}

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
  
  if err = initGL(); err != nil{
    log.Fatalf("%v\n", err);
  }
  defer destroyGL();
  
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

func initGL()(err error) {
  if err = loadTextures(); err != nil{
    return;
  }
  
  gl.ShadeModel(gl.SMOOTH); /// 启用阴影平滑,设置平滑着色,阴影平滑通过多边形精细的混合色彩，并对外部光进行平滑
  
  /*
  色彩值的范围从0.0f到1.0f。0.0f代表最黑的情况，1.0f就是最亮的情况。
  glClearColor 后的第一个参数是Red Intensity(红色分量),第二个是绿色，第三个是蓝色。最大值也是1.0f，代表特定颜色分量的最亮情况。
  最后一个参数是Alpha值。当它用来清除屏幕的时候，我们不用关心第四个数字。现在让它为0.0f.
  */
  gl.ClearColor(0, 0, 0, 0); ///设置清除屏幕时所用的颜色,黑色背景
  
  /*
  接下来的三行必须做的是关于depth buffer(深度缓存)的。将深度缓存设想为屏幕后面的层。
  深度缓存不断的对物体进入屏幕内部有多深进行跟踪。
  本程序其实没有真正使用深度缓存，但几乎所有在屏幕上显示3D场景OpenGL程序都使用深度缓存。
  它的排序决定那个物体先画。这样您就不会将一个圆形后面的正方形画到圆形上来。深度缓存是OpenGL十分重要的部分。
  */
  gl.ClearDepth(1); ///设置深度缓存
  gl.DepthFunc(gl.LEQUAL); ///所作深度测试的类型
  gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST); ///告诉系统对透视进行修正,这会十分轻微的影响性能。但使得透视图看起来好一点
  
  gl.Enable(gl.DEPTH_TEST); ///启用深度测试
  gl.Enable(gl.TEXTURE_2D); ///启用纹理映射
  return;
}

func destroyGL(){
  gl.DeleteTextures(textures); ///释放纹理
  textures = nil;
}

func loadTextures()(err error){
  textures = make([]gl.Texture, len(texturefiles));  ///创建文件存储空间数组
  gl.GenTextures(textures); ///告诉OpenGL我们想生成一个纹理名字
  for i := range texturefiles{
    ///告诉OpenGL将纹理名字 texture[0] 绑定到纹理目标上。
    ///2D纹理只有高度(在 Y 轴上)和宽度(在 X 轴上)。主函数将纹理名字指派给纹理数据。
    ///本例中我们告知OpenGL， texture[i] 处的内存已经可用。我们创建的纹理将存储在 texture[i] 的 指向的内存区域。
    textures[i].Bind(gl.TEXTURE_2D); ///绑定

    if !glfw.LoadTexture2D(texturefiles[i], 0){  ///加载图片
      return errors.New("Failed to load texture: " + texturefiles[i]);
    }
    /*
    下面的两行告诉OpenGL在显示图像时，当它比放大得原始的纹理大 ( GL_TEXTURE_MAG_FILTER )
    或缩小得比原始得纹理小( GL_TEXTURE_MIN_FILTER )时OpenGL采用的滤波方式。
    通常这两种情况下我都采用 GL_LINEAR 。这使得纹理从很远处到离屏幕很近时都平滑显示。
    使用 GL_LINEAR 需要CPU和显卡做更多的运算。如果您的机器很慢，您也许应该采用 GL_NEAREST 。
    过滤的纹理在放大的时候，看起来斑驳的很『译者注：马赛克啦』。您也可以结合这两种滤波方式。
    在近处时使用 GL_LINEAR ，远处时 GL_NEAREST 。
    */
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR); ///线形滤波
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR); ///线形滤波
  }
  return;
}

func drawScene(){
  gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT); /// 清除屏幕及深度缓存
  gl.LoadIdentity();  /// 重置模型观察矩阵 
  
  gl.Translatef(0, 0, -5);  /// 并移入屏幕 5.0
  gl.Rotatef (rotation[0], 1, 0, 0); ///以x为轴旋转 参数:角度,X,Y,Z
  gl.Rotatef (rotation[1], 0, 1, 0);
  gl.Rotatef (rotation[2], 0, 0, 1);
  
  rotation[0] += 0.3;
  rotation[1] += 0.2;
  rotation[2] += 0.4;
  
  /*
  选择我们使用的纹理。如果您在您的场景中使用多个纹理，
  应该使用来 glBindTexture(GL_TEXTURE_2D, texture[ 所使用纹理对应的数字 ]) 选择要绑定的纹理。
  当想改变纹理时，应该绑定新的纹理。有一点值得指出的是，您不能在 glBegin() 和 glEnd() 之间绑定纹理，
  必须在 glBegin() 之前或 glEnd() 之后绑定。注意在后面是如何使用 glBindTexture 来指定和绑定纹理的。
  */
  textures[0].Bind (gl.TEXTURE_2D); ///选择纹理
  
  /*
为了将纹理正确的映射到四边形上，必须将纹理的右上角映射到四边形的右上角，纹理的左上角映射到四边形的左上角，
纹理的右下角映射到四边形的右下角，纹理的左下角映射到四边形的左下角。
如果映射错误的话，图像显示时可能上下颠倒，侧向一边或者什么都不是。
glTexCoord2f 的第一个参数是X坐标。 0.0f 是纹理的左侧。 0.5f 是纹理的中点， 1.0f 是纹理的右侧。 
glTexCoord2f 的第二个参数是Y坐标。 0.0f 是纹理的底部。 0.5f 是纹理的中点， 1.0f 是纹理的顶部。

所以纹理的左上坐标是 X：0.0f，Y：1.0f ，四边形的左上顶点是 X： -1.0f，Y：1.0f 。其余三点依此类推。
*/  
  gl.Begin(gl.QUADS); ///绘制正方形
      gl.TexCoord2f(0, 0);  ///前面
	  gl.Vertex3f(-1, -1, 1); ///左下
	  gl.TexCoord2f(1, 0);
	  gl.Vertex3f(1, -1, 1); ///右下
	  gl.TexCoord2f(1, 1);
	  gl.Vertex3f(1, 1, 1); ///右上
	  gl.TexCoord2f(0, 1);
	  gl.Vertex3f(-1, 1, 1);  ///左上
	  
	  gl.TexCoord2d(1, 0); ///后面
	  gl.Vertex3f(-1, -1, -1); ///右下
	  gl.TexCoord2d(1, 1);
	  gl.Vertex3f(-1, 1, -1); ///右上
	  gl.TexCoord2d(0, 1);
	  gl.Vertex3f(1, 1, -1);///左上
	  gl.TexCoord2d(0, 0);
	  gl.Vertex3f(1, -1, -1); //左下
	  
	  gl.TexCoord2d(0, 1); ///上面
	  gl.Vertex3f(-1, 1, -1);
	  gl.TexCoord2d(0, 0);
	  gl.Vertex3f(-1, 1, 1);
	  gl.TexCoord2d(1, 0);
	  gl.Vertex3f(1, 1, 1);
	  gl.TexCoord2d(1, 1);
	  gl.Vertex3f(1, 1, -1);
	  
	  gl.TexCoord2d(1, 1); ///下面
	  gl.Vertex3f(-1, -1, -1);
	  gl.TexCoord2d(0, 1);
	  gl.Vertex3f(1, -1, -1);
	  gl.TexCoord2d(0, 0);
	  gl.Vertex3f(1, -1, 1);
	  gl.TexCoord2d(1, 0);
	  gl.Vertex3f(-1, -1, -1);
	  
	  gl.TexCoord2d(1, 0); ///右面
	  gl.Vertex3f(1, -1, -1);
	  gl.TexCoord2d(1, 1);
	  gl.Vertex3f(1, 1, -1);
	  gl.TexCoord2d(0, 1);
	  gl.Vertex3f(1, 1, 1);
	  gl.TexCoord2d(0, 0);
	  gl.Vertex3f(1, -1, 1);
	  
	  gl.TexCoord2d(0, 0); ///左面
	  gl.Vertex3f(-1, -1, -1);
	  gl.TexCoord2d(1, 0);
	  gl.Vertex3f(-1, -1, 1);
	  gl.TexCoord2d(1, 1);
	  gl.Vertex3f(-1, 1, 1);
	  gl.TexCoord2d(0, 1);
	  gl.Vertex3f(-1, 1, -1);
	  
  gl.End(); ///正方形绘制结束
  
  glfw.SwapBuffers(); ///必须交换显示区才能展现
}

/*
下面的代码的作用是重新设置OpenGL场景的大小，而不管窗口的大小是否已经改变(假定您没有使用全屏模式)。
甚至您无法改变窗口的大小时(例如您在全屏模式下)，它至少仍将运行一次--在程序开始时设置我们的透视图。
OpenGL场景的尺寸将被设置成它显示时所在窗口的大小。
*/
func onResize(w, h int) {
	if h == 0 {
        h = 1;
	}

	gl.Viewport(0, 0, w, h); ///重置当前的视口
	gl.MatrixMode(gl.PROJECTION); ///选择投影矩阵
	gl.LoadIdentity(); ///重置投影矩阵
	glu.Perspective(45.0, float64(w)/float64(h), 0.1, 100.0); ///设置视口的大小
	gl.MatrixMode(gl.MODELVIEW); ///选择模型观察矩阵
	gl.LoadIdentity(); ///重置模型观察矩阵
}
