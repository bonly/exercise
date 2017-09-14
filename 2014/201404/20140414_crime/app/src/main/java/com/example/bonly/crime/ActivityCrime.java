package com.example.bonly.crime;

import android.support.v4.app.Fragment;
import android.support.v4.app.FragmentActivity;
import android.os.Bundle;
import android.support.v4.app.FragmentManager;

public class ActivityCrime extends FragmentActivity{

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_crime);

        FragmentManager fm = getSupportFragmentManager(); //找到全局的容器管理器
        Fragment frag = fm.findFragmentById(R.id.fragment_container); //找容器
        if (frag == null){
            frag = new CrimeFragment(); //建容器中的frag
            fm.beginTransaction()  //加frag到容器中
                    .add(R.id.fragment_container, frag)
                    .commit();
        }
    }
}
