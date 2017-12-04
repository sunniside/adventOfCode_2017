inputFile = open("4.input", "r");

def partOne():
    validPhrases = 0;
    for line in inputFile:
        line = line.rstrip('\n');
        words = line.split(' ');
        before = len(words);
        after = len(set(words));
        if(before == after):
            validPhrases += 1;
    print validPhrases;

def isAnagram(s1,s2):
    s1_list = list(s1);
    s1_list.sort();
    s2_list = list(s2);
    s2_list.sort();
    return s1_list == s2_list;

def partTwo():
    validPhrases = 0;
    for line in inputFile:
        line = line.rstrip('\n');
        words = line.split(' ');
        anagram = False;
        for w in range (0, len(words)):
            for w2 in range (0, len(words)):
                if w == w2:
                    continue;
                if isAnagram(words[w], words[w2]):
                    anagram = True;
                    break;
        if not anagram:
            print words;
            validPhrases += 1;
    print validPhrases;

partOne();
partTwo();
