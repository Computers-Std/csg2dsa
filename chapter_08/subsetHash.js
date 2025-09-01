function isSubset(arr1, arr2) {
  let largeArray;
  let smallArray;
  let hashTable = {};

  // Determine which array is smaller
  if (arr1.length > arr2.length) {
    largeArray = arr1;
    smallArray = arr2;
  } else {
    largeArray = arr2;
    smallArray = arr1;
  }

  // Store all items from largeArray inside hashTable
  for (const value of largeArray) {
    hashTable[value] = true;
  }

  // Interate through each item in smallArray and return false
  // if we encounter an item not inside hashTable
  for (const value of smallArray) {
    if (!hashTable[value]) {
      return false;
    }
  }
  return true;
}
