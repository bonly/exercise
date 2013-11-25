package main

import (
        "errors"
        "log"

        "github.com/go-gl/gl"
        "github.com/go-gl/glfw"
        "github.com/go-gl/glu"
)

const (
        Title  = "Nehe 07";
        Width  = 640;
        Height = 480;
)

var (
        running      bool;
        textures     []gl.Texture = make([]gl.Texture, 3);    //3种纹理的储存空间
        texturefiles [1]string;
        light        bool;                                     // Display light? 跟踪光照是否打开
        rotation     [2]float32;                               // X/Y rotation. 旋转
        speed        [2]float32;                               // X/Y speed. 旋转速度
        z            float32    = -5;                          // Depth into the scene. 深入屏幕的距离
        ambient      []float32  = []float32{0.5, 0.5, 0.5, 1}; // ambient light colour. 环境光参数
        diffuse      []float32  = []float32{1, 1, 1, 1};       // diffuse light colour. 漫射光参数
        lightpos     []float32  = []float32{0, 0, 2, 1};       // Position of light source. 光源位置
        filter       int;                                      // Index of current texture to display.  滤波类型
)

func init() {
        texturefiles[0] = "/home/opt/Downloads/Crate.tga";
}

func main() {
        var err error;
        if err = glfw.Init(); err != nil {
                log.Fatalf("%v\n", err);
                return;
        }

        defer glfw.Terminate();

        if err = glfw.OpenWindow(Width, Height, 8, 8, 8, 8, 0, 8, glfw.Windowed); err != nil {
                log.Fatalf("%v\n", err);
                return;
        }

        defer glfw.CloseWindow();

        glfw.SetSwapInterval(1);
        glfw.SetWindowTitle(Title);
        glfw.SetWindowSizeCallback(onResize);
        glfw.SetKeyCallback(onKey);

        if err = initGL(); err != nil {
                log.Fatalf("%v\n", err);
                return;
        }

        defer destroyGL();

        running = true;
        for running && glfw.WindowParam(glfw.Opened) == 1 {
                drawScene();
        }
}

func onResize(w, h int) {
        if h == 0 {
                h = 1;
        }

        gl.Viewport(0, 0, w, h);
        gl.MatrixMode(gl.PROJECTION);
        gl.LoadIdentity();
        glu.Perspective(45.0, float64(w)/float64(h), 0.1, 100.0); //设置视窗的大小
        gl.MatrixMode(gl.MODELVIEW);
        gl.LoadIdentity();
}

func onKey(key, state int) {
        switch key {
        case glfw.KeyEsc:
                running = false;
        case 76: // L
                if state == 1 {
                        if light = !light; !light {
                                gl.Disable(gl.LIGHTING);
                        } else {
                                gl.Enable(gl.LIGHTING);
                        }
                }
        case 70: // F
                if state == 1 {
                        if filter++; filter >= len(textures) {
                                filter = 0;
                        }
                }
        case glfw.KeyPageup:
                z -= 0.2;
        case glfw.KeyPagedown:
                z += 0.2;
        case glfw.KeyUp:
                speed[0] -= 0.1;
        case glfw.KeyDown:
                speed[0] += 0.1;
        case glfw.KeyLeft:
                speed[1] -= 0.1;
        case glfw.KeyRight:
                speed[1] += 0.1;
        }
}

func initGL() (err error) {
        if err = loadTextures(); err != nil {
                return;
        }

        gl.ShadeModel(gl.SMOOTH); ///阴影模式设为平滑阴影
        gl.ClearColor(0, 0, 0, 0); ///背景色设为黑色
        gl.ClearDepth(1);
        gl.DepthFunc(gl.LEQUAL);
        gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST); ///启用优化透视计算
        gl.Enable(gl.DEPTH_TEST); ///启用深度测试
        gl.Enable(gl.TEXTURE_2D); ///启用2D纹理映射

/*
创建光源的数组。我们将使用两种不同的光。
第一种称为环境光。环境光来自于四面八方。所有场景中的对象都处于环境光的照射中。
第二种类型的光源叫做漫射光。漫射光由特定的光源产生，并在您的场景中的对象表面上产生反射。
处于漫射光直接照射下的任何对象表面都变得很亮，而几乎未被照射到的区域就显得要暗一些。
这样在我们所创建的木板箱的棱边上就会产生的很不错的阴影效果。
创建光源的过程和颜色的创建完全一致。前三个参数分别是RGB三色分量，最后一个是alpha通道参数。
因此，下面的代码我们得到的是半亮(0.5f)的白色环境光。如果没有环境光，未被漫射光照到的地方会变得十分黑暗。
*/
        gl.Lightfv(gl.LIGHT1, gl.AMBIENT, ambient); //设置环境光(半亮度环境光)
        //gl.Lightfv(gl.LIGHT1, gl.AMBIENT, diffuse); //设置漫射光
        gl.Lightfv(gl.LIGHT1, gl.DIFFUSE, diffuse); //设置漫射光(全亮度白光)
        gl.Lightfv(gl.LIGHT1, gl.POSITION, lightpos);//设置光源位置
        gl.Enable(gl.LIGHT1); //启用一号光源
