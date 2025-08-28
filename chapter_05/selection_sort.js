function selectionSort(array) {
  for (let i = 0; i < array.length; i++) {
    let lowestAt = i;
    for (let j = i; j < array.length; j++) {
      if (array[j] < array[lowestAt]) {
        lowestAt = j;
      }
    }

    if (lowestAt != i) {
      let temp = array[i];
      array[i] = array[lowestAt];
      array[lowestAt] = temp;
    }
  }
  return array;
}

let arr = [2, 5, 0, 1, 8, 3, 9];
console.log(selectionSort(arr));
