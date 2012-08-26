/** 所有的 Parser 类的基类，用于支持 [] 运算符，从而允许对匹配执行动作。*/

template<class DerivedParser>
class Parser
{
  public:
    ///两个辅助函数，用于转成实际对象的类型，后面全局运算符重载也会用到。
    DerivedParser& derived()
    {
      return static_cast<DerivedParser&> (*this);
    }

    const DerivedParser& derived() const
    {
      return static_cast<DerivedParser&> (*this);
    }

    template<class Action,class parser_with_action>
    parser_with_action operator[](Action const& actor) const
    {
      return parser_with_action(derived(), actor);
    }
};

/** 对任意 parser 基类对象调用 operator[] 产生一个类型支持 action 的对象
 * 要求 parser 基类也支持 parse 成员函数。由于 parser_with_action 依然是一个
 * parser，因此还可以再施加 [] 运算，对同一匹配执行多个动作。
 */
template <>
class parser_with_action : public Parser
{
  public:
    parser_with_action(Parser parser, Action action) :
      Parser(parser), m_action(action)
    {
    }

    template< class Iterator>
    bool parse(Iterator& begin, Iterator end) const
    {
      Iterator o = begin;
      if (Parser::parse(begin, end))
      {
        m_action(o, begin);
        return true;
      }
      return false;
    }
  private:
    Action m_action;
};

// 匹配单个字符的解析器，匹配的规则是与期望的值相等，这里出现了本文第一个
// parse 函数，本代码示例的 parse 函数接口如下：
// 参数：begin: 指向字符开始的迭代器，是个非 const 引用
// 参数：end: 指向字符的结尾后面的迭代器，用于检查输入是否结束。
// 返回 true 表示匹配成功，begin 会被移动，否则 begin 不变。
// 特别说明：spirit 中的 parse 函数的接口比本文所用的复杂，这里作了简化。

class char_parser: public Parser
{
  public:
    char_parser(char c) :
      m_ch(c)
    {
    }
    template
    bool parse(Iterator& begin, Iterator end) const
    {
      if (begin < end && *begin == m_ch)
      {
        ++begin;
        return true;
      }
      return false;
    }
  private:
    char m_ch;
};

// 顺序匹配的解析器，依次匹配连续的两个才算成功。
template
class sequence_parser :public parser< >
{
  public:
  sequence_parser(const FirstParser& p1, const SecondParser& p2) :
  m_p1(p1), m_p2(p2)
  {}

  template
  bool parse(Iterator& begin, Iterator end) const
  {
    if (m_p1.parse(begin, end) && m_p2.parse(begin, end))
    {
      return true;
    }
    return false;
  }
  private:
  const FirstParser& m_p1;
  const SecondParser& m_p2;
};

// 用运算符 >> 拼合两个解析器，产生顺序解析器。
// 这里的参数类型不能写成简单类型，而要写成 parser，是为了重载区分，
template
sequence_parser operator>>(parser const& a, parser const& b)
{
  return sequence_parser(a.derived(), b.derived());
}

// 二选一，任意一个匹配即为成功
template
class or_parser : public parser >
{
  public:
  or_parser(const FirstParser& p1, const SecondParser& p2) :
  m_p1(p1), m_p2(p2)
  {}
  template
  bool parse(Iterator& begin, Iterator end) const
  {
    if (m_p1.parse(begin, end) || m_p2.parse(begin, end))
    {
      return true;
    }
    return false;
  }
  private:
  const FirstParser& m_p1;
  const SecondParser& m_p2;
};

// 用 | 运算符声称“或”解析器
template
or_parser operator|(parser const& a, parser const& b)
{
  return or_parser(a.derived(), b.derived());
}

// 测试代码，我们写一个解析器来解析十六进制数的开头：0x 或者 0X。
//
// 我们的动作，仅仅是输出而已
struct out
{
    void operator()(const char* begin, const char* end) const
    {
      std::cout.write(begin, std::streamsize(end - begin)) << '\n';
    }
};

int main()
{
  const char* s = "0X";
  // 对应的正则表达式：0(x|X)
  bool b = (char_parser('0') >> (char_parser('x') | char_parser('X')))[out() // 我们的动作
        ].parse(s, s + 2);
}
