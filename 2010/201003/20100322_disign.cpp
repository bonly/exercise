#include <iostream>
using namespace std;
class Key //通货膨胀
{
    public:
        float key;
};

template<typename K>
class VO :public K//投资
{
    public:
        bool room()
        {
            bool result = (K::key>=2 && K::key<=5)?true:false;
            if(K::key > 5 && K::key < 10) result = false;
            if(K::key >=10) result = false;
            return result;
        }
        bool gold()
        {
            bool result = (K::key>=2 && K::key <= 5)?false:true;
            if(K::key > 5 && K::key < 10) result = false;
            if(K::key >= 10) result = false;
            return result;
        }
};

int main()
{
    VO<Key> vo;
    vo.key = 3.8;
    std::cerr << "room: " << vo.room() << endl;
    return 0;
}

