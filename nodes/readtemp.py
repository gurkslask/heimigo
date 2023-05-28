import os
def readTemp(sensor: str) -> (bool, float):
    dir = "/sys/bus/w1/devices/"
    filename = 'w1_slave'
    path = os.path.join(dir, sensor, filename)
    with open(path, 'r') as f:
        data = f.read()
        if 'YES' in data:
            temp_pos = data.find('t=')  # Look for t, this is where the temperature is
            data = data[temp_pos + 2:]  # remove 't='
            data = float(data) / 1000  # insert comma,
            return (False, data)
        else:
            return (True, 0.0)


if __name__ == '__main__':
    readTemp('28-0000052361be')
