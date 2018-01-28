using UnityEngine;
using System.Collections;
using strange.extensions.context.impl;

namespace techappen{

public class MainBootstrap : ContextView {
	void Start () {
		context = new MainContext(this);
	}
};

}
