void setup(){
  Serial.begin(115200);
  pinMode(2, INPUT);
}

void loop(){
  int switchStatus = digitalRead(2);
  Serial.println(switchStatus);
}