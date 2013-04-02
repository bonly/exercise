#include <nana/gui/wvl.hpp>
#include <nana/gui/widgets/button.hpp>
#include <nana/gui/widgets/progress.hpp>

class example
    : public nana::gui::form
{
public:
    example()
    {
        btn_start_.create(*this, nana::rectangle(10, 10, 100, 20));
        btn_start_.caption(STR("Start"));
        btn_start_.make_event<nana::gui::events::click>(*this, &example::_m_start);
        btn_cancel_.create(*this, nana::rectangle(120, 10, 100, 20));
        btn_cancel_.caption(STR("Cancel"));
        btn_cancel_.make_event<nana::gui::events::click>(*this, &example::_m_cancel);
        prog_.create(*this, nana::rectangle(10, 40, 280, 20));
    }

private:
    void _m_start()
    {
        working_ = true;
        btn_start_.enabled(false);
        prog_.amount(100);
        for(int i = 0; i < 100 && working_; ++i)
        {
            nana::system::sleep(1000); //a long-running simulation
            prog_.value(i + 1);
        }
        btn_start_.enabled(true);
    }

    void _m_cancel()
    {
        working_ = false;
    }
private:
    bool working_;
    nana::gui::button btn_start_;
    nana::gui::button btn_cancel_;
    nana::gui::progress prog_;
};

int main()
{
    example ex;
    ex.show();
    nana::gui::exec();
    return 0;
} 