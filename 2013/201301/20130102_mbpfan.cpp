/*
 * auth: bonly
 */

#include <math.h>
#include <stdio.h>
#include <unistd.h>

/* lazy min/max... */
#define min(a,b) a < b ? a : b
#define max(a,b) a > b ? a : b

/* basic fan speed parameters */
unsigned short min_fan_speed=2000;
unsigned short max_fan_speed=6200;

/* temperature thresholds 
 * low_temp - temperature below which fan speed will be at minimum
 * high_temp - fan will increase speed when higher than this temperature
 * max_temp - fan will run at full speed above this temperature */
unsigned short low_temp=55;
unsigned short high_temp=65;
unsigned short max_temp=80;

/* temperature polling interval */
unsigned short polling_interval=10;


/* Controls the speed of the fan */
void set_fan_speed(unsigned short speed)
{
  FILE *file;

  //file=fopen("/sys/devices/platform/applesmc.768/fan1_output", "w");
  file=fopen("/sys/devices/platform/applesmc.768/fan1_min", "w");
  fprintf(file, "%d", speed);
  fclose(file);
}


/* Takes "manual" control of fan */
void prepare_fan()
{
  FILE *file;

  file=fopen("/sys/devices/platform/applesmc.768/fan1_manual", "w");
  fprintf(file, "%d", 1);
  fclose(file);
}


/* Returns average CPU temp in degrees (ceiling) */
unsigned short get_temp()
{
  FILE *file;
  unsigned short temp; 
  unsigned int t0, t1;
  
  //file=fopen("/sys/devices/platform/coretemp.0/temp1_input", "r");
  file=fopen("/sys/devices/platform/coretemp.0/hwmon/hwmon1/temp1_input", "r");
  fscanf(file, "%d", &t0);
  fclose(file);
  
  //file=fopen("/sys/devices/platform/coretemp.1/temp1_input", "r");
  file=fopen("/sys/devices/platform/coretemp.0/hwmon/hwmon1/temp2_input", "r");
  fscanf(file, "%d", &t1);
  fclose(file);
  
  temp = (unsigned short)(ceil((float)(t0 + t1) / 2000.));
  return temp;
}


int main()
{
  unsigned short old_temp, new_temp, fan_speed, steps;
  short temp_change;
  float step_up, step_down;
  
  //prepare_fan();
  
  /* assume running on boot so set fan speed to minimum */
  new_temp = get_temp();
  fan_speed = 2000;
  set_fan_speed(fan_speed);
  sleep(polling_interval);

  step_up = (float)(max_fan_speed - min_fan_speed) / 
            (float)((max_temp - high_temp) * (max_temp - high_temp + 1) / 2);
  
  step_down = (float)(max_fan_speed - min_fan_speed) / 
              (float)((max_temp - low_temp) * (max_temp - low_temp + 1) / 2);
     
  while(1)
  {
    old_temp = new_temp;
    new_temp = get_temp();
    
    if(new_temp >= max_temp && fan_speed != max_fan_speed) {
      fan_speed = max_fan_speed;
    }
    
    if(new_temp <= low_temp && fan_speed != min_fan_speed) {
      fan_speed = min_fan_speed;
    }
    
    temp_change = new_temp - old_temp;
    
    if(temp_change > 0 && new_temp > high_temp && new_temp < max_temp) {
      steps = (new_temp - high_temp) * (new_temp - high_temp + 1) / 2;
      fan_speed = max(fan_speed, ceil(min_fan_speed + steps * step_up));
    }

    if(temp_change < 0 && new_temp > low_temp && new_temp < max_temp) {
      steps = (max_temp - new_temp) * (max_temp - new_temp + 1) / 2;
      fan_speed = min(fan_speed, floor(max_fan_speed - steps * step_down));
    }

    set_fan_speed(fan_speed);
    sleep(polling_interval);    
  }
  
  return 0;
}
