/*
混合的基本原理是就将要分色的图像各象素的颜色以及背景颜色均按照RGB规则各自分离之后，
根据－图像的RGB颜色分量*alpha值+背景的RGB颜色分量*(1-alpha值)－这样一个简单公式来混合之后，最后将混合得到的RGB分量重新合并。』
公式如下：
(Rs Sr + Rd Dr, Gs Sg + Gd Dg, Bs Sb + Bd Db, As Sa + Ad Da)
OpenGL按照上面的公式计算这两个象素的混色结果。小写的s和r分别代表源象素和目标象素。大写的S和D则是相应的混色因子。
这些决定了您如何对这些象素混色。绝大多数情况下，各颜色通道的alpha混色值大小相同，这样对源象素就有 (As, As, As, As)，
目标象素则有1, 1, 1, 1) - (As, As, As, As)。上面的公式就成了下面的模样:
(Rs As + Rd (1 - As), Gs As + Gd (1 - As), Bs As + Bs (1 - As), As As + Ad (1 - As))
这个公式会生成透明/半透明的效果。
OpenGL中的混色
在OpenGL中实现混色的步骤类似于我们以前提到的OpenGL过程。接着设置公式，并在绘制透明对象时关闭写深度缓存。
因为我们想在半透明的图形背后绘制 对象。这不是正确的混色方法，但绝大多数时候这种做法在简单的项目中都工作的很好。
Rui Martins 的补充： 正确的混色过程应该是先绘制全部的场景之后再绘制透明的图形。并且要按照与深度缓存相反的次序来绘制(先画最远的物体)。
考虑对两个多边形(1和2)进行alpha混合，不同的绘制次序会得到不同的结果。
(这里假定多边形1离观察者最近，那么正确的过程应该先画多边形2，再画多边形1。
正如您再现实中所见到的那样，从这两个<透明的>多边形背后照射来的光线总是先穿过多边形2，再穿过多边形1，最后才到达观察者的眼睛。)
在深度缓存启用时，您应该将透明图形按照深度进行排序，并在全部场景绘制完毕之后再绘制这些透明物体。
否则您将得到不正确的结果。我知道某些时候这样做是很令人痛苦的，但这是正确的方法
*/
package main

import (
        "errors"
        "log"

        "github.com/go-gl/gl"
        "github.com/go-gl/glfw"
        "github.com/go-gl/glu"
)

const (
        Title  = "Nehe 08";
        Width  = 640;
        Height = 480;
)

var (
        running      bool;
        textures     []gl.Texture = make([]gl.Texture, 3);     //存储纹理
        texturefiles [1]string;                                //纹理文件名
        light        bool;                                     // Display light? 光照是否打开
        blend        bool;                                    // Perform blending?  
        rotation     [2]float32;                               // X/Y rotation.  旋转
        speed        [2]float32;                               // X/Y speed.  旋转速度
        z            float32    = -5;                          // Depth into the scene. 深度
        ambient      []float32  = []float32{0.5, 0.5, 0.5, 1}; // ambient light colour. 环境光参数
        diffuse      []float32  = []float32{1, 1, 1, 1};       // diffuse light colour. 漫射光参数
        lightpos     []float32  = []float32{0, 0, 2, 1};       // Position of light source. 光源位置
        filter       int;                                      // Index of current texture to display. 滤波类型
)

func init() {
        texturefiles[0] = "/home/opt/Downloads/Glass.tga";
}

func main() {
        var err error;
        if err = glfw.Init(); err != nil {
                log.Fatalf("%v\n", err);
                return;
        }

        defer glfw.Terminate();

        if err = glfw.OpenWindow(Width, Height, 8, 8, 8, 8, 8, 0, glfw.Windowed); err != nil {
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
        glu.Perspective(45.0, float64(w)/float64(h), 0.1, 100.0);
        gl.MatrixMode(gl.MODELVIEW);
        gl.LoadIdentity();
}

func onKey(key, state int) {
        switch key {
        case glfw.KeyEsc:
                running = false;
        case 'L':
                if state == 1 {
                        if light = !light; !light {
                                gl.Disable(gl.LIGHTING);
                        } else {
                                gl.Enable(gl.LIGHTING);
                        }
                }
        case 'F':
                if state == 1 {
                        if filter++; filter >= len(textures) {
                                filter = 0;
                        }
                }
        case 'B': // B
                if state == 1 {
                        if blend = !blend; blend {  //切换混合选项的 TRUE / FALSE
                                gl.Enable(gl.BLEND); //打开混合
                                gl.Disable(gl.DEPTH_TEST); //关闭深度测试
                        } else {
                                gl.Disable(gl.BLEND); //关闭混合
                                gl.Enable(gl.DEPTH_TEST); //打开深度测试
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

        gl.ShadeModel(gl.SMOOTH);
        gl.ClearColor(0, 0, 0, 0);
        gl.ClearDepth(1);
        gl.DepthFunc(gl.LEQUAL);
        gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST);
        gl.Enable(gl.TEXTURE_2D);
        gl.Enable(gl.DEPTH_TEST);
        
        //alpha通道的值为 0.0意味着物体材质是完全透明的。1.0 则意味着完全不透明
        //以全亮度绘制此物体，并对其进行50%的alpha混合(半透明)。
        //当混合选项打开时，此物体将会产生50%的透明效果
        gl.Color4f(1, 1, 1, 0.5); //全亮度， 50% Alpha 混合
        gl.BlendFunc(gl.SRC_ALPHA, gl.ONE); //基于源象素alpha通道值的半透明混合函数

        gl.Lightfv(gl.LIGHT1, gl.AMBIENT, ambient);
        gl.Lightfv(gl.LIGHT1, gl.AMBIENT, diffuse);
        gl.Lightfv(gl.LIGHT1, gl.POSITION, lightpos);
        gl.Enable(gl.LIGHT1);
        return;
}

func destroyGL() {
        gl.DeleteTextures(textures);
        textures = nil;
}

func loadTextures() (err error) {
        gl.GenTextures(textures); //申请纹理

        // Texture 1
        textures[0].Bind(gl.TEXTURE_2D);//绑定纹理空间

        if !glfw.LoadTexture2D(texturefiles[0], 0) { //加载纹理
                return errors.New("Failed to load texture: " + texturefiles[0]);
        }
        //纹理缩放设置
        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST);
        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST);

        // Texture 2
        textures[1].Bind(gl.TEXTURE_2D);

        if !glfw.LoadTexture2D(texturefiles[0], 0) {
                return errors.New("Failed to load texture: " + texturefiles[0]);
        }

        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);

        // Texture 3
        textures[2].Bind(gl.TEXTURE_2D);

        if !glfw.LoadTexture2D(texturefiles[0], glfw.BuildMipmapsBit) {
                return errors.New("Failed to load texture: " + texturefiles[0]);
        }

        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
        gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR_MIPMAP_NEAREST);

        return;
}

