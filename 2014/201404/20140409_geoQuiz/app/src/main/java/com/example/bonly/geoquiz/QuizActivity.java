package com.example.bonly.geoquiz;

import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.support.design.widget.FloatingActionButton;
import android.support.design.widget.Snackbar;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.util.Log;
import android.view.View;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.Button;
import android.widget.TextView;
import android.widget.Toast;

import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.File;

import go.mypkg.Mypkg;

public class QuizActivity extends AppCompatActivity {
    private static final String TAG = "QuizActivity";
    private static final String KEY_INDEX = "index";
    private Button mTrueButton;
    private Button mFalseButton;
    private Button mNextButton;
    private TextView mQuestionTextView;
    private File mPath;
    private Button mCheatButton;
    private static int REQUEST_CODE_CHEAT = 0;

    private Question[] mQuestionBank = new Question[]{
            new Question(R.string.question_africa, false),
            new Question(R.string.question_oceans, true),
            new Question(R.string.question_americas, false),
            new Question(R.string.question_asia, true),
            new Question(R.string.question_mideast, false),
    };

    private int mCurrentIndex = 0;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        Log.d(TAG, "onCreate(Bundle) called");

        setContentView(R.layout.activity_quiz);
        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);

        mPath = getBaseContext().getExternalFilesDir(null);
        Log.d(TAG,mPath.toString());

        mQuestionTextView = (TextView) findViewById(R.id.question_text_view);
//        int question = mQuestionBank[mCurrentIndex].getTextResId();
//        mQuestionTextView.setText(question);

        mCheatButton = (Button) findViewById(R.id.cheat_button);
        mCheatButton.setOnClickListener(new View.OnClickListener(){
            @Override
            public void onClick(View v){
//                Intent it = new Intent(QuizActivity.this, CheatActivity.class);
//                startActivity(it);
                Intent it = new Intent(QuizActivity.this, CheatActivity.class);
                it.putExtra("Answer_is_true", mQuestionBank[mCurrentIndex].isAnswerTrue());
//                startActivity(it);

                startActivityForResult(it, REQUEST_CODE_CHEAT);
            }
        });

        mTrueButton = (Button) findViewById(R.id.true_button);
        mTrueButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v){
                /*
                Toast.makeText(QuizActivity.this,
                        R.string.correct_toast,
                        Toast.LENGTH_SHORT).show();
                        */
                checkAnswer(true);
            }
        });

        mFalseButton = (Button) findViewById(R.id.false_button);
        mFalseButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                /*
                Toast.makeText(QuizActivity.this,
                        R.string.incorrect_toast,
                        Toast.LENGTH_SHORT).show();
                */
                checkAnswer(false);
            }
        });

        mNextButton = (Button) findViewById(R.id.next_button);
        mNextButton.setOnClickListener(new View.OnClickListener(){
                @Override
                public void onClick(View v){
                    mCurrentIndex = (mCurrentIndex+1) % mQuestionBank.length;
//                    int question = mQuestionBank[mCurrentIndex].getTextResId();
//                    mQuestionTextView.setText(question);
                    updateQuestion();
                }
        });

        if (savedInstanceState != null){
            mCurrentIndex = savedInstanceState.getInt(KEY_INDEX, 0);
        }
        updateQuestion();

        FloatingActionButton fab = (FloatingActionButton) findViewById(R.id.fab);
        fab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
//                Process process = null;
//                DataOutputStream os = null;
//                DataInputStream is = null;
//                try{
//                    process =  Runtime.getRuntime().exec("/system/xbin/su");
//                    os = new DataOutputStream(process.getOutputStream());
//                    is = new DataInputStream(process.getInputStream());
//                    os.writeBytes ("/data/tmp/update_hosts" + "\n");
////                    os.writeBytes("exit \n");
//                    os.flush();
//                }catch (Exception e){
//                    Log.e("bonly", "Unexpected err: " + e.getMessage());
//                }finally{
//                    try {
//                        if (os != null) {
//                            os.close();
//                        }
//                        if (is != null) {
//                            is.close();
//                        }
//                        process.destroy();
//                    }catch (Exception e){
//                    }
//                }
                String ret = Mypkg.Update(mPath.toString());
//                String ret = "ok";
                Snackbar.make(view, ret, Snackbar.LENGTH_LONG)
                        .setAction("Action", null).show();
            }
        });
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data){
        if (resultCode != Activity.RESULT_OK){
            return;
        }
        if (requestCode == REQUEST_CODE_CHEAT){
            if (data == null){
                return;
            }
            if (data.getBooleanExtra("Cheated", false)) {
                String txt = "你已作弊";
                Toast.makeText(this, txt, Toast.LENGTH_SHORT).show();
            }
        }
    }

    private void updateQuestion(){
        int question = mQuestionBank[mCurrentIndex].getTextResId();
        mQuestionTextView.setText(question);
    }

    private void checkAnswer(boolean userPressedTrue){
        boolean answerIsTrue = mQuestionBank[mCurrentIndex].isAnswerTrue();

        int messageResId = 0;
        if (userPressedTrue == answerIsTrue){
            messageResId = R.string.correct_toast;
        }else{
            messageResId = R.string.incorrect_toast;
        }
        Toast.makeText(this, messageResId, Toast.LENGTH_SHORT).show();
    }
    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_quiz, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        //noinspection SimplifiableIfStatement
        if (id == R.id.action_settings) {
            return true;
        }

        return super.onOptionsItemSelected(item);
    }

    @Override
    public void onStart(){
        super.onStart();
        Log.d(TAG, "onStart() called", new Exception());
    }

    @Override
    public void onPause(){
        super.onPause();
        Log.d(TAG, "onPause() calles");
    }

    @Override
    public  void onStop(){
        super.onStop();
        Log.d(TAG, "onStop() calles");
    }

    @Override
    public void onDestroy(){
        super.onDestroy();
        Log.d(TAG, "onDestory() calles");
    }

    @Override
    public void onSaveInstanceState(Bundle savedInstanceState){
        super.onSaveInstanceState(savedInstanceState);
        Log.i(TAG, "onSaveInstanceState");
        savedInstanceState.putInt(KEY_INDEX, mCurrentIndex);
    }

}
