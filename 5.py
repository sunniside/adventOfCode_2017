inputFile = open("5.input", "r");

offsets = [];
for line in inputFile:
    line = line.rstrip('\n');
    offsets.append(int(line));

def partOne(offsetIndex):
    offsets[offsetIndex] += 1;

def partTwo(jumps, offsetIndex):
    if(jumps >= 3):
        offsets[offsetIndex] -= 1;
    else:
        offsets[offsetIndex] += 1;

def jump(offsetIndex, steps):
    while offsetIndex < len(offsets):
        jumps = offsets[offsetIndex];
        #partOne(offsetIndex);
        partTwo(jumps, offsetIndex);
        offsetIndex = offsetIndex + jumps;
        steps += 1;
    print steps;

jump(0, 0);
