using System;
using UnityEngine;
using strange.extensions.context.impl;

namespace techappen{
	
public class MainContext : SignalContext {
	public MainContext(MonoBehaviour contextView) : base(contextView){}
	
	protected override void mapBindings(){
		base.mapBindings();

		commandBinder.Bind<StartSignal>().To<MainStartupCommand>();
	}
}

};
