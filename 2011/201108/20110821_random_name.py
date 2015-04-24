import random
parts = {}

with open('parts.txt', 'r') as f:
    currentList = []
    for line in f.readlines():
        line = line.strip()
        if line.startswith('[') and line.endswith(']'):
            currentList = []
            parts[line[1:-1]] = currentList
        else:
            currentList.append(line.strip())


for i in xrange(10):    
    print ''.join(random.choice(parts[partName]) for partName in sorted(parts))
