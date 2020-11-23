import sys, os


import sys

def main():
    complete_str = '' 
    start_str = "---\nvms:\n"
    complete_str = start_str
    # print command line arguments
    with open("keys.txt") as f:
        content = f.readlines()
    # you may also want to remove whitespace characters like `\n` at the end of each line
    content = [x.strip() for x in content] 
    for arg in sys.argv[1:]:
        for i in range(int(arg)):
            complete_str = complete_str + '  - name: vm'+str(i)+'\n    partitions: '+str(arg)+'\n    p_key: '+content[i]+'\n    id: '+str(i)+'\n'
            
    print(complete_str)
    with open("list.yaml", "w") as text_file:
        text_file.write(complete_str)

if __name__ == "__main__":
    main()
