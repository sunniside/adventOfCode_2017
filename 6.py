memoryBanks = [4,10,4,1,8,4,9,14,5,1,14,15,0,15,3,5];
seenMemoryBanks = [];
cycles = 0;

def evaluateSeenMemoryBanks(currentMemoryBank):
    global cycles;
    global seenMemoryBanks;
    cycles += 1;
    for memBank in seenMemoryBanks:
        if(currentMemoryBank == memBank):
            seenMemoryBanks.append(list(currentMemoryBank));
            print cycles;
            for index in range(len(seenMemoryBanks)):
                if(seenMemoryBanks[index] == memBank):
                    print len(seenMemoryBanks) - 1 - index;
                    break;
            return False;

    seenMemoryBanks.append(list(currentMemoryBank));
    return True;

def redistribute():
    global memoryBanks;
    keepGoing = True;
    while keepGoing:
        index = 0;
        pointsToDistribute = 0;
        for index in range(len(memoryBanks)):
            if(memoryBanks[index] > pointsToDistribute):
                pointsToDistribute = memoryBanks[index];
                largestIndex = index;
        memoryBanks[largestIndex] = 0;
        index = largestIndex + 1;
        if(index >= len(memoryBanks)):
            index = 0;
        while pointsToDistribute > 0:
            memoryBanks[index] += 1;
            pointsToDistribute -= 1;
            index += 1;
            if(index >= len(memoryBanks)):
                index = 0;
        keepGoing = evaluateSeenMemoryBanks(memoryBanks);

seenMemoryBanks.append(list(memoryBanks));
redistribute();