func drawScene() {
        gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);
        gl.LoadIdentity();

        gl.Translatef(0, 0, z);

        gl.Rotatef(rotation[0], 1, 0, 0);
        gl.Rotatef(rotation[1], 0, 1, 0);

        rotation[0] += speed[0];
        rotation[1] += speed[1];

        textures[filter].Bind(gl.TEXTURE_2D);

        gl.Begin(gl.QUADS)
        // Front Face
        gl.Normal3f(0, 0, 1) // Normal Pointing Towards Viewer
        gl.TexCoord2f(0, 0)
        gl.Vertex3f(-1, -1, 1) // Point 1 (Front)
        gl.TexCoord2f(1, 0)
        gl.Vertex3f(1, -1, 1) // Point 2 (Front)
        gl.TexCoord2f(1, 1)
        gl.Vertex3f(1, 1, 1) // Point 3 (Front)
        gl.TexCoord2f(0, 1)
        gl.Vertex3f(-1, 1, 1) // Point 4 (Front)
        // Back Face
        gl.Normal3f(0, 0, -1) // Normal Pointing Away From Viewer
        gl.TexCoord2f(1, 0)
        gl.Vertex3f(-1, -1, -1) // Point 1 (Back)
        gl.TexCoord2f(1, 1)
        gl.Vertex3f(-1, 1, -1) // Point 2 (Back)
        gl.TexCoord2f(0, 1)
        gl.Vertex3f(1, 1, -1) // Point 3 (Back)
        gl.TexCoord2f(0, 0)
        gl.Vertex3f(1, -1, -1) // Point 4 (Back)
        // Top Face
        gl.Normal3f(0, 1, 0) // Normal Pointing Up
        gl.TexCoord2f(0, 1)
        gl.Vertex3f(-1, 1, -1) // Point 1 (Top)
        gl.TexCoord2f(0, 0)
        gl.Vertex3f(-1, 1, 1) // Point 2 (Top)
        gl.TexCoord2f(1, 0)
        gl.Vertex3f(1, 1, 1) // Point 3 (Top)
        gl.TexCoord2f(1, 1)
        gl.Vertex3f(1, 1, -1) // Point 4 (Top)
        // Bottom Face
        gl.Normal3f(0, -1, 0) // Normal Pointing Down
        gl.TexCoord2f(1, 1)
        gl.Vertex3f(-1, -1, -1) // Point 1 (Bottom)
        gl.TexCoord2f(0, 1)
        gl.Vertex3f(1, -1, -1) // Point 2 (Bottom)
        gl.TexCoord2f(0, 0)
        gl.Vertex3f(1, -1, 1) // Point 3 (Bottom)
        gl.TexCoord2f(1, 0)
        gl.Vertex3f(-1, -1, 1) // Point 4 (Bottom)
        // Right face
        gl.Normal3f(1, 0, 0) // Normal Pointing Right
        gl.TexCoord2f(1, 0)
        gl.Vertex3f(1, -1, -1) // Point 1 (Right)
        gl.TexCoord2f(1, 1)
        gl.Vertex3f(1, 1, -1) // Point 2 (Right)
        gl.TexCoord2f(0, 1)
        gl.Vertex3f(1, 1, 1) // Point 3 (Right)
        gl.TexCoord2f(0, 0)
        gl.Vertex3f(1, -1, 1) // Point 4 (Right)
        // Left Face
        gl.Normal3f(-1, 0, 0) // Normal Pointing Left
        gl.TexCoord2f(0, 0)
        gl.Vertex3f(-1, -1, -1) // Point 1 (Left)
        gl.TexCoord2f(1, 0)
        gl.Vertex3f(-1, -1, 1) // Point 2 (Left)
        gl.TexCoord2f(1, 1)
        gl.Vertex3f(-1, 1, 1) // Point 3 (Left)
        gl.TexCoord2f(0, 1)
        gl.Vertex3f(-1, 1, -1)
        gl.End()

        glfw.SwapBuffers();
}