/*
光源的位置。前三个参数和glTranslate中的一样。依次分别是XYZ轴上的位移。
由于我们想要光线直接照射在木箱的正面，所以XY轴上的位移都是0.0f。
第三个值是Z轴上的位移。为了保证光线总在木箱的前面，所以我们将光源的位置朝着观察者(就是您哪。)挪出屏幕。
我们通常将屏幕也就是显示器的屏幕玻璃所处的位置称作Z轴的0.0f点。所以Z轴上的位移最后定为2.0f。
假如您能够看见光源的话，它就浮在您显示器的前方。当然，如果木箱不在显示器的屏幕玻璃后面的话，您也无法看见箱子。
*/        
        return;
}

func destroyGL() {
        gl.DeleteTextures(textures);
        textures = nil;
}

func loadTextures() (err error) {
        gl.GenTextures(textures);  //申请纹理处理

        // Texture 1
        textures[0].Bind(gl.TEXTURE_2D); //绑定将处理的纹理

        if !glfw.LoadTexture2D(texturefiles[0], 0) { //加载纹理
                return errors.New("Failed to load texture: " + texturefiles[0]);
        }
/*
LINEAR线性滤波的纹理贴图。这需要机器有相当高的处理能力，但它们看起来很不错。
我们接着要创建的第一种纹理使用 GL_NEAREST方式。从原理上讲，这种方式没有真正进行滤波。
它只占用很小的处理能力，看起来也很差。唯一的好处是这样我们的工程在很快和很慢的机器上都可以正常运行。
您会注意到我们在 MIN 和 MAG 时都采用了GL_NEAREST,你可以混合使用 GL_NEAREST 和 GL_LINEAR。
纹理看起来效果会好些，但我们更关心速度，所以全采用低质量贴图。MIN_FILTER在图像绘制时小于贴图的原始尺寸时采用。
MAG_FILTER在图像绘制时大于贴图的原始尺寸时采用。
*/
        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST); //小于原始尺寸时用NEAREST处理
        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST); //大于原始尺寸时用NEAREST处理

        // Texture 2
        textures[1].Bind(gl.TEXTURE_2D); //绑定第二个纹理

        if !glfw.LoadTexture2D(texturefiles[0], 0) {//加载纹理 
                return errors.New("Failed to load texture: " + texturefiles[0]);
        }

        //纹理缩放时采用的处理方式
        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);

/*
当图像在屏幕上变得很小的时候，很多细节将会丢失。刚才还很不错的图案变得很难看。
当您告诉OpenGL创建一个 mipmapped的纹理后，OpenGL将尝试创建不同尺寸的高质量纹理。
当您向屏幕绘制一个 mipmapped纹理的时候，OpenGL将选择它已经创建的外观最佳的纹理(带有更多细节)来绘制，而不仅仅是缩放原先的图像(这将导致细节丢失)。
我曾经说过有办法可以绕过OpenGL对纹理宽度和高度所加的限制——64、128、256，等等。
办法就是 gluBuild2DMipmaps。据我的发现，您可以使用任意的位图来创建纹理。OpenGL将自动将它缩放到正常的大小。
gluBuild2DMipmaps(GL_TEXTURE_2D, 3, TextureImage[0]->sizeX, TextureImage[0]->sizeY, GL_RGB, GL_UNSIGNED_BYTE, TextureImage[0]->data); 
使用三种颜色(红，绿，蓝)来生成一个2D纹理。
TextureImage[0]->sizeX 是位图宽度，extureImage[0]->sizeY 是位图高度，GL_RGB意味着我们依次使用RGB色彩。
GL_UNSIGNED_BYTE 意味着纹理数据的单位是字节。TextureImage[0]->data指向我们创建纹理所用的位图。
*/
        // Texture 3
        textures[2].Bind(gl.TEXTURE_2D); //绑定处理第三个纹理 

        if !glfw.LoadTexture2D(texturefiles[0], glfw.BuildMipmapsBit) { //加载纹理
                return errors.New("Failed to load texture: " + texturefiles[0]);
        }

        //指定缩放处理方式
        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR_MIPMAP_NEAREST); //golang用这个方式自动处理gluBuild2DMipmaps?

        return;
}

func drawScene() {
        gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT); //清除屏幕和深度缓存
        gl.LoadIdentity(); //重置当前的模型观察矩阵

        gl.Translatef(0, 0, z); //移入/移出屏幕 z 个单位

        gl.Rotatef(rotation[0], 1, 0, 0); //绕X轴旋转
        gl.Rotatef(rotation[1], 0, 1, 0); //绕Y轴旋转

        rotation[0] += speed[0];
        rotation[1] += speed[1];

        textures[filter].Bind(gl.TEXTURE_2D); //绑定的纹理

