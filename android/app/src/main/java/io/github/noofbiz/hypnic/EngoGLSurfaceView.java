package io.github.noofbiz.hypnic;

import android.content.Context;
import android.content.res.Resources;
import android.opengl.GLSurfaceView;
import android.util.Log;
import android.util.AttributeSet;
import android.view.MotionEvent;

import javax.microedition.khronos.egl.EGLConfig;
import javax.microedition.khronos.opengles.GL10;

import java.io.FileInputStream;
import java.io.FileOutputStream;

import androidglue.*;

public class EngoGLSurfaceView extends GLSurfaceView {

    private class EngoRenderer implements Renderer {

        private boolean mErrored;

        @Override
        public void onDrawFrame(GL10 gl) {
            if (mErrored) {
                return;
            }
            try {
                Androidglue.update();
            } catch (Exception e) {
                Log.e("Go Error", e.toString());
                mErrored = true;
            }
        }

        @Override
        public void onSurfaceCreated(GL10 gl, EGLConfig config) {
        }

        @Override
        public void onSurfaceChanged(GL10 gl, int width, int height) {
        }
    }

    public EngoGLSurfaceView(Context context) {
        super(context);
        initialize();
    }

    public EngoGLSurfaceView(Context context, AttributeSet attrs) {
        super(context, attrs);
        initialize();
    }

    private void initialize() {
        setEGLContextClientVersion(2);
        setEGLConfigChooser(8, 8, 8, 8, 0, 0);
        setRenderer(new EngoRenderer());
    }

    @Override
    public void onLayout(boolean changed, int left, int top, int right, int bottom) {
        super.onLayout(changed, left, top, right, bottom);

        try {
            if (!Androidglue.isRunning()) {
                try {
                    FileInputStream file = getContext().openFileInput("opts.json");
                    byte[] b = new byte[file.available()];
                    file.read(b);
                    Androidglue.loadFromJava(b);
                    file.close();
                } catch (java.io.FileNotFoundException e) {
                    Androidglue.loadFromJava(new byte[0]);
                } catch (java.io.IOException e) {
                    Androidglue.loadFromJava(new byte[0]);
                }
                Androidglue.start(Resources.getSystem().getDisplayMetrics().widthPixels, Resources.getSystem().getDisplayMetrics().heightPixels);
            }
        } catch (Exception e) {
            Log.e("Go Error", e.toString());
        }
    }

    @Override
    public boolean onTouchEvent(MotionEvent e) {
        for (int i = 0; i < e.getPointerCount(); i++) {
            Androidglue.touch((int)e.getX(i), (int)e.getY(i), i, e.getActionMasked());
        }
        if (Androidglue.needsSaving()) {
            byte[] fileContents = Androidglue.theOptionsBytes();
            FileOutputStream outputStream;
            try {
                outputStream = getContext().openFileOutput("opts.json",Context.MODE_PRIVATE);
                outputStream.write(fileContents);
                outputStream.close();
            } catch (Exception ex) {
                Log.println(Log.WARN, "basdfa", "Asfdsdf");
            }
        }
        return true;
    }

    @Override
    public void onPause() {
        try {
            if (Androidglue.isRunning()) {
                Androidglue.stop();
            }
        } catch (Exception e) {
            Log.e("Go Error", e.toString());
        }
    }
}