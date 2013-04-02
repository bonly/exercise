/**
在某些情况下，耗时工作不能被取消，也不知道当前的进度。程序通常会用一直滚动的进度条表示正在处理。
*/
#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/button.hpp>
#include <nana/gui/widgets/progress.hpp>
#include <nana/threads/pool.hpp>

class example
    : public nana::gui::form
{
public:
    example()
    {
        using namespace nana::gui;

        btn_start_.create(*this, nana::rectangle(10, 10, 100, 20));
        btn_start_.caption(STR("Start"));
        btn_start_.make_event<events::click>(nana::threads::pool_push(pool_, *this, &example::_m_start));
        btn_start_.make_event<events::click>(nana::threads::pool_push(pool_, *this, &example::_m_ui_update));

        prog_.create(*this, nana::rectangle(10, 40, 280, 20));
        //prog_.style(false);
        prog_.unknown(false);

        this->make_event<events::unload>(*this, &example::_m_cancel);
    }

private:
    void _m_start()
    {
        btn_start_.enabled(false);
        nana::system::sleep(10000); //a blocking simulation
        btn_start_.enabled(true);
    }

    void _m_ui_update()
    {
        while(btn_start_.enabled() == false)
        {
            prog_.inc();
            nana::system::sleep(100);
        }
    }

    void _m_cancel(const nana::gui::eventinfo& ei)
    {
        if(false == btn_start_.enabled())
            ei.unload.cancel = true;
    }
private:
    nana::gui::button btn_start_;
    nana::gui::progress prog_;
    nana::threads::pool pool_;
};

int main()
{
    example ex;
    ex.show();
    nana::gui::exec();
    return 0;
} 