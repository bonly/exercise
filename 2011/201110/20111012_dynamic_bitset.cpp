#include <boost/dynamic_bitset.hpp>
#include <boost/utility.hpp>
using namespace boost;

void test_dynamic_bitset()
{
        using namespace boost;

        // 1. 构造
        dynamic_bitset<> db1;                              // 空的dynamic_bitset
        dynamic_bitset<> db2(10);                          // 大小为10的dynamic_bitset
        dynamic_bitset<> db3('\0x16', BOOST_BINARY(10101)); // 
        dynamic_bitset<> db4(std::string("0101"));         // 字符串构造

        // 2. resize
        db1.resize(8, true);
        assert(db1.to_ulong() == BOOST_BINARY(11111111));
        db1.resize(5);
        assert(db1.to_ulong() == BOOST_BINARY(11111));
        db1.clear();
        assert(db1.empty() && db1.size() == 0);

        // 3. push_back
        // dynamic_bitset可以像vector那样使用push_back()向容器末尾(二制数头部)追加一个值
        dynamic_bitset<> db5(5, BOOST_BINARY(01010));
        assert(db5.to_ulong() == BOOST_BINARY(01010));
        db5.push_back(true);    // 添加二进制位1
        assert(db5.to_ulong() == BOOST_BINARY(101010));
        db5.push_back(false);   // 添加二进制位0
        assert(db5.to_ulong() == BOOST_BINARY(0101010));

        // 4. block
        // dynamic_bitset使用block来存储二进制位, 一个block就可以存储32个二进制位
        assert(dynamic_bitset<>(32).num_blocks() == 1);
        assert(dynamic_bitset<>(33).num_blocks() == 2);

        // 5. 位运算
        dynamic_bitset<> db6(4, BOOST_BINARY(1010));
        db6[0] &= 1;    // 按位与运算
        db6[1] ^= 1;    // 按位异或运算
        db6[2] |= 1;    // 按位或运算
        assert(db6.to_ulong() == BOOST_BINARY_UL(1100));
        
        // 6. 访问元素
        dynamic_bitset<> db7(4, BOOST_BINARY(1100));
        assert(!db7.test(0) && !db7.test(1));
        assert(db7.any() && !db6.none());
        assert(db7.count() == 2);
        assert(db7.set().to_ulong() == BOOST_BINARY(1111));
        assert(db7.reset().to_ulong() == BOOST_BINARY(0000));
        assert(db7.set(0, 1).to_ulong() == BOOST_BINARY(0001)); 
        assert(db7.set(2, 1).to_ulong() == BOOST_BINARY(0101));
        assert(db7.reset(0).to_ulong() == BOOST_BINARY(0100));

        assert(db7.find_first() == 2);
        assert(db7.find_next(2) == ULONG_MAX);  // 没有找到的情况
}

int main(){
  test_dynamic_bitset();
  return 0;
}

