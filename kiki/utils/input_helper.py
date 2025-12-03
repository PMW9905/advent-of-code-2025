def readInput(path):
    problem = []
    with open(path, "r") as file:
         for line in file:
            problem.append(line.strip()) # .strip() to remove newline characters
    return problem


    # with open(...) as file
