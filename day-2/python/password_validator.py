import re


class record:
    def __init__(self, input):
        regex = "(^\\d*)-(\\d*)\\s([a-zA-Z]):\\s([a-zA-Z].*)"
        s = re.findall(regex, input)[0]

        self.lower = int(s[0])
        self.upper = int(s[1])
        self.char = s[2]
        self.password = s[3]
        self.char_count = s[3].count(s[2])


def valid(input):

    valid = 0

    for i in input:
        r = record(i)

        if r.char_count >= r.lower and r.char_count <= r.upper:
            valid += 1

    return valid


def valid_strict(input):

    valid = 0

    for i in input:
        r = record(i)

        if len(r.password) < r.upper:
            continue

        if (r.password[r.lower-1] == r.char) != (r.password[r.upper-1] == r.char):
            valid += 1

    return valid
