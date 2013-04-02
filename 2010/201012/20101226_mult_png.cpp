/*
混合图像是用于创建强大且迷人的用户界面的重要方法之一。这是一篇向导，用于展示如何使用Nana库来混合图像并显示混合的结果。

在开始这个向导之前，我们需要准备两张用于混合的图片。

image01

image02

这两张图片是PNG格式的
*/
#include <nana/gui/wvl.hpp>
#include <nana/gui/drawing.hpp>
#include <nana/gui/timer.hpp>

class tsform: public nana::gui::form
{
public:
	//因为我们准备的图片大小是 450*300,
	//这个窗口创建的大小与图片保持一致.
	tsform()
		:nana::gui::form(nana::gui::API::make_center(450, 300)), fade_rate_(0.1), delta_(0.1)
	{
		//打开图像文件.
		image1_.open(STR("image01.png"));
		image2_.open(STR("image02.png"));
		timer_.make_tick(nana::make_fun(*this, &tsform::_m_blend));
		timer_.interval(10);
	}
private:
	void _m_blend()
	{
		fade_rate_ += delta_;
		if(fade_rate_ > 1)
		{
			fade_rate_ = 1;
			delta_ = -0.01;
		}
		else if(fade_rate_ < 0)
		{
			fade_rate_ = 0;
			delta_ = 0.01;
		}
            
		//在混合之前，需要将图像拷贝到graphics对象中
		nana::paint::graphics graph1(450, 300);
		image1_.paste(graph1, 0, 0);

		nana::paint::graphics graph2(450, 300);
		image2_.paste(graph2, 0, 0);

		//执行混合，结果保存在graph2对象中.
		graph1.blend(graph2, 0, 0, fade_rate_);

		//显示结果，将graph2中的内容绘制到窗口上.
		nana::gui::drawing drawing(*this);
		drawing.clear();
		drawing.bitblt(0, 0, 450, 300, graph2, 0, 0);
		drawing.update();
	}
private:
	double    fade_rate_;
	double    delta_;
	nana::gui::timer    timer_;
	nana::paint::image image1_, image2_;
};

int main()
{
	tsform fm;
	fm.show();
	nana::gui::exec();
}

