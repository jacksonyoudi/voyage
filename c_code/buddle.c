/*************************************************************************
	> File Name: buddle.c
	> Author: 
	> Mail: 
	> Created Time: 日  8/ 7 21:33:11 2022
 ************************************************************************/

#include<stdio.h>

void buddle(int arr[], int n) {
    int i;
    int temp;
    for (i = 0;i<n-1; i++) {
        if (arr[i] > arr[i+1]) {
            temp = arr[i];
            arr[i] = arr[i+1];
            arr[i+1] = temp;
        }
}}

void buddleSort(int arr[], int n) {
    int i;
    for (i = n;i++;i--) {
        buddle(arr, i);
    }
}

void main() {
    int arr[] = {6,7,1,2,10,8,6,0};
    int i;
    buddleSort(arr, 8);
    for (i=0;i<8;i++) {
        printf("%d", i);

}
}
