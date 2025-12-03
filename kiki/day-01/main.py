from utils.input_helper import readInput
import os

def problemListToRotations(problem):
    rotations = [] 
    for line in problem:
        if line[0] == 'R':
            rotations.append(int(line[1:]))
        else:
            rotations.append(int(line[1:]) * -1)
    return rotations


def main():
    project_root_path = os.getcwd()
    path = f'{project_root_path}/day-01/input.txt'
    input = problemListToRotations(readInput(path)) # input is an arry of ints
    rotation = 50
    answer = 0

    for index in range(len(input)):
        rotation = (rotation + input[index])%100
        if rotation == 0:
            answer += 1

    print(answer)
        
if __name__ == '__main__':
    main()