import csv
import sys

def common_ids_with_bug(left, right):
    """ 

    Given two lists dictionaries with the following shape:
        {
            'id':number,
            'name':string
        }
    return a tuple of all the IDs that are common to both lists

    Warning: This version has a bug!!!
        1. What is the bug?
        2. How would you solve it?
        3. How would you test it?

    """

    common = []
    for element in left:
        if element in right:
            common.append(element)

    common = [x['id'] for x in common]

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

    common_elements = common_ids_with_bug(left_input, right_input)
    print("IDs that are common to", sys.argv[1], "and", sys.argv[2])
    for element in common_elements:
        print("    ", element)
