package com.example.bonly.crime;

import android.os.Bundle;
import android.support.v4.app.Fragment;
import android.support.v4.app.FragmentActivity;
import android.support.v4.app.FragmentManager;

/**
 * Created by bonly on 16-5-14.
 */
public abstract class SingleFragmentActivity extends FragmentActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_fragment); //指向xml里的空activity布局,指的是文件名

        FragmentManager fm = getSupportFragmentManager(); //找到全局的容器管理器
        Fragment frag = fm.findFragmentById(R.id.fragment_container); //找容器，指的是id名
        if (frag == null){
            frag = createFragment(); //建容器中的子类frag
            fm.beginTransaction()  //加frag到容器中
                    .add(R.id.fragment_container, frag)
                    .commit();
        }
    }
    protected abstract Fragment createFragment();
}
