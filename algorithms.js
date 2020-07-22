//CHALLENGE1
function triangular_triplet(arr) {
    const len = arr.length;
    if (len < 3) return 0;
    arr.sort((a, b) => a - b);
    let result;
    for (let i = 0; i < len - 2; i++) {
      if (arr[i] + arr[i + 1] > arr[i + 2]) {
        result = 1;
        break;
      }
      result = 0;
    }
    return result;
  }

//CHALLENGE 2
const nested = (s) => {
    store = []
   s.split("").map(char => {
       if(['[', '{', '('].includes(char) ) {
        store.push(char)
       }
       else if([']', '}', ')'].includes(char)){
           if (store.length == 0) return 0
           x = store.pop()
           if (char== ']' && x!= '['){
                return 0;
           }else if (char == '}' && x!= '{'){
               return 0;
           }
           else if(char == ')' && x!= '('){
               return 0;
           }    
       }
    return 1      
   });  
}

//CHALLENGE 3
const findDominator = (arr) => {
    const arrTransform = arr.reduce((acc, val, index) => {
        if(acc[val]) acc[val].push(index);
        else acc[val] = [index];
        return acc
    }, {});
    let dominatorCollection;
    const isDominant = Object.values(arrTransform).some((value) => {
        dominatorCollection = value;
        return value.length > arr.length /2;
    });
    return isDominant? dominatorCollection : -1;
}

console.log(triangular_triplet([10, 2, 5, 1, 8, 20]))
console.log(nested("{[()()]}"))
console.log(findDominator([3, 4, 2, 2, -1, 3]))

