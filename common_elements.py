import csv
import sys

def common_ids(left, right):
    """ 

    Given two lists of dictionaries with the following shape:
        {
            'id':number,
            'name':string
        }
    return a list of all the IDs that are common to both lists
    """

    common = []
    for l in left:
        for r in right:
            if l == r:
                common.append(l)

    return common


if __name__ == '__main__':

    if len(sys.argv) == 3:
        with open(sys.argv[1]) as fp:
            left_input = list(csv.DictReader(fp))
        with open(sys.argv[2]) as fp:
            right_input = list(csv.DictReader(fp))
    else:
        print("USAGE:", sys.argv[0], "<left.csv> <right.csv>", file=sys.stderr)
        sys.exit(1)

    common_elements = common_ids(left_input, right_input)
    print("IDs that are common to", sys.argv[1], "and", sys.argv[2])
    for element in common_elements:
        print("    ", element)
