# CHALLENGE 1
def triangular_triplet(arr):
    array_length = len(arr)
    if array_length<3 :
        return 0
    arr.sort()
    for i in range(array_length/2):
        if arr[i] + arr[i+1] > arr[i+2]:
            return 1
    

# CHALLENGE 2
def nested(s):
    store = []
    for i in s:
        if i in ['[', '{', '(']:
            store.append(i)
        elif i in [']', '}', ')']:
            if len(store)== 0:
                return 0
            y = store.pop()
            if i == ']' and y!= '[':
                return 0
            elif i == '}' and y!= '{':
                return 0
            elif i == ')' and y!= '(':
                return 0
    return 1


# CHALLENG3
def dominator(arr):
    obj = {}
    temp = []
    for i in arr:
        occurrences = arr.count(i)
        obj[i] = occurrences
    max_key = max(obj, key = obj.get)
    max_val = max(obj.values())
    if max_val > len(arr)/2 :
        for i in range(len(arr)):
            if arr[i] == max_key:              
                temp.append(i)
        return temp
    return -1

 print(triangular_triplet([10, 2, 5, 1, 8, 20])) 
print(dominator([3,4, 3, 2, 3, -1, 3, 3]))
print(nested("{[()()]}"))
print(nested("([)()]"))