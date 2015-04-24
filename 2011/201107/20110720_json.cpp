std::locale oldLocale;
std::locale utf8Locale(oldLocale,
new boost::program_options::detail::utf8_codecvt_facet());

std::wistringstream jsonIStream;
jsonIStream.str((wchar_t*)_bstr_t(content->c_str()));

boost::property_tree::wptree ptParse;
boost::property_tree::json_parser::read_json(jsonIStream, ptParse);

// method 1:
//BOOST_AUTO(child, ptParse.get_child(L""));
//for (BOOST_AUTO(pos, child.begin()); pos != child.end(); ++pos)
// method 2:
//for (BOOST_AUTO(pos, ptParse.begin()); pos != ptParse.end(); ++pos)
// method 3:
BOOST_FOREACH(boost::property_tree::wptree::value_type &v, ptParse.get_child(L""))
{
USES_CONVERSION;
cout << W2A(v.second.get(L"k1").c_str()) << endl;
cout << W2A(v.second.get(L"k2").c_str()) << endl;
}
