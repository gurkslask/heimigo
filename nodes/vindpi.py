from readtemp import readTemp
import random
import time
from paho.mqtt import client as mqtt_client


broker = "192.168.1.171"
port = 1883
topic = "/mosquitto/SUN_GT2"
client_id = f'python-mqtt-{random.randint(0,1000)}'

def connect_mqtt():
    def on_connect(client, userdata, flags, rc):
        if rc == 0:
            print("Connected to MQTT Broker!")
        else:
            print("Failed to connect, return code %d\n", rc)
    # Set Connecting Client ID
    client = mqtt_client.Client(client_id)
    # client.username_pw_set(username, password)
    client.on_connect = on_connect
    client.connect(broker, port)
    return client

def publish(client):
    msg_count = 0
    old_temp = 0.0
    diff = 0.5
    while True:
        time.sleep(random.randint(1, 5))
        (err, temp) = readTemp('28-0000052378dd')
        if err:
            temp = 0,0
            print("Problem with reading temperature")
        else:
            print(f"Read temp {temp}, old temp: {old_temp}")
            
        
        msg = f'{temp}'
        # jjmsg = f'Messages: {temp}'
        # Only publish if difference in temp is greater than 0.5
        if old_temp + diff < temp or old_temp - diff > temp:
            old_temp = temp
            result = client.publish(topic, msg)

            status = result[0]
            if status == 0:
                print(f"Send {msg} to {topic}")
            else:
                print(f"Failed to send {msg} to {topic}")
        msg_count += 1


def run():
    client = connect_mqtt()
    client.loop_start()
    publish(client)

if __name__ == '__main__':
    run()