/*
Normal就是法线的意思，所谓法线是指经过面(多边形）上的一点且垂直于这个面(多边形)的直线。
使用光源的时候必须指定一条法线。法线告诉OpenGL这个多边形的朝向，并指明多边形的正面和背面。
如果没有指定法线，什么怪事情都可能发生：不该照亮的面被照亮了，多边形的背面也被照亮....。对了，法线应该指向多边形的外侧。

看着木箱的前面您会注意到法线与Z轴正向同向。这意味着法线正指向观察者－您自己。这正是我们所希望的。
对于木箱的背面，也正如我们所要的，法线背对着观察者。
如果立方体沿着X或Y轴转个180度的话，前侧面的法线仍然朝着观察者，背面的法线也还是背对着观察者。
换句话说，不管是哪个面，只要它朝着观察者这个面的法线就指向观察者。由于光源紧邻观察者，任何时候法线对着观察者时，这个面就会被照亮。
并且法线越朝着光源，就显得越亮一些。如果您把观察点放到立方体内部，你就会法线里面一片漆黑。因为法线是向外指的。
如果立方体内部没有光源的话，当然是一片漆黑。
*/
        gl.Begin(gl.QUADS);
	        // Front Face 前面
	        gl.Normal3f(0, 0, 1); // Normal Pointing Towards Viewer 法线指向观察者
	        gl.TexCoord2f(0, 0);
	        gl.Vertex3f(-1, -1, 1); // Point 1 (Front)
	        gl.TexCoord2f(1, 0);
	        gl.Vertex3f(1, -1, 1); // Point 2 (Front)
	        gl.TexCoord2f(1, 1);
	        gl.Vertex3f(1, 1, 1); // Point 3 (Front)
	        gl.TexCoord2f(0, 1);
	        gl.Vertex3f(-1, 1, 1); // Point 4 (Front)
	        // Back Face 后面
	        gl.Normal3f(0, 0, -1); // Normal Pointing Away From Viewer 法线背向观察者
	        gl.TexCoord2f(1, 0);
	        gl.Vertex3f(-1, -1, -1); // Point 1 (Back)
	        gl.TexCoord2f(1, 1);
	        gl.Vertex3f(-1, 1, -1); // Point 2 (Back)
	        gl.TexCoord2f(0, 1);
	        gl.Vertex3f(1, 1, -1); // Point 3 (Back)
	        gl.TexCoord2f(0, 0);
	        gl.Vertex3f(1, -1, -1); // Point 4 (Back)
	        // Top Face 顶面
	        gl.Normal3f(0, 1, 0); // Normal Pointing Up  法线向上
	        gl.TexCoord2f(0, 1);
	        gl.Vertex3f(-1, 1, -1); // Point 1 (Top)
	        gl.TexCoord2f(0, 0);
	        gl.Vertex3f(-1, 1, 1); // Point 2 (Top)
	        gl.TexCoord2f(1, 0);
	        gl.Vertex3f(1, 1, 1); // Point 3 (Top)
	        gl.TexCoord2f(1, 1);
	        gl.Vertex3f(1, 1, -1); // Point 4 (Top)
	        // Bottom Face 底面
	        gl.Normal3f(0, -1, 0); // Normal Pointing Down 法线朝下
	        gl.TexCoord2f(1, 1);
	        gl.Vertex3f(-1, -1, -1); // Point 1 (Bottom)
	        gl.TexCoord2f(0, 1);
	        gl.Vertex3f(1, -1, -1); // Point 2 (Bottom)
	        gl.TexCoord2f(0, 0);
	        gl.Vertex3f(1, -1, 1); // Point 3 (Bottom)
	        gl.TexCoord2f(1, 0);
	        gl.Vertex3f(-1, -1, 1); // Point 4 (Bottom)
	        // Right face 右面
	        gl.Normal3f(1, 0, 0); // Normal Pointing Right 法线朝右
	        gl.TexCoord2f(1, 0);
	        gl.Vertex3f(1, -1, -1); // Point 1 (Right)
	        gl.TexCoord2f(1, 1);
	        gl.Vertex3f(1, 1, -1); // Point 2 (Right)
	        gl.TexCoord2f(0, 1);
	        gl.Vertex3f(1, 1, 1); // Point 3 (Right)
	        gl.TexCoord2f(0, 0);
	        gl.Vertex3f(1, -1, 1);// Point 4 (Right)
	        // Left Face 左面
	        gl.Normal3f(-1, 0, 0); // Normal Pointing Left  法线朝左
	        gl.TexCoord2f(0, 0);
	        gl.Vertex3f(-1, -1, -1);// Point 1 (Left)
	        gl.TexCoord2f(1, 0);
	        gl.Vertex3f(-1, -1, 1); // Point 2 (Left)
	        gl.TexCoord2f(1, 1);
	        gl.Vertex3f(-1, 1, 1);// Point 3 (Left)
	        gl.TexCoord2f(0, 1);
	        gl.Vertex3f(-1, 1, -1);
        gl.End();

        glfw.SwapBuffers();
}
