#include <stdio.h>

#include <iostream>
#include <sstream>
#include <string>
#include <locale>

#include "boost/property_tree/ptree.hpp"
#include "boost/property_tree/json_parser.hpp"
#include "boost/property_tree/xml_parser.hpp"

#include "boost/program_options/detail/convert.hpp"
#include "boost/program_options/detail/utf8_codecvt_facet.hpp"

int main(int argc, char **argv)
{
    /* The data format
     * <root>
     *  <num>1</num>
     *  <str>Test</str>
     * </root>
     */
    /* test UTF-8 format */
    try
    {
        /* create boost utf8 codecvt */
        std::locale oldLocale;
        std::locale utf8Locale(oldLocale,
            new boost::program_options::detail::utf8_codecvt_facet());
        std::wcout.imbue(utf8Locale);

        /* create the wptree for save the UTF-8 data */
        boost::property_tree::wptree datum;
        datum.put(L"root.num", 100);
        datum.put(L"root.str", L"wstring");

        /* output XML string */
        std::wostringstream xmlOutputStream;
        xmlOutputStream.imbue(utf8Locale);
        boost::property_tree::xml_parser::write_xml(xmlOutputStream,
            datum);
        std::wcout << L"XML format:" << std::endl;
        std::wcout << xmlOutputStream.str() << std::endl;

        /* output JSON string */
        std::wostringstream jsonOutputStream;
        jsonOutputStream.imbue(utf8Locale);
        boost::property_tree::json_parser::write_json(jsonOutputStream,
            datum);
        std::wcout << L"JSON format:" << std::endl;
        std::wcout << jsonOutputStream.str() << std::endl;

        /* read datum from JSON stream */
        boost::property_tree::wptree wptParse;
        std::wistringstream jsonIStream;
        jsonIStream.imbue(utf8Locale);
        jsonIStream.str(jsonOutputStream.str());
        boost::property_tree::json_parser::read_json(jsonIStream,
            wptParse);
        int num = wptParse.get<int>(L"root.num");
        std::wstring wstrVal = wptParse.get<std::wstring>(L"root.str");
        std::wcout << L"Num=" << std::dec << num
            << L" Str=" << wstrVal << std::endl << std::endl;
    }
    catch (...)
    {
        printf("create boost::property_tree::wptree failed\n");
    }

    return 0;
}

