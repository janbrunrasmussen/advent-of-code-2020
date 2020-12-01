
def checksum2(input):
    for n1 in input:
        for n2 in input:
            if n1 + n2 == 2020:
                return n1 * n2

    return []

def checksum3(input):
    for n1 in input:
        for n2 in input:
            for n3 in input:
                if n1 + n2 + n3 == 2020:
                    return n1 * n2 * n3
                    

    return []
