package com.example.bonly.crime;

import android.support.v4.app.Fragment;

/**
 * Created by bonly on 16-5-14.
 */
public class CrimeListActivity extends SingleFragmentActivity{
    @Override
    protected Fragment createFragment() {
        return new CrimeListFragment();
    }
}
