#define STATIC_CHECK(expr) {char unnamed[(expr)?1:0];}

template <class To, class From>
To safe_reinterpret_cast(From from)
{
    STATIC_CHECK(sizeof(From)<=sizeof(To));
    return reinterpret_cast<To>(from);
}


int main()
{
  char k='a';
  char *p = &k;
  char m = safe_reinterpret_cast<char>(p);

  return 0;
}

