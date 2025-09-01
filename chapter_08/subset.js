function isSubset(array1, array2) {
  let largerArray;
  let smallerArray;

  if (array1.length > array2.length) {
    largerArray = array1;
    smallerArray = array2;
  } else {
    largerArray = array2;
    smallerArray = array1;
  }

  // Iterate through smaller array
  for (let i = 0; i < smallerArray.length; i++) {
    let foundMatch = false;

    for (let j = 0; j < largerArray.length; j++) {
      if (smallerArray[i] === largerArray[j]) {
        foundMatch = true;
        break;
      }
    }
    if (foundMatch === false) {
      return false;
    }
  }
  return true;
}

let arrLarge = [1, 3, 4, 5, 6];
let arrSmall = [3, 4];

console.log(isSubset(arrSmall, arrLarge));
