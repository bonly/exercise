#include <iostream>
#include <string>

std::string trim(const std::string& str,
                 const std::string& whitespace = " \t")
{
    const auto strBegin = str.find_first_not_of(whitespace);
    if (strBegin == std::string::npos)
        return ""; // no content

    const auto strEnd = str.find_last_not_of(whitespace);
    const auto strRange = strEnd - strBegin + 1;

    return str.substr(strBegin, strRange);
}

std::string reduce(const std::string& str,
                   const std::string& fill = " ",
                   const std::string& whitespace = " \t")
{
    // trim first
    auto result = trim(str, whitespace);

    // replace sub ranges
    auto beginSpace = result.find_first_of(whitespace);
    while (beginSpace != std::string::npos)
    {
        const auto endSpace = result.find_first_not_of(whitespace, beginSpace);
        const auto range = endSpace - beginSpace;

        result.replace(beginSpace, range, fill);

        const auto newStart = beginSpace + fill.length();
        beginSpace = result.find_first_of(whitespace, newStart);
    }

    return result;
}

int main(void)
{
    const std::string foo = "    too much\t   \tspace\t\t\t  ";
    const std::string bar = "one\ntwo";

    std::cout << "[" << trim(foo) << "]" << std::endl;
    std::cout << "[" << reduce(foo) << "]" << std::endl;
    std::cout << "[" << reduce(foo, "-") << "]" << std::endl;

    std::cout << "[" << trim(bar) << "]" << std::endl;
}