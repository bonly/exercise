#include <memory>
#include <boost/optional.hpp>

namespace boost {
	template <class T>
		T* begin(boost::optional<T>& opt) noexcept
		{
			return opt.is_initialized() ?
				std::addressof(opt.get()) :
				nullptr;
		}

	template <class T>
		T* end(boost::optional<T>& opt) noexcept
		{
			return opt.is_initialized() ?
				std::addressof(opt.get()) + 1 :
				nullptr;
		}
}

#include <iostream>
int main()
{
	boost::optional<int> opt = 3;

	for (int& x : opt) {
		std::cout << x << std::endl;
	}
}
