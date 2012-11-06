#include <vector>
#include <map>
#include <boost/shared_ptr.hpp>
#include <boost/make_shared.hpp>
#include <boost/foreach.hpp>
using namespace std;
using namespace boost;

class Base
{
    public:
        int  base;
        virtual void print()
        { 
            clog << "Base" << endl;
            clog << "base: " << base << endl;
        }
};
typedef boost::shared_ptr<Base> base_ptr;

class Inv : public Base
{
    public:
        int inv;
        int other;
        base_ptr decoration;
        virtual void print()
        { 
            clog << "Inv" << endl;
            clog << "base: " << base 
                 << "\ninv: " << inv
                 << "\nother: " << other << endl;
            if (decoration!=0)
               decoration->print();
        }
};
typedef boost::shared_ptr<Inv> inv_ptr;

int newbase(vector<shared_ptr<Base> > &b)
{
     shared_ptr<Base> nb = boost::make_shared<Base>();
     nb->base = 30;
     b.push_back (nb);
     return 0;
}

int newInv(vector<base_ptr> &b)
{
     inv_ptr nb = boost::make_shared<Inv>();
     nb->base = 40;
     nb->inv = 50;
     nb->other = 60;
     b.push_back (nb);
     return 0;
}

int modInv(base_ptr &b)
{
    inv_ptr inv = dynamic_pointer_cast<Inv>(b);
    //static_pointer_cast会导致失去动态性支持
    inv->base = 11;
    inv->inv = 12;
    inv->other = 13;
    inv_ptr dec = make_shared<Inv>();
    dec->base = 14;
    dec->inv = 15;
    dec->other =16;
    inv->decoration = dec;
    dec->decoration = make_shared<Base>();
    dec->decoration->base = 17;
    return 0;
}

int main()
{
    vector<base_ptr> vb;
    newbase(vb);
    newInv(vb);
    modInv(vb[1]);
    BOOST_FOREACH (base_ptr p, vb)
    {
       p->print();
    }


    //vector<inv_ptr> vi;
    //newbase(vi); // 参数不匹配
    return 0;
}
