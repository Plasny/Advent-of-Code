#include <math.h>
#include <stdio.h>
#include <stdlib.h>

int is_number(char c) { 
    return c >= '0' && c <= '9'; 
}

int check_number(char *arr, int size, int y, int *x) {
  char c = arr[y * size + *x];
  int num = c - '0';

  int i_l = 0;
  while (1) {
    i_l++;
    if (*x - i_l < -1) {
      break;
    }

    c = arr[y * size + *x - i_l];
    if (!is_number(c)) {
      break;
    }

    num = (c - '0') * pow(10, i_l) + num;
  }

  int i_r = 0;
  while (1) {
    *x += 1;
    if (*x >= size) {
      break;
    }

    c = arr[y * size + *x];
    if (!is_number(c)) {
      break;
    }

    num = num * 10 + (c - '0');
  }

  return num;
}

int check_around(char *arr, int size, int y, int x) {
  int sum = 0;
  for (int i = y - 1; i <= y + 1; i++) {
    if (i < 0 || i >= size) {
      continue;
    }

    for (int j = x - 1; j <= x + 1; j++) {
      if (j < 0 || j >= size || (i == y && j == x)) {
        continue;
      }

      if (is_number(arr[i * size + j])) {
        sum += check_number(arr, size, i, &j);
      }
    }
  }

  return sum;
};

int main(int argc, char *argv[]) {
  FILE *f;
  size_t len = 0;
  char *arr = NULL;
  char *line = NULL;
  int sum = 0;

  if (argc == 2) {
    f = fopen(argv[1], "r");
  } else {
    f = stdin;
  }
  if (f == NULL) {
    printf("Could not open file\n");
    return 1;
  }

  int size = getline(&line, &len, f) - 1;
  arr = malloc(sizeof(char) * size * size);
  rewind(f);

  int i = 0;
  while (getline(&line, &len, f) != -1) {
    for (int j = 0; j < size; j++) {
      arr[i * size + j] = line[j];
    }
    i++;
  }

  fclose(f);

  for (int i = 0; i < size; i++) {
    // printf("%d\t", i);
    int row_sum = 0;
    for (int j = 0; j < size; j++) {
      // printf("%c", arr[i * size + j]);
      char c = arr[i * size + j];
      if (!is_number(c) && c != '.') {
        row_sum += check_around(arr, size, i, j);
      }
    }
    sum += row_sum;
    // printf("\t%d\n", row_sum);
  }

  printf("%d\n", sum);

  free(arr);
  return 0;
}
