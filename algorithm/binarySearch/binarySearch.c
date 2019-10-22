#include <stdio.h>

int binarySearch(int* narray, int num)
{
    int array[5] = {1,3,5,7,9};
    int left = 0;
    int right = sizeof(array)/sizeof(int)-1;
    printf("left=%d, right=%d\n", left, right);

    while(left < right) 
    {
        printf("left=%d, right=%d\n", left, right);
        int mid = left + (right - left +1) / 2;
        if (array[mid] == num)
        {
            return mid;
        } else if(array[mid] < num) {
            left = mid + 1;
        } else if(array[mid] > num) {
            right = mid -1;
        }
    }
    return -1;
}

int main()
{
    int array[] = {1,3,5,7,9};
    int rst = binarySearch(array, 3);
    printf("rst=%d\n", rst);
    return 0;
}