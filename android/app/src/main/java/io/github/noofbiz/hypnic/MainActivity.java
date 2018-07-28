package io.github.noofbiz.hypnic;

import android.hardware.SensorEventListener;
import android.opengl.GLSurfaceView;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.content.Context;

import android.hardware.SensorEventListener;
import android.hardware.Sensor;
import android.hardware.SensorEvent;
import android.hardware.SensorManager;
import android.util.Log;

import androidglue.*;

public class MainActivity extends AppCompatActivity implements SensorEventListener {

    private SensorManager senSensorManager;
    private Sensor senAccelerometer;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        senSensorManager = (SensorManager) getSystemService(Context.SENSOR_SERVICE);
        try {
            senAccelerometer = senSensorManager.getDefaultSensor(Sensor.TYPE_GAME_ROTATION_VECTOR);
            boolean ispresent = senSensorManager.registerListener(this, senAccelerometer, SensorManager.SENSOR_DELAY_GAME);
            Androidglue.accelerometerPresent(ispresent);
        } catch (Exception e) {
            Log.println(Log.WARN, "exception", "problem loading accelerometer");
        }
    }

    private GLSurfaceView glSurfaceView() {
        return (GLSurfaceView)this.findViewById(R.id.glview);
    }

    @Override
    protected void onPause() {
        super.onPause();
        this.glSurfaceView().onPause();
        senSensorManager.unregisterListener(this);
    }

    @Override
    protected void onResume() {
        super.onResume();
        this.glSurfaceView().onResume();
        senSensorManager.registerListener(this, senAccelerometer, SensorManager.SENSOR_DELAY_GAME);
    }

    @Override
    public void onSensorChanged(SensorEvent event) {
        if (event.sensor.getType() == Sensor.TYPE_GAME_ROTATION_VECTOR) {
            Androidglue.accelerometerValue(event.values[0]);
        }
    }

    @Override
    public void onAccuracyChanged(Sensor sensor, int accuracy) {

    }
}
