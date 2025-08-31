function isPalindrome(string) {
    let leftIndex = 0;
    let rightIndex = string.length - 1;

    while (leftIndex < string.length / 2) {
        if (string[leftIndex] !== string[rightIndex]) {
            return false;
        }

        leftIndex++;
        rightIndex--;
    }
    return true;
}
